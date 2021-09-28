package main

import (
    "fmt"
)


func main() {
    foo()
}

// func (r receiver) identifier(params) (return(s)) { ... }

func foo() {
    fmt.Println("hello from foo")
}
