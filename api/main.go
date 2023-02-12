package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, req *http.Request){
    fmt.Fprintf(w, "hello from %s\n", req.RequestURI)
}
func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8090", nil)
}
