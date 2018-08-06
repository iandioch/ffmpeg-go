package main

import (
    "net/http"
)

func serveRoot(w http.ResponseWriter, r *http.Request) {
    message := "Hello world"
    w.Write([]byte(message))
}

func main() {
    http.HandleFunc("/", serveRoot)
    err := http.ListenAndServe(":9008", nil)
    if err != nil {
        panic(err)
    }
}
