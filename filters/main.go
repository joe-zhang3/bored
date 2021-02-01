package main

import (
  "fmt"
  "plugin"
  "net/http"
)

func main(){
  p, err := plugin.Open("epid.so")

  if err != nil {
    fmt.Printf("cannot find the so file")
    panic(err)
  }

  s, err := p.Lookup("PluginMain")

  if err != nil {
    fmt.Printf("cannot find the symbol")
  }

  plug, _ := s.(func(http.ResponseWriter, *http.Request))

  request, _ := http.NewRequest(http.MethodPost, "http://localhost:8080", nil)
  plug(nil, request)
}

