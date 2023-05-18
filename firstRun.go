package main

import (
  "time"
  "strings"
  "os"
  "os/exec"
)

func main() {
  // Get the user's defined profile folder.
  userHome := strings.TrimSpace(os.Getenv("userprofile"))
  if userHome == "" {
    os.Exit(0)
  }
  
  // Is this the first time this application has run for this user?
  if _, pathErr := os.Stat(userHome + "\\AppData\\Local\\ApplicationStarter"); !os.IsNotExist(pathErr) {
    os.Exit(0)
  }
  
  // Stop Windows Explorer.
  _ = exec.Command("C:\\Windows\\System32\\Taskkill.exe", "/f", "/im", "explorer.exe").Run()
  
  // Make sure the Google Drive mount point folder exists.
  if _, gDriveErr := os.Stat(userHome + "\\Google Drive"); os.IsNotExist(gDriveErr) {
    _ = os.Mkdir(userHome + "\\Google Drive", 0750)
  }
  
  // Set user folder redirects.
  _ = exec.Command("C:\\Windows\\regedit.exe", "/S", "C:\\Program Files\\Application Starter\\setPerUser.reg").Run()
  
  // Check if Google Drive is ready by checking for G:\My Drive...
  _, pathErr := os.Stat(userHome + "\\Google Drive\\My Drive")
  // ...if not, start it...
  if pathErr != nil {
    _ = exec.Command("C:\\Program Files\\Google\\Drive File Stream\\launch.bat").Start()
  }
  // ...and wait for it to be ready...
  for pathErr != nil {
    time.Sleep(1 * time.Second)
    _, pathErr = os.Stat(userHome + "\\Google Drive\\My Drive");
  }
  // ...and wait for G:\My Drive\Desktop to be ready...
  tries := 1
  _, pathErr = os.Stat(userHome + "\\Google Drive\\My Drive\\Desktop")
  for pathErr != nil && tries < 60 {
    _ = os.Mkdir(userHome + "\\Google Drive\\My Drive\\Desktop", 0750)
    _ = os.Mkdir(userHome + "\\Google Drive\\My Drive\\Documents", 0750)
    time.Sleep(1 * time.Second)
    _, pathErr = os.Stat(userHome + "\\Google Drive\\My Drive\\Desktop")
    tries = tries + 1
  }
  
  // Non-peristantly map G: drive to Google Drive for the local user only.
  _ = exec.Command("C:\\Windows\\System32\\subst.exe", "G:", userHome + "\\Google Drive").Start()
  
  // Re-start Windows Explorer.
  _ = exec.Command("C:\\Windows\\Explorer.exe").Start()
  
  // Create the application data folder - really just used as a "first run" marker.
  _ = os.Mkdir(userHome + "\\AppData\\Local\\ApplicationStarter", 0750)
}
