package main

import (
    "log"
    "zeekay.io/zip"
)

type helloJson struct {
    Hello string `json:"hello"`
}

func main() {
    // You can easily return JSON data with res.JSON
    zip.Get("/json", func(req zip.Req, res zip.Res) {
        res.JSON(&helloJson{Hello: "world!"})
    })

    // ...and decode JSON sent in a POST body with req.JSON(),
    // for example:
    // $ curl 127.0.0.1:1337/json -d '{"hello": "world"}'
    zip.Post("/json", func(req zip.Req, res zip.Res) {
        data := helloJson{}
        req.JSON(&data)

        // ...or so something more interesting with it.
        log.Println(data)
    })
    zip.Listen(":1337")
}
