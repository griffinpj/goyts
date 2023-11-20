package lib

import (
    "net/http"
    "os"
    "io"
    "path/filepath"
    "io/ioutil"
    "fmt"
    "bytes"
    "mime/multipart"
    "errors"
)

func TranscribeWithWhisper (videoId string) (error, [] byte) {
    whisperHost, exists := os.LookupEnv("WHISPER_HOST")
    if !exists {
        return errors.New("Whisper Host not found"), nil
    }

    url := fmt.Sprintf("http://%s/asr?encode=true&task=transcribe&language=en&word_timestamps=false&output=txt", whisperHost)

    fileName := fmt.Sprintf("audio/%s.mp3", videoId)
    audio, err := os.Open(fileName)
    if err != nil {
        return err, nil
    }

    defer audio.Close()

    body := &bytes.Buffer{}
    writer := multipart.NewWriter(body)

    reqBody, err := writer.CreateFormFile("audio_file", filepath.Base(audio.Name()))
    if err != nil {
        return err, nil
    }

    io.Copy(reqBody, audio)
    writer.Close()

    req, err := http.NewRequest("POST", url, body)
    if err != nil {
        return err, nil
    }

    req.Header.Add("Content-Type", writer.FormDataContentType())
    client := &http.Client{}

    res, err := client.Do(req)
    if err != nil {
        return err, nil
    }

    content, err := ioutil.ReadAll(res.Body)
    if err != nil {
        return err, nil
    }

    return nil, content
}
