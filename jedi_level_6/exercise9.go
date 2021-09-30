package main

import (
    "fmt"
)

func foo(x int) int {
    return x * x
}

func bar(f func(x int) int, ii int) {
    fmt.Println(f(ii)) 
}

func main() {
    bar(foo, 7)
}
