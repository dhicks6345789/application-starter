package main

import (
  "fmt"
  "time"
  "os"
  "os/exec"
  //"golang.org/x/sys/windows/registry"
)

func main() {
  driveErr := exec.Command("C:\\Program Files\\Google\\Drive File Stream\\launch.bat").Start()
  if driveErr != nil {
    fmt.Println(driveErr)
  }
  
  regErr := exec.Command("cmd", "/C", "REG", "ADD", "HKEY_LOCAL_MACHINE\\SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion\\Winlogon", "/v", "Shell", "/d", "Explorer.exe", "/f").Start()
  if regErr != nil {
    fmt.Println(regErr)
  }
  /*regKey, regErr := registry.OpenKey(registry.HKEY_LOCAL_MACHINE, "Software\Microsoft\Windows NT\CurrentVersion\Winlogon", registry.QUERY_VALUE)
  if regErr != nil {
    fmt.Println(regErr)
  }
  regKey.SetStringValue("Shell", "Explorer.exe")
  regErr = regKey.Close()*/
  
  tries := 1
  _, pathErr := os.Stat("G:\\My Drive");
  for os.IsNotExist(pathErr) && tries < 10 {
    fmt.Println("Google Drive not ready yet.")
    time.Sleep(5 * time.Second)
    _, pathErr = os.Stat("G:\\My Drive");
    tries = tries + 1
  }
  
  fmt.Println("Done starting.")
  time.Sleep(10 * time.Second)
  
  explorerErr := exec.Command("C:\\Windows\\explorer.exe").Start()
  if explorerErr != nil {
    fmt.Println(explorerErr)
  }
}
