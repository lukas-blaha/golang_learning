package main

import (
    "fmt"
)

func main() {
    i := incrementor()
    j := incrementor()

    fmt.Println(i())
    fmt.Println(i())
    fmt.Println(i())
    fmt.Println(j())
    fmt.Println(j())
    fmt.Println(j())
}

func incrementor() func() int {
    var x int
    return func() int {
        x++
        return x
    }
}
