package main

import (
    "fmt"
)

type newType int
var x newType
var y int

func main() {
    fmt.Printf("x: %v\n", x)
    fmt.Printf("type of x: %T\n", x)
    x = 42
    fmt.Printf("x: %v\n", x)
    y = int(x)
    fmt.Printf("y: %v\n", y)
    fmt.Printf("type of y: %T\n", y)
}
