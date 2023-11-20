package routes


import (
    "net/http"
    "fmt"
    "time"
)

func LogRequestTime(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// Call the next handler in the chain
		next.ServeHTTP(w, r)

		// Log the request time
		fmt.Printf("%s request --- %s --- %v\n", r.Method, r.URL.Path, time.Since(startTime))
	})
}

