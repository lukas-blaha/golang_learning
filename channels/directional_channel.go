package main

import (
    "fmt"
)

func main() {
    c := make(chan <- int, 2)

    c <- 42

    fmt.Println(<-c)
}
