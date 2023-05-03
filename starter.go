package main

import "fmt"
import "time"
import "os"
import "os/exec"

func main() {
  if _, err := os.Stat("G:\\My Drive"); os.IsNotExist(err) {
    fmt.Println("Google Drive not ready yet.")
    time.Sleep(5 * time.Second)
  }
  err := exec.Command("C:\\Windows\\explorer.exe").Start()
  if err != nil {
    fmt.Println(err)
  }
}
