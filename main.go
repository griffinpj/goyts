package main

import (
    "net/http"
    "fmt"
    "time"
    "html/template"
)

type Route struct {
    Path string
    Method string
    Handler http.Handler
    Middleware [] func (http.Handler) http.Handler
}

type ViewData struct {
    Message string
}

func renderTemplate(templatePath string, data interface {}) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(templatePath)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})
}

func logRequestTime(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// Call the next handler in the chain
		next.ServeHTTP(w, r)

		// Log the request time
		fmt.Printf("%s request --- %s --- %v\n", r.Method, r.URL.Path, time.Since(startTime))
	})
}

func createRoutes () [] Route {
    return [] Route {
        { 
            Path: "/", 
            Method: "GET", 
            Handler: renderTemplate("templates/index.html", ViewData {
                Message: "Hello, World",
            }),
            Middleware: [] func (http.Handler) http.Handler { 
                logRequestTime, 
            },
        },
        { 
            Path: "/yt-summary", 
            Method: "POST", 
            Handler: http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
                fmt.Fprint(w, "summary")
            }), 
            Middleware: [] func (http.Handler) http.Handler { 
                logRequestTime, 
            },
        },
    } 
}

func setupRoutes () {
    for _, route := range createRoutes() {
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
