package main

import (
    "fmt"
)


var x int
var y string
var z bool

func main() {
    fmt.Printf("x: %v y: %v z: %v\n", x, y, z)
    fmt.Printf("x: %T y: %T z: %T\n", x, y, z)
}
