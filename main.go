package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"strconv"
    "strings"
)

const OutDir = "out/"

func extractYoutube(ytURL string, start int, duration int, filePath string) error {
	s := fmt.Sprintf("%v", start)
	t := fmt.Sprintf("%v", duration)

	e := exec.Command("ffmpeg", "-t", t, "-ss", s, "-i", ytURL, "-vf", "scale=350:-1", "-an", "-r", "15", "-crf", "24", filePath)
	stdout, err := e.CombinedOutput()
	if err != nil {
        fmt.Printf("ERROR:\n%v\n", string(stdout))
        return err
	}

    return nil
}

func generateFileName() string {
    out, err := exec.Command("uuidgen").Output()
    if err != nil {
        panic(err)
    }
    s := fmt.Sprintf("%v.gif", strings.TrimSpace(string(out)))
    fmt.Printf("%v\n", s)
    return s
}

func serveGif(w http.ResponseWriter, r *http.Request) {
    parts := strings.Split(r.URL.Path[1:], "/")
    path := OutDir + parts[len(parts)-1]
    fmt.Printf("Serving gif: %v\n", path)
    http.ServeFile(w, r, path)
}

func serveYoutubeExtractor(w http.ResponseWriter, r *http.Request) {
	ytURL := r.URL.Query().Get("video")
	start := r.URL.Query().Get("start")
	dur := r.URL.Query().Get("dur")

	startI, err := strconv.Atoi(start)
	if err != nil {
		panic(err)
	}

	durI, err := strconv.Atoi(dur)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Received args: %v %v %v\n", ytURL, startI, durI)

    fileName := generateFileName()
    fullPath := OutDir + fileName

	err = extractYoutube(ytURL, startI, durI, fullPath)
    if err != nil {
        panic(err)
    }

    http.Redirect(w, r, "/gif/" + fileName, 302)
}

func serveRoot(w http.ResponseWriter, r *http.Request) {
	message := "Hello world"
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", serveRoot)
	http.HandleFunc("/yt_extract", serveYoutubeExtractor)
    http.HandleFunc("/gif/", serveGif)
	err := http.ListenAndServe(":9008", nil)
	if err != nil {
		panic(err)
	}
}
