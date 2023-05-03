package main

import "fmt"
import "time"
import "os"
import "os/exec"

func main() {
  _, pathErr := os.Stat("G:\\My Drive");
  for os.IsNotExist(pathErr) {
    fmt.Println("Google Drive not ready yet.")
    time.Sleep(5 * time.Second)
    _, pathErr = os.Stat("G:\\My Drive");
  }
  err := exec.Command("C:\\Windows\\explorer.exe").Start()
  if err != nil {
    fmt.Println(err)
  }
}
