package zip

import (
    "io"
    "log"
    "net/http"
    "encoding/json"
)

type Req struct {
    req *http.Request
}

type Res struct {
    res http.ResponseWriter
    Header http.Header
}

func (r *Res) WriteString(s string) {
    io.WriteString(r.res, s)
}

func (r *Res) JSON(value interface{}) {
    r.res.Header().Set("Content-Type", "application/javascript")
    json.NewEncoder(r.res).Encode(value)
}

type Route func(Req, Res)

func Get(url string, handler Route) {
    http.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
        req := Req{r}
        res := Res{res: w, Header: w.Header()}
        handler(req, res)
    })
}

func Listen(bind string) {
    log.Println("Listening on " + bind)
    log.Fatal(http.ListenAndServe(bind, nil))
}
