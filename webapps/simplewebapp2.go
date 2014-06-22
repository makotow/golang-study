package main

import (
  "fmt"
  "net/http"
  "log"
)

func get_status() (string) {
  return "on"
}

func switch_handler(writer http.ResponseWriter, request *http.Request) {
  status := get_status()
  fmt.Fprintf(writer, status)
}

func main() {
  http.HandleFunc("/switch/", switch_handler)

  log.Fatal(http.ListenAndServe(":8080", nil))
}
