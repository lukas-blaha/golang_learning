package main

import (
    "fmt"
)

type person struct {
    first string
    last string
}

func changeMe(p *person) {
    p.last = "Banana"
    // (*p).last = "Banana"
}

func main() {
    bond := person {
        first: "James",
        last: "Bond",
    }
    fmt.Println(bond)
    changeMe(&bond)
    fmt.Println(bond)
}
