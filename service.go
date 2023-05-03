package main

import (
  "fmt"
	"log"
  "strings"
  "net/http"
)

func main() {
  http.HandleFunc("/", func (theResponseWriter http.ResponseWriter, theRequest *http.Request) {
    if strings.HasPrefix(theRequest.URL.Path, "/setExplorer") {
		  fmt.Println("Handle setExplorer")
	  }
  })
  log.Fatal(http.ListenAndServe("localhost:8090", nil))
}
