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

func runAndGetOutput(theName string, theArgs ...string) string {
  cmd := exec.Command(theName, theArgs...)
  out, err := cmd.CombinedOutput()
  if err != nil {
    debug("Running command " + theName + " - errror: " + err.Error())
    return ""
  }
  return string(out)
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
  userHome := runAndGetOutput("C:\\Windows\\System32\\cmd.exe", "/C", "echo", "%userprofile%")
  if userHome == "" {
    os.Exit(0)
  }
  userHome = strings.TrimSpace(userHome)
  debug("User Home: " + userHome)
  // Is this the first time this application has run for this user?
  if _, pathErr := os.Stat(userHome + "\\AppData\\Local\\ApplicationStarter"); os.IsNotExist(pathErr) {
    debug("This is user first login.")
    _ = runAndGetOutput("C:\\Windows\\System32\\cmd.exe", "/C", "mkdir %userprofile%\\AppData\\Local\\ApplicationStarter 2>&1")
    _ = runAndGetOutput("C:\\Windows\\System32\\cmd.exe", "/C", "echo > %userprofile%\\AppData\\Local\\ApplicationStarter\\firstRun.txt")
    _ = runAndGetOutput("copy", "C:\Program Files\Application Starter\\starter.exe", "%userprofile%\\AppData\\Microsoft\\Windows\\Start Menu\\Programs\\Startup")
    os.Exit(0)
  }
  //firstRun := false
  if _, pathErr := os.Stat(userHome + "\\AppData\\Local\\ApplicationStarter\\starter.txt"); os.IsNotExist(pathErr) {
    debug("This is a valid run.")
    //_ = runAndGetOutput("C:\\Windows\\System32\\cmd.exe", "/C", "echo > %userprofile%\\AppData\\Local\\ApplicationStarter\\starter.txt")
    if _, firstRunErr := os.Stat(userHome + "\\AppData\\Local\\ApplicationStarter\\firstRun.txt"); !os.IsNotExist(firstRunErr) {
      firstRun = true
      debug("This is a valid first run.")
      _ = runAndGetOutput("C:\\Windows\\System32\\cmd.exe", "/C", "del /q /f %userprofile%\\AppData\\Local\\ApplicationStarter\\firstRun.txt 2>&1")
    }
  } else {
    debug("This is not a valid run.")
    _ = runAndGetOutput("C:\\Windows\\System32\\cmd.exe", "/C", "del /q /f %userprofile%\\AppData\\Local\\ApplicationStarter\\starter.txt 2>&1")
    os.Exit(0)
  }
  
  // Pause so Explorer has time to start properly.
  time.Sleep(4 * time.Second)
  
  // Stop Windows Explorer.
  debug("Stopping Windows Explorer...")
  _ = runAndGetOutput("C:\\Windows\\System32\\Taskkill.exe", "/f", "/im", "explorer.exe")
  
  // Set user folder redirects.
  _ = runAndGetOutput("C:\\Windows\\regedit.exe", "/S", "C:\\Program Files\\Application Starter\\setPerUser.reg")
  
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
    _ = runAndGetOutput("C:\\Windows\\System32\\cmd.exe", "/C", "mkdir", "G:\\My Drive\\Desktop")
    time.Sleep(1 * time.Second)
    _, pathErr = os.Stat("G:\\My Drive\\Desktop");
    tries = tries + 1
  }
  
  // Re-start Windows Explorer.
  err := exec.Command("C:\\Windows\\Explorer.exe").Start()
  if err != nil {
    debug(err.Error())
  }
  
  /*if firstRun {
    _ = runAndGetOutput("C:\\Windows\\System32\\cmd.exe", "/C", "del /q /f %userprofile%\\AppData\\Local\\ApplicationStarter\\starter.txt 2>&1")
  }*/
}
