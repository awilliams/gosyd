package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)
	go func() {
		for i := 0; ; i++ {
			messages <- fmt.Sprintf("%d\tHola mundo", i)
			time.Sleep(time.Second)
		}
	}()
	for msg := range messages {
		fmt.Println(msg)
	}
}
