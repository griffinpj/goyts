package lib

import (
	"fmt"
	"goyts/utils"
	"io"
	"os"
	"github.com/kkdai/youtube/v2"
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

func extractAudio (video string, keepOriginal bool) error {
    err := utils.ExtractAudio(fmt.Sprintf("tmp/%s.mp4", video), fmt.Sprintf("audio/%s.mp3", video))
	if err != nil {
        return err
	}

    // Remove temporary video file
    if !keepOriginal {
        err = os.Remove(fmt.Sprintf("tmp/%s.mp4", video))
        if err != nil {
            return err
        }
    }
    
    return nil
}

func GetAudio (video string) error {
    err := downloadVideo(video)
	if err != nil {
        return err
	}

    err = extractAudio(video, false)
	if err != nil {
        return err
	}

    return nil
}

