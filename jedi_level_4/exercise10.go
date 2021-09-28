package main

import (
    "fmt"
)

func main() {
    m := map[string][]string{
        "bond_james": []string{"Shaken, not stirred", "Martinis", "Women"},
        "moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
        "no_dr": []string{"Being evil", "Ice cream", "Sunsets"},
    }

    m["flemin_ian"] = []string{"steaks", "cigars", "espionage"}

    if _, ok := m["moneypenny_miss"]; ok {
        delete(m, "moneypenny_miss")
    }

    for k, v := range m {
        fmt.Printf("%v:\n", k)
        for i, val := range v {
            fmt.Printf("\tindex: %d\tvalue: %v\n", i, val)
        }
    }
}
