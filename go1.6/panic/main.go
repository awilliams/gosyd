package main

import (
	"sync"
	"time"
)

func main() {
	const N = 100
	var wg sync.WaitGroup

	for i := 0; i <= N; i++ {
		wg.Add(1)
		go func(i int) {
			if i == N {
				panic("💩💀💩💀") // HL
			}
			time.Sleep(time.Hour)
		}(i)
	}

	wg.Wait()
}
