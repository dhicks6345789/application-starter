package main

import (
  "fmt"
  "time"
  "strings"
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
  // Get the user's defined profile folder.
  userHome, err := runAndGetOutput("C:\\Windows\\System32\\cmd.exe", "/C", "echo", "%userprofile%")
  if err != nil {
    debug(err.Error())
  } else {
    userHome = strings.TrimSpace(userHome)
    debug("User Home: " + userHome)
    // Is this the first time this application has run for this user?
    if _, pathErr := os.Stat(userHome + "\\AppData\\Local\\ApplicationStarter"); os.IsNotExist(pathErr) {
      debug("This is user first login.")
      _, mkdirErr := runAndGetOutput("C:\\Windows\\System32\\cmd.exe", "/C", "mkdir", "%userprofile%\\AppData\\Local\\ApplicationStarter")
      if mkdirErr != nil {
        debug(mkdirErr.Error())
      }
      os.Exit(0)
    }
  }
  
  // Pause so Explorer has time to start properly.
  time.Sleep(2 * time.Second)
  
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
  
  // Check if Google Drive is ready by checking for G:\My Drive...
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
  // ...and wait for it to be ready...
  for pathErr != nil && tries < 60 {
    debug("Google Drive not ready yet.")
    time.Sleep(1 * time.Second)
    _, pathErr = os.Stat("G:\\My Drive");
    tries = tries + 1
  }
  // ...and wait for G:\My Drive\Desktop to be ready...
  tries = 1
  _, pathErr = os.Stat("G:\\My Drive\\Desktop");
  for pathErr != nil && tries < 60 {
    debug("Desktop folder not ready yet.")
    _, mkdirErr := runAndGetOutput("C:\\Windows\\System32\\cmd.exe", "/C", "mkdir", "G:\\My Drive\\Desktop")
    if mkdirErr != nil {
      debug(mkdirErr.Error())
    }
    time.Sleep(1 * time.Second)
    _, pathErr = os.Stat("G:\\My Drive\\Desktop");
    tries = tries + 1
  }
  
  // Re-start Windows Explorer.
  err = exec.Command("C:\\Windows\\Explorer.exe").Start()
  if err != nil {
    debug(err.Error())
  }
}
