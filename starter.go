package main

import "fmt"
import "time"
import "os"
import "os/exec"

func main() {
  tries := 1
  _, pathErr := os.Stat("G:\\My Drive");
  for os.IsNotExist(pathErr) && tries < 5 {
    fmt.Println("Google Drive not ready yet.")
    time.Sleep(5 * time.Second)
    _, pathErr = os.Stat("G:\\My Drive");
    tries = tries + 1
  }
  err := exec.Command("C:\\Windows\\explorer.exe").Start()
  if err != nil {
    fmt.Println(err)
  }
}
