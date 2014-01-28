package zip

import (
    "net/http"
    "encoding/json"
)

// Request struct
type Req struct {
    *http.Request
}

func (r *Req) JSON(s interface{}) (err error) {
    decoder := json.NewDecoder(r.Body)
    return decoder.Decode(&s)
}
