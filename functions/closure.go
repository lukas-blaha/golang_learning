package main

import (
    "fmt"
)

var x int

func main() {
    fmt.Println(x)
    x++
    fmt.Println(x)

    {
        y := 42
        fmt.Println(y)
    }
    // fmt.Println(y) --> this won't work, "y" is defined in a block of code above
    // and it is only in scope of the block
}

func foo() func() int {
    var x int
    return func() int {
        x++
        return x
    }
}
