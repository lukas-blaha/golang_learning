package main

import (
    "fmt"
)

func main() {
    s1 := foo()
    fmt.Println(s1)

    x := bar()
    fmt.Printf("%T\n", x)
    fmt.Println(x())
}

// returning string
func foo() string {
    s := "Hello world"
    return s 
}

// returning func int
func bar() func() int {
    return func() int {
        return 451
    }
}
