package main

import (
  "fmt"
  "log"
  "strings"
  "os/exec"
  "net/http"
)

func main() {
  http.HandleFunc("/", func (theResponseWriter http.ResponseWriter, theRequest *http.Request) {
    if strings.HasPrefix(theRequest.URL.Path, "/setExplorer") {
      fmt.Println("Handle setExplorer")
      
      cmd := exec.Command("C:\\Windows\\System32\\reg.exe", "Query", "HKEY_USERS")
      out, err := cmd.CombinedOutput()
      if err != nil {
        fmt.Printf("Query to registry failed: %s\n", err)
      } else {
        for _, user := range strings.Split(string(out), "\n") {
          userSplit := strings.Split(user, "\\")
          if len(userSplit) == 2 {
            userID := strings.TrimSpace(userSplit[1])
            if userID != ".DEFAULT" {
              fmt.Println("AAA" + userID + "BBB")
            }
          }
        }
        err = exec.Command("C:\\Windows\\regedit.exe", "/S", "C:\\Program Files\\Application Starter\\setExplorer.reg").Start()
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
