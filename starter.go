package main

import "fmt"
import "time"
import "os"
import "os/exec"

func main() {
  _, err := os.Stat("G:\\My Drive");
  for os.IsNotExist(err) {
    fmt.Println("Google Drive not ready yet.")
    time.Sleep(5 * time.Second)
    _, err = os.Stat("G:\\My Drive");
  }
  err := exec.Command("C:\\Windows\\explorer.exe").Start()
  if err != nil {
    fmt.Println(err)
  }
}
