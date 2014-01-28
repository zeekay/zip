package main

import "zeekay.io/zip"

func main() {
    // You can return a string,
    zip.Get("/", func(req zip.Req, res zip.Res) {
        res.End("hello world!")
    })

    // bytes,
    zip.Get("/bytes", func(req zip.Req, res zip.Res) {
        res.End([]byte("hello bytes!"))
    })

    // Too magic? You can also be explicit:
    zip.Get("/write", func(req zip.Req, res zip.Res) {
        res.Write([]byte("hello bytes!"))
    })

    zip.Get("/writeString", func(req zip.Req, res zip.Res) {
        res.WriteString("hello bytes!")
    })

    zip.Run(":1337")
}
