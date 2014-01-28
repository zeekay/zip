# zip
A simple and fast web framework for go.

## Install
Zip can be installed with `go get`:

```bash
$ go get zeekay.io/zip
```

## Usage
A simple hello world web application would look something like this:

```go
// hello.go

package main

import (
    "zeekay.io/zip"
)

func main() {
    zip.Get("/", func(req zip.Req, res zip.Res) {
        res.WriteString("hello world!")
    })
    zip.Listen(":1337")
}
```

You can run the server with `go run`:

```bash
$ go run hello.go
```

## Examples
There are a few examples in [examples/][examples] for you to play with:

- [hello.go][hello.go]
- [json.go][json.go]
- [websocket.go][websocket.go]

[examples]:     https://github.com/zeekay/zip/blob/master/examples
[hello.go]:     https://github.com/zeekay/zip/blob/master/examples/hello/hello.go
[json.go]:      https://github.com/zeekay/zip/blob/master/examples/json/json.go
[websocket.go]: https://github.com/zeekay/zip/blob/master/examples/websocket/websocket.go
