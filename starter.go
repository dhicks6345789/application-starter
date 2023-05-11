package main

import (
  "fmt"
  "time"
  "strings"
  "os"
  "os/exec"
)

var debugOn string

func debug(theMessage string) {
  if debugOn == "true" {
    fmt.Println(theMessage)
  }
}

func main() {
  // Get the user's defined profile folder.
  userHome := strings.TrimSpace(os.Getenv("userprofile"))
  if userHome == "" {
    os.Exit(0)
  }
  
  // Pause so Explorer has time to start properly.
  time.Sleep(2 * time.Second)
  
  // Stop Windows Explorer.
  _ = exec.Command("C:\\Windows\\System32\\Taskkill.exe", "/f", "/im", "explorer.exe").Run()
  
  // Set user folder redirects.
  _ = exec.Command("C:\\Windows\\regedit.exe", "/S", "C:\\Program Files\\Application Starter\\setPerUser.reg").Run()
  
  // Check if Google Drive is ready by checking for G:\My Drive...
  tries := 1
  _, pathErr := os.Stat("G:\\My Drive");
  // ...if not, start it...
  if pathErr != nil {
    debug("Starting Google Drive...")
    _ = exec.Command("C:\\Program Files\\Google\\Drive File Stream\\launch.bat").Start()
  }
  // ...and wait for it to be ready...
  for pathErr != nil && tries < 60 {
    debug("Google Drive not ready yet.")
    time.Sleep(1 * time.Second)
    _, pathErr = os.Stat("G:\\My Drive");
    tries = tries + 1
  }
  // ...and wait for G:\My Drive\Desktop to be ready...
  tries = 1
  _, pathErr = os.Stat("G:\\My Drive\\Desktop");
  for pathErr != nil && tries < 60 {
    debug("Desktop folder not ready yet.")
    _ = os.Mkdir("G:\\My Drive\\Desktop", 0750)
    time.Sleep(1 * time.Second)
    _, pathErr = os.Stat("G:\\My Drive\\Desktop");
    tries = tries + 1
  }
  
  // Re-start Windows Explorer.
  _ = exec.Command("C:\\Windows\\Explorer.exe").Start()
}
