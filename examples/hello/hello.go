package main

import (
    "zeekay.io/zip"
)

func main() {
    zip.Get("/", func(req zip.Req, res zip.Res) {
        res.Header.Set("Content-Type", "text/plain")
        res.WriteString("hello world!")
    })
    zip.Listen(":1337")
}
