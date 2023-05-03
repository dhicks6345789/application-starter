package main

import (
  "fmt"
  "net/http"
)

func main() {
  http.HandleFunc("/", func (theResponseWriter http.ResponseWriter, theRequest *http.Request) {
	  if strings.HasPrefix(theRequest.URL.Path, "/setExplorer") {
		  fmt.println("Handle setExplorer")
	  }
  })
}
