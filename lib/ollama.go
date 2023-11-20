package lib

import ( 
    "os"
    "errors"
    "fmt"
    "bytes"
    "io"
    "net/http"
    "strings"
)

func OllamaSummary (summary string) (error, [] byte) {
    ollamaHost, exists := os.LookupEnv("OLLAMA_HOST")
    if !exists {
        return errors.New("Ollama Host not found"), nil
    }

    url := fmt.Sprintf("http://%s/api/generate", ollamaHost)

    jsonBody := fmt.Sprintf(`{
        "model": "llama2:13b",
        "prompt": "Please summarize the following in as many needed paragraphs, '%s'",
        "stream": false
    }`, strings.Replace(summary, "\n", "", -1))

    var jsonData = []byte(jsonBody)

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
    if err != nil {
        return err, nil
    }

    return nil, body
}
