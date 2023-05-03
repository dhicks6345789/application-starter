package main

import "fmt"
import "time"
import "os"
import "os/exec"

func main() {
  driveErr := exec.Command("C:\\Program Files\\Google\\Drive File Stream\\launch.bat").Start()
  if driveErr != nil {
    fmt.Println(driveErr)
  }
  
  regErr := exec.Command("REG ADD HKEY_LOCAL_MACHINE\\SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion\\Winlogon /v Shell /d Explorer.exe /f").Start()
  if regErr != nil {
    fmt.Println(regErr)
  }
  
  tries := 1
  _, pathErr := os.Stat("G:\\My Drive");
  for os.IsNotExist(pathErr) && tries < 10 {
    fmt.Println("Google Drive not ready yet.")
    time.Sleep(5 * time.Second)
    _, pathErr = os.Stat("G:\\My Drive");
    tries = tries + 1
  }
  
  explorerErr := exec.Command("C:\\Windows\\explorer.exe").Start()
  if explorerErr != nil {
    fmt.Println(explorerErr)
  }
}
