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
  // Stop Windows Explorer.
  //debug("Stopping Windows Explorer...")
  //err := exec.Command("C:\\Windows\\System32\\Taskkill.exe", "/f", "/im", "explorer.exe").Run()
  //if err != nil {
    //debug(err.Error())
  //}
  //if debugOn == "true" {
    //time.Sleep(10 * time.Second)
  //}
  
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
  for pathErr != nil && tries < 60 {
    debug("Google Drive not ready yet.")
    time.Sleep(1 * time.Second)
    _, pathErr = os.Stat("G:\\My Drive");
    tries = tries + 1
  }
  
  // Set the Shell registry value temporarily back to "Explorer.exe" so that Windows Explorer starts in "shell" mode, displaying the desktop, taskbar and
  // so on. For this we need to have elevated privalages, so we ask a service running as the system user to do the operation.
  //callEndpoint("http://localhost:8090/setExplorer")
  //time.Sleep(60 * time.Second)
  
  // Start Windows Explorer to display the desktop.
  //err = exec.Command("C:\\Windows\\explorer.exe").Start()
  //if err != nil {
    //debug(err.Error())
  //}
  
  //// Set the Shell registry value back to this application.
  //callEndpoint("http://localhost:8090/setStarter")
  
  if debugOn == "true" {
    time.Sleep(30 * time.Second)
  }
}
