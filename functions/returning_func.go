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
    // this print does exactly the same as the print above, without assigning it
    // to var
    fmt.Println(bar()())
}

// returning string
func foo() string {
    s := "Hello world"
    return s 
}

// returning func() int
func bar() func() int {
    return func() int {
        return 451
    }
}
