package main

import (
    "fmt"
)

func main() {
    func (s string) {
        text := s
        fmt.Println(text)
    }("Hello world")
    // fmt.Println(text)
}
