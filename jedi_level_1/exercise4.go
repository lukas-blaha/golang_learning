package main

import (
    "fmt"
)

type newType int
var x newType

func main() {
    fmt.Printf("x: %v\n", x)
    fmt.Printf("type of x: %T\n", x)
    x = 42
    fmt.Printf("x: %v\n", x)
}
