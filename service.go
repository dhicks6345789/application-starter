package main

import (
  "os"
  "fmt"
  "log"
  "errors"
  "strings"
  "os/exec"
  "net/http"
  "io/ioutil"
)

func runAndGetOutput(theName string, theArgs ...string) (string, error) {
  result := ""
  cmd := exec.Command(theName, theArgs...)
  out, err := cmd.CombinedOutput()
  if err != nil {
    //result = "Error running command: " + strings.Join(theArgs, " ") + " - result: " + err.Error()
    return "", err
  } //else {
    //result = string(out)
  //}
  return result, nil
}

/* An application intended to run as a Windows service (installed via NSSM) to handle requests from its companion application to set various Windows registry entries.
This service should run as the system user so it has permissions to set registry entries. */
func main() {
  http.HandleFunc("/", func (theResponseWriter http.ResponseWriter, theRequest *http.Request) {
    // Handle the "setExplorer" endpoint - set the user shell to "Explorer.exe", also make sure per-user registry settings are set.
    if strings.HasPrefix(theRequest.URL.Path, "/setExplorer") {
      fmt.Println("Handle setExplorer")
      
      // Get a list of users on this machine.
      out, err := runAndGetOutput("C:\\Windows\\System32\\reg.exe", "Query", "HKEY_USERS")
      
      //cmd := exec.Command("C:\\Windows\\System32\\reg.exe", "Query", "HKEY_USERS")
      //out, err := cmd.CombinedOutput()
      if err != nil {
        fmt.Printf("Query to registry failed: %s\n", err)
      } else {
        // Read the per-user registry settings template file.
        perUserFile, perUserErr := os.Open("C:\\Program Files\\Application Starter\\setPerUser.reg")
        if perUserErr != nil {
          fmt.Printf("Error opening setPerUser.reg: %s\n", perUserErr)
        }
        perUserText, _ := ioutil.ReadAll(perUserFile)
        perUserString := string(perUserText)
        perUserFile.Close()
        
        // Step through each user found, excluding special cases.
        for _, user := range strings.Split(string(out), "\n") {
          userSplit := strings.Split(user, "\\")
          if len(userSplit) == 2 {
            userID := strings.TrimSpace(userSplit[1])
            if userID != ".DEFAULT" && !strings.HasSuffix(userID, "_Classes") {
              pathString := "C:\\Program Files\\Application Starter\\Users\\" + userID + ".reg"
              // For each user, if we haven't already written their per-user registry settings, do so now. The user-named cached file acts as an
              // indicator we've already done the settings for that user.
              if _, pathErr := os.Stat(pathString); errors.Is(pathErr, os.ErrNotExist) {
                fileWriteErr := os.WriteFile(pathString, []byte(strings.ReplaceAll(perUserString, "HKEY_CURRENT_USER\\", "HKEY_USERS\\" + userID + "\\")), 0644)
                if fileWriteErr != nil {
                  fmt.Printf("Error writing file:: %s\n", pathString)
                } else {
                  _ = exec.Command("C:\\Windows\\regedit.exe", "/S", pathString).Start()
                }
              }
            }
          }
        }
        
        // Set the user shell (for all users) to be "Explorere.exe".
        err = exec.Command("C:\\Windows\\regedit.exe", "/S", "C:\\Program Files\\Application Starter\\setExplorer.reg").Start()
        
        // Return an "OK" message for the calling application.
        if err != nil {
          fmt.Fprint(theResponseWriter, err)
        } else {
          fmt.Fprint(theResponseWriter, "OK")
        }
      }
    }
    if strings.HasPrefix(theRequest.URL.Path, "/setStarter") {
      fmt.Println("Handle setStarter")
      err := exec.Command("C:\\Windows\\regedit.exe", "/S", "C:\\Program Files\\Application Starter\\setStarter.reg").Start()
      if err != nil {
        fmt.Fprint(theResponseWriter, err)
      } else {
        fmt.Fprint(theResponseWriter, "OK")
      }
    }
  })
  log.Fatal(http.ListenAndServe("localhost:8090", nil))
}
