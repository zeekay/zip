# zip
A zippy web micro-framwork. Just a tiny bit of sugar for Go's excellent
`net/http` package. Inspired by [bottle][bottle], [express][express], [web.go][web.go], etc.

## Install
Zip can be installed with `go get`:

```bash
$ go get zeekay.io/zip
```

## Usage
A simple hello world web app looks like this:

```go
package main

import "zeekay.io/zip"

func main() {
    zip.Get("/", func(req zip.Req, res zip.Res) {
        res.End("hello world!")
    })

    // Run and listen on port 1337.
    zip.Run(":1337")
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
[bottle]:       http://bottlepy.org
[express]:      http://expressjs.com
[web.go]:       http://webgo.io
