package main

import "fmt"
import "time"
import "os"
import "os/exec"

func main() {
  driveErr := exec.Command("C:\\Program Files\\Google\Drive File Stream\\launch.bat").Start()
  if driveErr != nil {
    fmt.Println(driveErr)
  }
  
  tries := 1
  _, pathErr := os.Stat("G:\\My Drive");
  for os.IsNotExist(pathErr) && tries < 5 {
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
