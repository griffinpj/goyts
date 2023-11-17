package lib

import (
    "io"
    "os"
    "fmt"
    "github.com/kkdai/youtube/v2"
)

func GetAudio (video string) {
	client := youtube.Client{}

	audio, err := client.GetVideo(video)
	if err != nil {
		panic(err)
	}

	formats := audio.Formats.WithAudioChannels() // only get videos with audio
	stream, _, err := client.GetStream(audio, &formats[0])
	if err != nil {
		panic(err)
	}

	defer stream.Close()
    
    newFileName := fmt.Sprintf("tmp/%s.mp4", video)
	file, err := os.Create(newFileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}
}

