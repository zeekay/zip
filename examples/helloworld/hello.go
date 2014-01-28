package main

import (
    "zeekay.io/zip"
)

func hello() string {
    return "hello world!"
}

func main() {
    zip.Get("/", hello)
    zip.Listen(":8000")
}
