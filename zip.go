package zip

import (
    "io"
    "log"
    "net/http"
    "code.google.com/p/go.net/websocket"
    "encoding/json"
)

type Req struct {
    req *http.Request
}

type Res struct {
    res http.ResponseWriter
    Header http.Header
}

type Handler func(Req, Res)

type WebSocketHandler func(*websocket.Conn)

var (
    handlers = make(map[string]Handler)
)

func (r *Res) WriteString(s string) (int, error) {
    return io.WriteString(r.res, s)
}

func (w *Res) Write(data []byte) (int, error) {
   return w.res.Write(data)
}

func (r *Res) JSON(value interface{}) {
    r.res.Header().Set("Content-Type", "application/javascript")
    json.NewEncoder(r.res).Encode(value)
}

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

func WebSocket(url string, handler WebSocketHandler) {
  http.Handle(url, websocket.Handler(handler))
}

func Listen(bind string) {
    log.Println("Listening on " + bind)
    log.Fatal(http.ListenAndServe(bind, nil))
}
