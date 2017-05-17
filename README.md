# slice-go
[![Build Status](https://travis-ci.org/spinlock/slice-go.svg)](https://travis-ci.org/spinlock/slice-go)

#### How to setup & install
```bash
$ mkdir -p $GOPATH/src/github.com/spinlock
$ cd $_
$ git clone https://github.com/spinlock/slice-go.git
$ make install
```

#### How to use it

```go
package main

import (
    "fmt"

    "github.com/spinlock/slice-go/unsafe2"
)

func init() {
    unsafe2.SetMaxOffheapBytes(unsafe2.MinOffheapSlice * 10)
}

func main() {
    slice := unsafe2.MakeSlice(unsafe2.MinOffheapSlice)
    fmt.Println(slice.Type())

    copy(slice.Buffer(), "hello")

    fmt.Println(string(slice.Buffer()[:5]))
    for i := 0; i < 5; i++ {
        p := slice.Slice2(i, i+1).Buffer()
        fmt.Println(string(p))
        p[0]++
    }
    fmt.Println(string(slice.Buffer()[:5]))
}
```

```bash
$ go run main.go
cgo_slice:default
hello
h
e
l
l
o
ifmmp

$ go run -tags "cgo_jemalloc" main.go
cgo_slice:jemalloc
hello
h
e
l
l
o
ifmmp

```
