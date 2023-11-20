package routes

import (
    "net/http"
    "fmt"
    "errors"
    "os"
    lib "goyts/lib"
)

func YTSummary (w http.ResponseWriter, r *http.Request) {
    videoId := r.FormValue("video")

    if _, err := os.Stat(fmt.Sprintf("audio/%s.mp3", videoId)); errors.Is(err, os.ErrNotExist) {
        // path/to/whatever does not exist
        err := lib.GetAudio(videoId)
        if err != nil {
            panic(err)
        }
    }

    err, content := lib.TranscribeWithWhisper(videoId)
    if err != nil {
        panic(err)
    }

    err, summary := lib.OllamaSummary(string(content))
    if err != nil {
        panic(err)
    }

    fmt.Fprintf(w, string(summary))
}
