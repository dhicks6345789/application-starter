package main

import (
  "os"
  "time"
  "strings"
  "os/exec"
)

func main() {
  // Get the user's defined profile folder.
  userName := strings.TrimSpace(os.Getenv("username"))
  userHome := strings.TrimSpace(os.Getenv("userprofile"))
  if userHome == "" {
    os.Exit(0)
  }
  
  // Make the user's local (and, hopefully, unused) Desktop folder read-only.
  _ = exec.Command("C:\\Windows\\System32\\icacls.exe", userHome + "\\Desktop", "/inheritance:r", "/grant:r", "" + userName + ":R").Run()
  
  // If this is a user's first run, we need to quit so the first run application can run instead.
  if _, pathErr := os.Stat(userHome + "\\AppData\\Local\\ApplicationStarter"); os.IsNotExist(pathErr) {
    os.Exit(0)
  }
  
  // Stop Windows Explorer.
  _ = exec.Command("C:\\Windows\\System32\\Taskkill.exe", "/f", "/im", "explorer.exe").Run()
  
  // Check if Google Drive is ready by checking for G:\My Drive...
  tries := 1
  _, pathErr := os.Stat(userHome + "\\Google Drive\\My Drive");
  // ...if not, start it...
  if pathErr != nil {
    _ = exec.Command("C:\\Program Files\\Google\\Drive File Stream\\launch.bat").Start()
  }
  // ...and wait for it to be ready.
  for pathErr != nil && tries < 60 {
    time.Sleep(1 * time.Second)
    _, pathErr = os.Stat(userHome + "\\Google Drive\\My Drive");
    tries = tries + 1
  }
  
  // Non-peristantly map G: drive to Google Drive for the local user only.
  _ = exec.Command("C:\\Windows\\System32\\subst.exe", "G:", userHome + "\\Google Drive").Start()
  
  // Re-start Windows Explorer.
  _ = exec.Command("C:\\Windows\\Explorer.exe").Start()
}
