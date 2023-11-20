package utils

import (
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func ExtractAudio (src string, dest string) error {
    // Convert local video file to audio via ffmpeg
    return ffmpeg.
        Input(src, ffmpeg.KwArgs{"format": "mp4"}).
        Output(dest, ffmpeg.KwArgs{"format": "mp3"}).
        OverWriteOutput().
        ErrorToStdOut().
        Run()
}
