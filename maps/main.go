package main

import (
    "fmt"
)


func main() {
    m := map[string]int{
        "James": 32,
        "Miss Moneypenny": 27,
    }
    fmt.Println(m)
    fmt.Println(m["James"])
    fmt.Println(m["Barnabas"])
    v, ok := m["Barnabas"]
    fmt.Println(v)
    fmt.Println(ok)

    // add a new element to map
    m["todd"] = 33

    if v, ok := m["Barnabas"]; ok {
        fmt.Println("THIS IS THE IF PRINT", v)
    }

    for k, v := range m {
        fmt.Println(k, v)
    }

    xi := []int{4, 5, 6, 7, 8, 9, 42}
    for i, v := range xi {
        fmt.Println(i, v)
    }

    // delete an item from map
    delete(m, "James")
    fmt.Println(m)

    if v, ok := m["Ian Fleming"]; ok {
        fmt.Println("value:", v)
        delete(m, "Ian Fleming")
    }
    fmt.Println(m)

}
