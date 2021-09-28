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

    m := map[string]person{
        sb.last: sb,
        mb.last: mb,
    }

    for _, v := range m {
        fmt.Printf("%s %s likes these types of ice cream:\n" ,v.first, v.last)
        for _, flavor := range v.favIcecream {
            fmt.Printf("\t%s\n", flavor)
        }
    }
}
