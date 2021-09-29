package main

import (
    "fmt"
)


func main() {
    foo()
    bar("James")
    s1 := woo("Moneypenny")
    fmt.Println(s1)
    x, y := mouse("Ian", "Fleming")
    fmt.Println(x)
    fmt.Println(y)
}

// func (r receiver) identifier(params) (return(s)) { ... }

func foo() {
    fmt.Println("hello from foo")
}

// everything in GO is PASS BY VALUE
func bar(s string) {
    fmt.Println("hello,", s)
}

func woo(s string) string {
    return fmt.Sprint("Hello from woo, ", s)
}

func mouse(fn, ln string) (string, bool) {
    a := fmt.Sprint(fn, " ", ln, `, says "hello"`)
    b := false
    return a, b
}
