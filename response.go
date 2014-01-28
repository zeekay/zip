package zip

import (
    "io"
    "net/http"
    "encoding/json"
)

// Response
type Res struct {
    res http.ResponseWriter
    Header http.Header
}

func (r *Res) WriteString(s string) (int, error) {
    return io.WriteString(r.res, s)
}

func (w *Res) Write(data []byte) (int, error) {
   return w.res.Write(data)
}

func (r *Res) JSON(j interface{}) {
    r.res.Header().Set("Content-Type", "application/javascript")
    json.NewEncoder(r.res).Encode(j)
}
