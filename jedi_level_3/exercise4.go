package main

import (
    "fmt"
)


func main() {
    born, currentYear := 1994, 2021
    year := born

    for {
        if year <= currentYear {
            fmt.Println(year)
        } else {
            break
        }
        year++
    }
}
