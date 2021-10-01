package main

import (
    "fmt"
)

func main() {
    cs := make(chan<- int) // send only
    cr := make(<-chan int) // receive only

    cs <- 42
    cr <- 43

    fmt.Println(<-cs)
    fmt.Println(<-cr)
    fmt.Println("------")
    fmt.Printf("%T\n", cs)
    fmt.Printf("%T\n", cr)
}
