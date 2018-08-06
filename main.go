package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"strconv"
    "strings"
)

func extractYoutube(ytURL string, start int, duration int) string {
	s := fmt.Sprintf("%v", start)
	t := fmt.Sprintf("%v", duration)

	e := exec.Command("ffmpeg", "-t", t, "-ss", s, "-i", ytURL, "-vf", "scale=350:-1", "-an", "-r", "15", "-crf", "24", "banana.gif")
	fmt.Printf("%s %s\n", e.Path, e.Args)
	stdout, err := e.CombinedOutput()
	if err != nil {
		fmt.Printf("%v\n", string(stdout))
		panic(err)
	}

	res := string(stdout)
	fmt.Printf("%s\n", res)
	return res
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

    generateFileName()

	res := extractYoutube(ytURL, startI, durI)

	w.Write([]byte("response: " + ytURL + ", " + start + ", " + dur + ", " + res))
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
