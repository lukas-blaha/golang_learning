package main

import (
    "fmt"
    "golang.org/x/crypto/bcrypt"
)

func main() {
    s := `password123`
    bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
    if err != nil {
        fmt.Println(err)
    }

    loginPass := `password123`

    err = bcrypt.CompareHashAndPassword(bs, []byte(loginPass))
    if err != nil {
        fmt.Println("YOU CAN'T LOGIN!")
        return
    }
    fmt.Println("You're logged in")
}
