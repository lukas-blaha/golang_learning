package main

import (
    "fmt"
)

func foo() int {
    return 1
}

func bar() (int, string) {
    return 2, "baaar"
}

func main() {
    fmt.Println(foo())
    fmt.Println(bar())
}
