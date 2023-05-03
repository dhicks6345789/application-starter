package main

import "fmt"
import "os/exec"

func main() {
  if _, err := os.Stat("G:\\My Drive"); os.IsNotExist(err) {
    fmt.Println("Path does not exist!")
  }
  err := exec.Command("C:\\Windows\\explorer.exe").Start()
  if err != nil {
    fmt.Println(err)
  }
}
