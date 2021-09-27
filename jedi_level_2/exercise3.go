package main

import (
    "fmt"
)


const (
    a int = 42
    b float64 = 36.95
    c string = "something"
)

const (
    d = 42
    e = 36.95
    f = "something"
)

func main() {
    fmt.Printf("a: %v\t%T\nb: %v\t%T\nc: %v\t%T\n", a, a, b, b, c, c)
    fmt.Printf("d: %v\t%T\ne: %v\t%T\nf: %v\t%T\n", d, d, e, e, f, f)
}
