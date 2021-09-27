package main

import (
    "fmt"
)

func main() {
    a := 3
    b := 7
    g := (a == b)
    h := (a <= b)
    i := (a >= b)
    j := (a != b)
    k := (a < b)
    l := (a > b)

    fmt.Printf("g: %v\nh: %v\ni: %v\nj: %v\nk: %v\nl: %v\n", g, h, i, j, k, l)
}
