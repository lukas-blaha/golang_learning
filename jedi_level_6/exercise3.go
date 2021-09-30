package main

import (
    "fmt"
)

func foo() {
    fmt.Println("hello from foo")
}

func bar() {
    fmt.Println("hello from bar")
}

func main() {
    defer foo()
    bar()
}
