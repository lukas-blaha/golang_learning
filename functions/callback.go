package main

import (
    "fmt"
)

func main() {
   // callback means assigning the function as argument to another 
   ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
   fmt.Printf("%T\n", ii)
   fmt.Printf("sum of all numbers: %d\n", sum(ii...))

   s2 := even(sum, ii...)
   fmt.Printf("%T\n", s2)
   fmt.Printf("sum of all even numbers: %d\n", s2)

   s3 := odd(sum, ii...)
   fmt.Printf("%T\n", s3)
   fmt.Printf("sum of all odd numbers: %d\n", s3)
}

func sum(xi ...int) int {
    fmt.Printf("%T\n", xi)
    total := 0

    for _, v := range xi {
        total += v
    }

    return total
}

func even(f func(xi ...int) int, vi ...int) int {
    var yi []int

    for _, v := range vi {
        if v % 2 == 0 {
            yi = append(yi, v)
        }
    }

    return f(yi...)
}

func odd(f func(xi ...int) int, vi ...int) int {
    var yi []int

    for _, v := range vi {
        if v % 2 != 0 {
            yi = append(yi, v)
        }
    }

    return f(yi...)
}
