package main

import "fmt"
import "os/exec"

func main() {
  err := exec.Command("C:\\Windows\\explorer.exe").Run()
  if err != nil {
    fmt.Println(err)
  }
}
