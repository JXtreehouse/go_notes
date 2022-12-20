package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	s := make([]int, 0, 1000)
	for i := 0; i < 1000; i++ {
		v := i
		wg.Add(1)
		go func() {
			s = append(s, v)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("%v\n", len(s))
}
