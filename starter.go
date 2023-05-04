package main

import (
  "fmt"
  "time"
  "os"
  "os/exec"
  "net/http"
)

var debugOn bool
func debug(theMessage string) {
  if debugOn {
    fmt.Println(theMessage)
  }
}

func main() {
  // Make sure Google Drive is started.
  err := exec.Command("C:\\Program Files\\Google\\Drive File Stream\\launch.bat").Start()
  if err != nil {
    debug(err.Error())
  }
  
  // Set the Shell registry value temporarily back to "Explorer.exe" so that Windows Explorer starts in "shell" mode, displaying the desktop, taskbar and
  // so on. For this we need to have elevated privalages, so we ask a service running as the system user to do the operation.
  _, _ = http.Get("http://localhost:8090/setExplorer")
  
  // Wait for Google Drive to be ready.
  tries := 1
  _, pathErr := os.Stat("G:\\My Drive");
  for os.IsNotExist(pathErr) && tries < 60 {
    time.Sleep(1 * time.Second)
    _, pathErr = os.Stat("G:\\My Drive");
    tries = tries + 1
  }
  
  // Start Windows Explorer to display the desktop.
  err = exec.Command("C:\\Windows\\explorer.exe").Start()
  if err != nil {
    fmt.Println(err)
  }
  
  // Set the Shell registry value back to this application.
  _, _ = http.Get("http://localhost:8090/setStarter")
}
