/// +build OMIT

package main

import (
	"fmt"
	"sync"
)

func main() {
	const workers = 100 // what if we have 1, 2, 25?

	var wg sync.WaitGroup
	m := map[int]int{}

	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go func(i int) {
			for j := 0; j < i; j++ {
				m[i]++ // Concurrent access // HL
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println(m)
}
