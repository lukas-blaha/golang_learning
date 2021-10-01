package main

import (
    "fmt"
    "runtime"
    "sync"
    "sync/atomic"
)

func main() {
    var wg sync.WaitGroup

    var counter int64
    gs := 50

    wg.Add(gs)

    for i := 0; i < gs; i++ {
        go func () {
            atomic.AddInt64(&counter, 1)
            runtime.Gosched()
            wg.Done()
        }()
    }

    wg.Wait()
    fmt.Println(counter)
}
