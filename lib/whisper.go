package lib

import (
    "os"
    "fmt"
    "errors"
    utils "goyts/utils"
)

func TranscribeWithWhisper (videoId string) (error, [] byte) {
    whisperHost, exists := os.LookupEnv("WHISPER_HOST")
    if !exists {
        return errors.New("Whisper Host not found"), nil
    }

    url := fmt.Sprintf("http://%s/asr?encode=true&task=transcribe&language=en&word_timestamps=false&output=txt", whisperHost)
    fileName := fmt.Sprintf("audio/%s.mp3", videoId)
    err, content := utils.PostFile(fileName, url)
    if err != nil {
        return err, nil
    }

    return nil, content
}
