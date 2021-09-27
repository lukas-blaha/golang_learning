package main

import "fmt"


var a int
type hotdog int
var b hotdog

func main() {
    a = 42
    b = 43
    fmt.Printf("var a of type %T and value %v\n", a, a)
    fmt.Printf("var a of type %T and value %v\n", b, b)
    a = int(b)
    fmt.Printf("var a of type %T and value %v\n", a, a)
}
