package lib

import (
    "io"
    "os"
    "fmt"
    "github.com/kkdai/youtube/v2"
    ffmpeg "github.com/u2takey/ffmpeg-go"
)

func downloadVideo (video string) error {
	client := youtube.Client{}

    // Get YT Video Stream
	videoCtx, err := client.GetVideo(video)
	if err != nil {
        return err
	}

	formats := videoCtx.Formats.WithAudioChannels()
	stream, _, err := client.GetStream(videoCtx, &formats[0])
	if err != nil {
        return err
	}

	defer stream.Close()
   
    // Create temporary local video file
	file, err := os.Create(fmt.Sprintf("tmp/%s.mp4", video))
	if err != nil {
        return err
	}

	defer file.Close()

    // Copy YT Video Stream to local file
	_, err = io.Copy(file, stream)
	if err != nil {
        return err
	}
    return nil
}

func extractAudio (video string) error {
    // Convert local video file to audio via ffmpeg
    err := ffmpeg.
        Input(fmt.Sprintf("tmp/%s.mp4", video), ffmpeg.KwArgs{"format": "mp4"}).
        Output(fmt.Sprintf("audio/%s.mp3", video), ffmpeg.KwArgs{"format": "mp3"}).
        OverWriteOutput().
        ErrorToStdOut().
        Run()

	if err != nil {
        return err
	}

    // Remove temporary video file
	err = os.Remove(fmt.Sprintf("tmp/%s.mp4", video))
	if err != nil {
        return err
	}
    
    return nil
}

func GetAudio (video string) error {
    err := downloadVideo(video)
	if err != nil {
        return err
	}

    err = extractAudio(video)
	if err != nil {
        return err
	}

    return nil
}

