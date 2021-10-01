package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    var mu sync.Mutex

    counter := 0
    gs := 50

    wg.Add(gs)

    for i := 0; i < gs; i++ {
        go func () {
            mu.Lock()
            counter++
            mu.Unlock()
            wg.Done()
        }()
    }

    wg.Wait()
    fmt.Println(counter)
}
