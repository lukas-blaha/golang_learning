package main

import (
    "fmt"
)


func main() {
    favSport := "mtb"

    switch favSport {
        case "mtb":
            fmt.Println("Wow, that's awesome sport, a lot of adrenalin too...")
        case "workout":
            fmt.Println("Building muscle with your own weight is pretty cool")
    }
}
