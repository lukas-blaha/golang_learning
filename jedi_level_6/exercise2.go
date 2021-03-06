package main

import (
    "fmt"
)

func foo(x ...int) int {
    sum := 0
    for _, v := range x {
        sum += v
    }
    return sum
}

func bar(x []int) int {
    sum := 0
    for _, v := range x {
        sum += v
    }
    return sum
}

func main() {
    xi := []int{2, 3, 4, 5, 6, 7}
    fmt.Println(foo(xi...))
    fmt.Println(bar(xi))
}
