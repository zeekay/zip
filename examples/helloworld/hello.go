// hello.go

package main

import (
    "zeekay.io/zip"
)

type json struct {
    Hello string `json:"hello"`
}

func main() {
    zip.Get("/", func(req zip.Req, res zip.Res) {
        res.WriteString("hello world!")
    })

    zip.Get("/json", func(req zip.Req, res zip.Res) {
        res.JSON(&json{Hello: "world!"})
    })
    zip.Listen(":1337")
}
