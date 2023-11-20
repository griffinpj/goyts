package utils

import (
    "os"
    "net/http"
    "io"
    "path/filepath"
    "bytes"
    "mime/multipart"
)

func PostFile (fileName string, url string) (error, [] byte) {
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

    content, err := io.ReadAll(res.Body)
    return err, content
}

func PostJson (json string, url string) (error, [] byte) {
    jsonData := [] byte(json)

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
    if err != nil {
        return err, nil
    }

    req.Header.Set("Content-Type", "application/json; charset=UTF-8")

    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
        return err, nil
    }

    defer res.Body.Close()
   
    body, err := io.ReadAll(res.Body)
    return err, body
}
