package main
import (
    "fmt"
    "sync"
    "sync/atomic"
    "time"
)

var (
    x int64
    mx sync.Mutex
    wg sync.WaitGroup
)

func Add () {
    x++
    wg.Done()
}