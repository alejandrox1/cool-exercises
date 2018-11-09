// Synchronized thread-safe buffer.
package main

import (
    "log"
    "sync"
    "time"
)

var (
    safe   bool
    start  time.Time
    group  *sync.WaitGroup
    buffer *Buffer
    alphas []string
)

// Buffer is used to concatenate strings together. Note that concatenation is
// associative but not commutative.
type Buffer struct {
    lock sync.Mutex
    buffer string
}

// String returns the buffer value.
func (b *Buffer) String() string {
    return b.buffer
}

// Concat is an thread-unsage operation.
func (b *Buffer) Concat(s string) {
    time.Sleep(1 * time.Second)
    b.buffer += s
    log.Println(b.buffer)
}

// SafeConcat is a thread-safe operation.
func (b *Buffer) SafeConcat(s string) {
    b.lock.Lock()
    defer b.lock.Unlock()

    b.Concat(s)
}


func write(idx int, safe bool) {
    defer group.Done()

    if idx >= len(alphas) {
        return
    }

    if safe {
        buffer.SafeConcat(alphas[idx])
    } else {
        buffer.Concat(alphas[idx])
    }
}
