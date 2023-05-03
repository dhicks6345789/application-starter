package main

import "fmt"
import "exec"

func main() {
  fmt.Println("Hello, world!")
  err := exec.Command("C:\\Windows\\explorer.exe").Run()
  if err != nil {
    log.Fatal(err)
  }
}
