package routes

import (
    "net/http"
    "fmt"
    lib "goyts/lib"
)

func YTSummary (w http.ResponseWriter, r *http.Request) {
    videoId := r.FormValue("video")
    err := lib.GetAudio(videoId)
    if err != nil {
        panic(err)
    }

    err, content := lib.TranscribeWithWhisper(videoId)
    if err != nil {
        panic(err)
    }

    err, summary := lib.OllamaSummary(string(content))
    if err != nil {
        panic(err)
    }

    fmt.Printf("%s", summary)

    fmt.Fprintf(w, string(summary))
}
