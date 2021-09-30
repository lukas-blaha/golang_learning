package main

import (
    "fmt"
    "math"
)

type square struct {
    a float64
}

func (sq square) area() float64 {
    return math.Pow(sq.a, 2)
}

type circle struct {
    r float64
}

func (ci circle) area() float64 {
    return math.Pi * math.Pow(ci.r, 2)
}

type shape interface {
    area() float64
}

func info(sh shape) {
    fmt.Println(sh.area())
}

func main() {
    sq1 := square {
        a: 4,
    }

    ci1 := circle {
        r: 5,
    }

    info(sq1)
    info(ci1)
}
