package main

import (
    "fmt"
)


func main() {
    // 3 ways to write for loop
    //
    // The first way
    // x := 1
    // for x < 10 {
    //     fmt.Println(x)
    //     x++
    // }
    // The second way
    // for y := 0; y < 10; y++ {
    //     fmt.Println(y)
    // }
    // The third way
    // Alternative to WHILE loop
    // x := 1
    // for {
    //     if x > 9 {
    //         break
    //     }
    //     fmt.Println(x)
    //     x++
    // }

    // NESTING LOOP
    //
    // for i := 0; i <= 10; i++ {
    //     fmt.Printf("The outer loop: %d\n", i)
    //     for j := 0; j < 3; j++ {
    //         fmt.Printf("\tThe inner loop: %d\n", j)
    //     }
    // }

    x := 0
    for {
        x++
        if x > 100 {
            break
        }
        if x%2 != 0 {
            continue
        }
        fmt.Println(x)
    }
}
