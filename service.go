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
      err := exec.Command("C:\\Windows\\regedit.exe", "/S", "C:\\Program Files\\Application Starter\\setExplorer.reg").Start()
      if err != nil {
        fmt.Fprint(theResponseWriter, err)
      } else {
        fmt.Fprint(theResponseWriter, "OK")
      }
	  }
  })
  log.Fatal(http.ListenAndServe("localhost:8090", nil))
}
