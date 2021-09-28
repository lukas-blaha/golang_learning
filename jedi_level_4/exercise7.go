package main

import (
    "fmt"
)


func main() {
    jb := []string{"James", "Bond", "Shaken, not stirred"}
    mp := []string{"Miss", "Moneypenny", "Hellooooo, James"}
    x := [][]string{jb, mp}

    for i, xs := range x {
        fmt.Println("Record:", i)
        for j, v := range xs {
            fmt.Printf("\tindex: %d\t value: %v\n", j, v)
        }
    }
}
