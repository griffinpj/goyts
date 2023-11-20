package main

import (
    "net/http"
    "fmt"
    "github.com/joho/godotenv"
    Routes "goyts/routes"
    "log"
)

// init is invoked before main()
func init() {
    // loads values from .env into the system
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
}

func setupRoutes () {
    for _, route := range Routes.InitRoutes() {
        handler := route.Handler
        localRoute := route

        for _, middleware := range route.Middleware {
            handler = middleware(handler)
        }

        http.HandleFunc(route.Path, func(w http.ResponseWriter, r *http.Request) {
            if r.Method == localRoute.Method {
                handler.ServeHTTP(w, r)
            } else {
                http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
            }
        })
    }
}

func main () {
    setupRoutes()

    port := 8080
	fmt.Printf("Server running on :%d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
