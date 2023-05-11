package main

import (
  "os"
  "fmt"
  "time"
  "strings"
  "os/exec"
)

func main() {
  // Get the user's defined profile folder.
  userHome := strings.TrimSpace(os.Getenv("userprofile"))
  userDomain := strings.TrimSpace(os.Getenv("userdomain"))
  userName := strings.TrimSpace(os.Getenv("username"))
  if userHome == "" {
    os.Exit(0)
  }
  fmt.Println(userDomain)
  fmt.Println(userName)
  
  // Make the user's local (and, hopefully, unused) Desktop folder read-only.
  out, err := exec.Command("C:\\Windows\\System32\\icacls.exe", "\"" + userHome + "\\Desktop\\*\"", "/deny", "\"%userdomain%\\%username%\":(OI)(WA)").CombinedOutput()
  if err != nil {
    fmt.Println(err.Error())
  }
  fmt.Println(string(out))
  os.Exit(0)
  
  // If this is a user's first run, we need to quit so the first run application can run instead.
  if _, pathErr := os.Stat(userHome + "\\AppData\\Local\\ApplicationStarter"); os.IsNotExist(pathErr) {
    os.Exit(0)
  }
  
  // Stop Windows Explorer.
  _ = exec.Command("C:\\Windows\\System32\\Taskkill.exe", "/f", "/im", "explorer.exe").Run()
  
  // Check if Google Drive is ready by checking for G:\My Drive...
  tries := 1
  _, pathErr := os.Stat("G:\\My Drive");
  // ...if not, start it...
  if pathErr != nil {
    _ = exec.Command("C:\\Program Files\\Google\\Drive File Stream\\launch.bat").Start()
  }
  // ...and wait for it to be ready.
  for pathErr != nil && tries < 60 {
    time.Sleep(1 * time.Second)
    _, pathErr = os.Stat("G:\\My Drive");
    tries = tries + 1
  }
  
  // Re-start Windows Explorer.
  _ = exec.Command("C:\\Windows\\Explorer.exe").Start()
}
