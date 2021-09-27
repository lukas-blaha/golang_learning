package main

import (
    "fmt"
)


func main() {
    n := 10
    m := n << 1
    fmt.Printf("%d\t%b\t%#x\n", n, n, n)
    fmt.Printf("%d\t%b\t%#x\n", m, m, m)
}
