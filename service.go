package main

import (
  "fmt"
  "strings"
  "net/http"
)

func main() {
  http.HandleFunc("/", func (theResponseWriter http.ResponseWriter, theRequest *http.Request) {
    if strings.HasPrefix(theRequest.URL.Path, "/setExplorer") {
		  fmt.Println("Handle setExplorer")
	  }
  })
}
