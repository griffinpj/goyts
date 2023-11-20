package routes

import ( "net/http" )

func InitRoutes () [] Route {
    return [] Route {
        { 
            Path: "/", 
            Method: "GET", 
            Handler: RenderTemplate("templates/index.html", ViewData {
                Message: "GOYTS",
            }),
            Middleware: [] func (http.Handler) http.Handler { 
                LogRequestTime, 
            },
        },
        { 
            Path: "/yt-summary", 
            Method: "POST", 
            Handler: http.HandlerFunc(YTSummary), 
            Middleware: [] func (http.Handler) http.Handler { 
                LogRequestTime, 
            },
        },
    } 
}
