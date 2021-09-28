package main

import (
    "fmt"
)

func main() {
    x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
    fmt.Printf("%v\n%v\n%v\n%v\n", x[:5], x[5:], x[2:7], x[1:6])

}
