package main

import (
    "fmt"
)


func main() {
    // xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
    // x := sum(xi...)
    x := sum()
    fmt.Println(x)
}

func sum(x ...int) int {
    fmt.Println(x)

    sum := 0
    for _, v := range x {
        sum += v
    }
    return sum
}
