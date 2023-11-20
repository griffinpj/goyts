package lib

import ( 
    "os"
    "errors"
    "fmt"
)

func OllamaSummary (summary string) (error, string) {
    ollamaHost, exists := os.LookupEnv("OLLAMA_HOST")
    if !exists {
        return errors.New("Ollama Host not found"), ""
    }

    // url := fmt.Sprintf("http://%s/api/generate", ollamaHost)

    // TODO summarize transcription with ollama llama2
    // curl http://localhost:11434/api/generate -d '{
    //     "model": "llama2",
    //     "prompt": "Why is the sky blue?",
    //     "stream": false
    // }'

    return nil, summary
}
