package main

import (
    "fmt"
)

type person struct {
    name    string
}

type human interface {
    speak()
}

func (p *person) speak() {
    fmt.Println("hello, I am", p.name)
}

func saySomething(h human) {
    h.speak()
}

func main() {
    p1 := person {
        name: "James Bond",
    }

    p1.speak()
    saySomething(&p1)
    // saySomething(p1) --> this won't work
    
}
