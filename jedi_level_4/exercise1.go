package main

import (
    "fmt"
)


func main() {
    x := [5]int{0, 4, 10, 18, 28}
    for i, v := range x {
        fmt.Println(i, v)
    }

    fmt.Printf("%T\n", x)
}
