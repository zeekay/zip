package zip

import (
    "net/http"
)

type Handler func(Req, Res)

var (
    handlers = make(map[string]Handler)
)

func AddHandler(url, method string, handler Handler) {
    // Add to handlers
    handlers[method] = handler

    http.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
        // Create request
        req := Req{r}

        // Create response
        header := w.Header()
        res := Res{res: w, Header: header}

        // Set header
        header.Set("Server", "zip.go")

        // handlers added with Route() get all requests
        if method == "" { handler(req, res); return }

        // method handlers bail if method doesn't match
        if handler, ok := handlers[r.Method]; ok {
            handler(req, res)
        } else {
            w.WriteHeader(http.StatusMethodNotAllowed)
        }
    })
}

func Route(url string, handler Handler) {
    AddHandler(url, "", handler)
}

func Head(url string, handler Handler) {
    AddHandler(url, "HEAD", handler)
}

func Get(url string, handler Handler) {
    AddHandler(url, "GET", handler)
}

func Patch(url string, handler Handler) {
    AddHandler(url, "PATCH", handler)
}

func Post(url string, handler Handler) {
    AddHandler(url, "POST", handler)
}

func Put(url string, handler Handler) {
    AddHandler(url, "PUT", handler)
}

func Delete(url string, handler Handler) {
    AddHandler(url, "DELETE", handler)
}

func Options(url string, handler Handler) {
    AddHandler(url, "OPTIONS", handler)
}
