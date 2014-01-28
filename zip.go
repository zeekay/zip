// hello.go

package zip

import (
    "log"
    "io"
    "net/http"
)

type Route func() string

func Get(url string, handler Route) {
    http.HandleFunc(url, func(res http.ResponseWriter, req *http.Request) {
        res.Header().Set("Content-Type", "text/plain")
        io.WriteString(res, handler())
    })
}

func Listen(bind string) {
    log.Println("Listening on " + bind)
    log.Fatal(http.ListenAndServe(bind, nil))
}
