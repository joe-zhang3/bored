package main

import (
  "net/http"
  "fmt"
)

func PluginMain(w http.ResponseWriter, r *http.Request){
  fmt.Printf("hellow")
}

