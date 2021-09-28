package main

import (
    "fmt"
)


type person struct {
    first string
    last string
    favIcecream []string
}

func main() {
    sb := person{
        first: "Stanicka",
        last: "Blahova",
        favIcecream: []string{
            "lemon",
            "strawberry",
            "vanilla",
        },
    }

    mb := person{
        first: "Miska",
        last: "Blaha",
        favIcecream: []string{
            "vanilla",
            "cherries",
            "caramel",
        },
    }

    fmt.Printf("%s %s likes these types of ice cream:\n", sb.first, sb.last)
    for _, v := range sb.favIcecream {
        fmt.Printf("\t%s\n", v)
    }

    fmt.Printf("%s %s likes these types of ice cream:\n", mb.first, mb.last)
    for _, v := range mb.favIcecream {
        fmt.Printf("\t%s\n", v)
    }
}
