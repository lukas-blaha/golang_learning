package main

import (
    "fmt"
)


func main() {
    // SLICE allows you to group together VALUES of the same TYPE
    // x := type{values} // COMPOSITE LITERAL
    x := []int{4, 5, 7, 8, 42}
    fmt.Println(len(x))

    // for i, v := range x {
    //     fmt.Println(i, v)
    // }
    fmt.Println(x[:3])
    fmt.Println(x[1:])
    fmt.Println(x[1:4])

    // append function
    x = append(x, 77, 88, 99, 1014)
    fmt.Println(x)

    y := []int{234, 456, 678, 987}
    x = append(x, y...)
    fmt.Println(x)

    x = append(x[:2], x[4:]...)
    fmt.Println(x)

    // make function
    m := make([]int, 10, 12)
    fmt.Println(m)
    fmt.Println(len(m))
    fmt.Println(cap(m))
    m[0] = 42
    m[9] = 999
    // m[10] = 3423 // this won't work, use append instead
    m = append(m, 3423, 3424)
    fmt.Println(m)
    fmt.Println(len(m))
    fmt.Println(cap(m))

    m = append(m, 3425)
    fmt.Println(m)
    fmt.Println(len(m))
    fmt.Println(cap(m))

    //multi-dimension slices
    jb := []string{"James", "Bond", "chocolate", "martini"}
    fmt.Println(jb)
    mp := []string{"Miss", "Moneypenny", "strawberry", "hazelnut"}
    fmt.Println(mp)

    xp := [][]string{jb, mp}
    fmt.Println(xp[1][2])
    fmt.Println(xp)
}



