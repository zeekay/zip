package zip

import (
    "log"
    "net/http"
)

func Listen(bind string) {
    log.Println("Listening on " + bind)
    log.Fatal(http.ListenAndServe(bind, nil))
}
