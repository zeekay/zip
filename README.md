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
```

You can run the server with `go run`:

```bash
$ go run hello.go
```

## Examples
You can checkout some examples of using Zip here: [examples/][examples]. To run
them clone this repository and use `go run`.

[examples]: https://github.com/zeekay/zip/tree/master/examples
