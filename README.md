# zip
A simple and fast web framework for go.


## Getting started
This is a simple hello world web application:

```go
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

## Run it
Put the above code in a package called `hello.go`. To run it, type:

```bash
$ go run hello.go
```
