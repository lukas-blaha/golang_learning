package main

import (
    "fmt"
)

func foo() func() int {
    return func() int {
        return 7
    }
}

func main() {
    x := func () func() {
        return func() {
            fmt.Println("test")
        }
    }
    x()()

    f := foo()
    fmt.Println(f())
}
