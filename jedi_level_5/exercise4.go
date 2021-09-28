package main

import (
    "fmt"
)


func main() {
    p1 := struct {
        first string
        friends map[string]int
        favDrinks []string
    } {
        first: "James",
        friends: map[string]int {
            "Moneypenny": 555,
            "Q": 777,
            "M": 888,
        },
        favDrinks: []string{
            "Martini",
            "Water",
        },
    }

    fmt.Printf("%s:\n  friends:\n", p1.first)

    for k, v := range p1.friends {
        fmt.Printf("    %s:\n      name: %s\n      number: %d\n", k, k, v) 
    }

    fmt.Printf("  favDrinks:\n")
    for _, v := range p1.favDrinks {
        fmt.Printf("    - %s\n", v)
    }
}
