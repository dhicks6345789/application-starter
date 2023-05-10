package main

import (
  "fmt"
  "time"
  "os"
  "os/exec"
  "net/http"
  "io/ioutil"
)

var debugOn string

func debug(theMessage string) {
  if debugOn == "true" {
    fmt.Println(theMessage)
  }
}

func runAndGetOutput(theName string, theArgs ...string) (string, error) {
  cmd := exec.Command(theName, theArgs...)
  out, err := cmd.CombinedOutput()
  return string(out), err
}

func callEndpoint(theEndpoint string) {
  debug("Calling endpoint: " + theEndpoint)
  resp, err := http.Get(theEndpoint)
  if err != nil {
    debug(err.Error())
  }
  body, err := ioutil.ReadAll(resp.Body)
  debug("Result: " + string(body))
  resp.Body.Close()
}

func main() {
  // If we're running the first time a user has logged in, the user's defined user profile folder won't actually exist.
  firstLogin := false
  userHome, err := runAndGetOutput("C:\\Windows\\System32\\cmd.exe", "/C", "echo", "%userprofile%")
  if err != nil {
    debug(err.Error())
  } else {
    debug("User Home: " + userHome)
    if _, pathErr := os.Stat(userHome + "\AppData\\Local\\ApplicationStarter"); os.IsNotExist(pathErr) {
      firstLogin = true
      _, mkdirErr := runAndGetOutput("C:\\Windows\\System32\\cmd.exe", "/C", "mkdir", "%userprofile%\\AppData\\Local\\ApplicationStarter")
      if mkdirErr != nil {
        debug(mkdirErr.Error())
      }
    }
  }
  
  // Pause so Explorer has time to start properly.
  time.Sleep(2 * time.Second)
  
  if firstLogin {
    os.Exit(0)
  }
  
  // Stop Windows Explorer.
  debug("Stopping Windows Explorer...")
  err = exec.Command("C:\\Windows\\System32\\Taskkill.exe", "/f", "/im", "explorer.exe").Run()
  if err != nil {
    debug(err.Error())
  }
  
  // Set user folder redirects.
  err = exec.Command("C:\\Windows\\regedit.exe", "/S", "C:\\Program Files\\Application Starter\\setPerUser.reg").Run()
  if err != nil {
    debug(err.Error())
  }
  
  // Check if Google Drive is ready...
  tries := 1
  _, pathErr := os.Stat("G:\\My Drive");
  // ...if not, start it...
  if pathErr != nil {
    debug("Starting Google Drive...")
    err := exec.Command("C:\\Program Files\\Google\\Drive File Stream\\launch.bat").Start()
    if err != nil {
      debug(err.Error())
    }
  }
  // ...and wait for it to be ready.
  for pathErr != nil && tries < 60 && firstLogin == false {
    debug("Google Drive not ready yet.")
    time.Sleep(1 * time.Second)
    _, pathErr = os.Stat("G:\\My Drive");
    tries = tries + 1
  }
  
  // Re-start Windows Explorer.
  err = exec.Command("C:\\Windows\\Explorer.exe").Start()
  if err != nil {
    debug(err.Error())
  }
}
