package main

import (
    "fmt"
)


func main() {
    if true {
        fmt.Println("001")
    }
    if false {
        fmt.Println("002")
    }
    if !true {
        fmt.Println("003")
    }
    if !false {
        fmt.Println("004")
    }

    if x := 42; x == 42 {
        fmt.Println("005")
    }

    y := 42
    if y == 40 {
        fmt.Println("our value was 40")
    } else if y == 41 {
        fmt.Println("our value was 41")
    } else {
        fmt.Println("our value was not 40 or 41")
    }
}
