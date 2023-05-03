package main

import (
  "time"
)

func main() {
  for {
    time.Sleep(5 * time.Second)
  }
  select {} // Block so the program stays resident.
}
