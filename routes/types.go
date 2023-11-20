package routes

import ( "net/http" )

type Route struct {
    Path string
    Method string
    Handler http.Handler
    Middleware [] func (http.Handler) http.Handler
}

type ViewData struct {
    Message string
}
