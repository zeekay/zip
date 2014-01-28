package zip

import (
    "log"
    "net/http"
)

func Run(bind string) {
    log.Println("Listening on " + bind)
    log.Fatal(http.ListenAndServe(bind, nil))
}
