package main

import (
    "fmt"
    "sort"
)

func main() {
    xi := []int{2,5,8,3,1,6,0,8}
    xs := []string{"Stanicka", "Miska", "Lukas", "Max", "Fik" }
    fmt.Println(xi)
    fmt.Println(xs)
    sort.Ints(xi)
    sort.Strings(xs)
    fmt.Println(xi)
    fmt.Println(xs)
}
