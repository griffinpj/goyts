package routes

import (
    "net/http"
    "fmt"
    "time"
    "html/template"
    lib "goyts/lib"
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

func logRequestTime(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// Call the next handler in the chain
		next.ServeHTTP(w, r)

		// Log the request time
		fmt.Printf("%s request --- %s --- %v\n", r.Method, r.URL.Path, time.Since(startTime))
	})
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

func InitRoutes () [] Route {
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
                videoId := r.FormValue("video")
                err := lib.GetAudio(videoId)
                if err != nil {
                    panic(err)
                }

                fmt.Fprint(w, fmt.Sprintf("Video ID %s converted to audio", videoId))
            }), 
            Middleware: [] func (http.Handler) http.Handler { 
                logRequestTime, 
            },
        },
    } 
}
