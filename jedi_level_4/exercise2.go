package main

import (
    "fmt"
)

func main() {
    x := []int{0, 4, 8, 12, 16, 20, 24, 28, 32, 36}

    for i, v := range x {
        fmt.Println(i, v)
    }
    fmt.Printf("%T\n", x)
}
