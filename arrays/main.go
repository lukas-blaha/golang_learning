package main

import (
    "fmt"
)


func main() {
    // DO NOT USE ARRAYs, USE SLICEs INSTEAD
    var x [5]int
    fmt.Println(x)
    x[4] = 42
    fmt.Println(x)
    fmt.Println(len(x))
}
