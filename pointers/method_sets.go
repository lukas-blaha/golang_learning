package main

import (
    "fmt"
    "math"
)

type circle struct {
    radius float64
}

type shape interface {
    area() float64
}

func (c *circle) area() float64 {
    return math.Pi * math.Pow(c.radius, 2)
}

func info(s shape) {
    fmt.Println("area", s.area())
}

func main() {
    c := circle{5}
    info(&c)
}
