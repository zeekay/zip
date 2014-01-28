package main

import (
    "log"
    "zeekay.io/zip"
)

type helloJSON struct {
    Hello string `json:"hello"`
}

func main() {
    // res.End can handle JSON too!
    zip.Get("/", func(req zip.Req, res zip.Res) {
        res.End(&helloJSON{Hello: "world!"})
    })

    // Or more explicitly:
    zip.Get("/json", func(req zip.Req, res zip.Res) {
        res.JSON(&helloJSON{Hello: "world!"})
    })

    // You can also decode JSON sent in a POST body with req.JSON():
    zip.Post("/json", func(req zip.Req, res zip.Res) {
        data := helloJSON{}
        req.JSON(&data)
        log.Println(data)
    })

    zip.Run(":1337")
}
