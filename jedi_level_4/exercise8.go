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

    for k, v := range m {
        fmt.Printf("%v:\n", k)
        for i, val := range v {
            fmt.Printf("\tindex: %d\tvalue: %v\n", i, val)
        }
    }
}
