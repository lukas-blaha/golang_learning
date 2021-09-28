package main

import (
    "fmt"
)


type vehicle struct {
    doors int
    color string
}

type truck struct {
    vehicle
    fourWheel bool
}

type sedan struct {
    vehicle
    luxury bool
}

func main() {
    daf := truck {
        vehicle: vehicle {
            doors: 2,
            color: "white",
        },
        fourWheel: true,
    }

    octavia := sedan {
        vehicle: vehicle {
            doors: 4,
            color: "black",
        },
        luxury: false,
    }

    fmt.Println(daf)
    fmt.Println(octavia)

    fmt.Printf("%T:\n  doors: %d\n  color: %s\n  4x4: %v\n", daf, daf.doors, daf.color, daf.fourWheel)
    fmt.Printf("%T:\n  doors: %d\n  color: %s\n  luxury: %v\n", octavia, octavia.doors, octavia.color, octavia.luxury)
}
