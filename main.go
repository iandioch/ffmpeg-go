package main

import (
    "net/http"
)

func serveYoutubeExtractor(w http.ResponseWriter, r *http.Request) {
    ytURL := r.URL.Query().Get("video")
    start := r.URL.Query().Get("start")
    dur := r.URL.Query().Get("dur")

    // Do ffmpeg stuff

    w.Write([]byte("response: " +  ytURL + ", " + start + ", " + dur))
}

func serveRoot(w http.ResponseWriter, r *http.Request) {
    message := "Hello world"
    w.Write([]byte(message))
}

func main() {
    http.HandleFunc("/", serveRoot)
    http.HandleFunc("/yt_extract", serveYoutubeExtractor)
    err := http.ListenAndServe(":9008", nil)
    if err != nil {
        panic(err)
    }
}
