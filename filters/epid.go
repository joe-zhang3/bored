package main

import (
	"fmt"
	"net/http"
)

func PluginMain(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("epid.so is called")
	w.Header().Add("X-EPID", "xxx-xxxx-xxxxx-x-xxxxx")
	w.WriteHeader(http.StatusOK)
}
