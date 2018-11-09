package main

import (
    "log"
    "sync"
    "testing"
    "time"
)

func concatenateStrings(safe bool) {
    group = new(sync.WaitGroup)
    alphas = []string{"a", "b", "c", "d", "e", "f", "g", "h", "j", "k"}
    buffer = &Buffer{}
    start  = time.Now()

    for i := 0; i < len(alphas); i++ {
        group.Add(1)
        go write(i, safe)
    }
    group.Wait()

    log.Printf("Result: %s in %s (safe=%t)\n", buffer, time.Since(start), safe)
}


func TestBuffer(t *testing.T) {
    concatenateStrings(false)

    concatenateStrings(true)
}
