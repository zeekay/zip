package main

import (
    "io"
    "zeekay.io/zip"
)

// Echo the data received on the WebSocket.
func main() {
    zip.WebSocket("/ws", func(ws zip.Conn) {
        io.Copy(ws, ws)
    })

    zip.Run(":1337")
}
