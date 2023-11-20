package lib

import ( 
    "os"
    "errors"
    "fmt"
    "strings"
    utils "goyts/utils"
)

func OllamaSummary (summary string) (error, [] byte) {
    ollamaHost, exists := os.LookupEnv("OLLAMA_HOST")
    if !exists {
        return errors.New("Ollama Host not found"), nil
    }

    url := fmt.Sprintf("http://%s/api/generate", ollamaHost)

    jsonBody := fmt.Sprintf(`{
        "model": "llama2:13b",
        "prompt": "Please summarize the following video transcription. Your response will only prefixed with Summary: X, where X is the start of your summary without introduction., '%s'",
        "stream": false
    }`, strings.Replace(summary, "\n", "", -1))

    err, body := utils.PostJson(jsonBody, url)
    if err != nil {
        return err, nil
    }

    return nil, body
}
