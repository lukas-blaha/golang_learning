package main

import (
    "fmt"
    "encoding/json"
)

type person struct {
    First string `json:"First"`
    Last string `json:"Last"`
    Age int `json:"Age"`
}

func main() {
    s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`
    bs := []byte(s)
    fmt.Printf("%T\n", s)
    fmt.Printf("%T\n", bs)

    var people []person
    err := json.Unmarshal(bs, &people)
    if err != nil {
        fmt.Println(err)
    }

    for _, v := range people {
        fmt.Println(v.First, v.Last, "is", v.Age, "years old\n")
    }
}
