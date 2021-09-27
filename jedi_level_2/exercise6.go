package main

import (
    "fmt"
)


const (
    thisYear = 2021
    a = (thisYear - iota)
    b = (thisYear - iota)
    c = (thisYear - iota)
)

func main () {
    fmt.Println(thisYear, a, b, c)
}
