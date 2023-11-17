package main

import (
    "net/http"
    "fmt"
    Routes "goyts/routes"
)

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
