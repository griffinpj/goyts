package main

import (
    "net/http"
    "fmt"
    "github.com/joho/godotenv"
    "log"
    "os"
    "errors"
    Routes "goyts/routes"
)

func init() {
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
    
    port, exists := os.LookupEnv("PORT")
    if !exists {
        panic(errors.New("No port defined"))
    }

	fmt.Printf("Server running on :%s...\n", port)

    err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
    if err != nil {
        panic (err)
    }
}
