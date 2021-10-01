package main

import (
    "fmt"
    "sync"
    "runtime"
)

func main() {
    var wg sync.WaitGroup
    counter := 0
    gs := 50
    wg.Add(gs)

    for i := 0; i < gs; i++ {
        go func () {
            counter++
            runtime.Gosched()
            wg.Done()
        }()
    }

    wg.Wait()
    fmt.Println(counter)
}
