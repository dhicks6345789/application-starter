package main

import (
  "fmt"
  "time"
  "os"
  "os/exec"
  "net/http"
)

func main() {
  // Make sure Google Drive is started.
  err := exec.Command("C:\\Program Files\\Google\\Drive File Stream\\launch.bat").Start()
  if err != nil {
    fmt.Println(err)
  }
  
  // Set the Shell registry value temporarily back to "Explorer.exe" so taht Windows Explorer starts in "shell" mode, displaying the desktop, taskbar and so on.
  HTTPResp, HTTPErr := http.Get("http://localhost:8090/setExplorer")
  if err != nil {
    fmt.Println(HTTErr)
  }
  fmt.Println(HTTPResp)
  
  // Wait for Google Drive to be ready.
  tries := 1
  _, pathErr := os.Stat("G:\\My Drive");
  for os.IsNotExist(pathErr) && tries < 12 {
    fmt.Println("Google Drive not ready yet.")
    time.Sleep(5 * time.Second)
    _, pathErr = os.Stat("G:\\My Drive");
    tries = tries + 1
  }
  
  // Start Windows Explorer to display the desktop.
  err = exec.Command("C:\\Windows\\explorer.exe").Start()
  if err != nil {
    fmt.Println(err)
  }
  
  // Set the Shell registry value back to this application.
  err = exec.Command("C:\\Windows\\regedit.exe", "/S", "C:\\Program Files\\Application Starter\\setStarter.reg").Start()
  if err != nil {
    fmt.Println(err)
  }
}
