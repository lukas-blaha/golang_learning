package main

import (
    "fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
    s := fmt.Sprintf("x: %v y: %v z: %v\n", x, y, z)
    fmt.Print(s)
}
