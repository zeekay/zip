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

func hello() string {
    return "hello world!"
}

func main() {
    zip.Get("/", hello)
    zip.Listen(":8000")
}
```

You can run the server with `go run`:

```bash
$ go run hello.go
```
