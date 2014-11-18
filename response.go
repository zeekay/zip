package zip

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Response
type Res struct {
	http.ResponseWriter
	res    http.ResponseWriter
	Header http.Header
}

func (r *Res) WriteString(s string) (int, error) {
	return io.WriteString(r.res, s)
}

// Encodes struct into JSON
func (r *Res) JSON(j interface{}) {
	json.NewEncoder(r.res).Encode(j)
}

// End takes string or []byte and calls appropriate method
func (r *Res) End(value interface{}) {
	switch v := value.(type) {
	default:
		fmt.Printf("unexpected type %T", v)
	case []byte:
		if b, ok := value.([]byte); ok {
			r.Write(b)
		}
	case string:
		if s, ok := value.(string); ok {
			r.WriteString(s)
		}
	case interface{}:
		r.JSON(value)
	}
}
