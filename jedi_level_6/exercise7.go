package main

import (
    "fmt"
)

func main() {
    x := func () int {
        return 7
    }

    fmt.Println(x())
}
