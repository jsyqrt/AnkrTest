package main

import (
	"fmt"
	"sync"
)

func main() {

	wakeA := make(chan struct{})
	wakeB := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(100)

	go func(start int) {
		for {
			select {
			case <-wakeA:
				fmt.Println(start)
				start += 2
				wakeB <- struct{}{}

				wg.Done()
			}
		}
	}(1)

	go func(start int) {
		for {
			select {
			case <-wakeB:
				fmt.Println(start)
				start += 2
				wakeA <- struct{}{}

				wg.Done()
			}
		}
	}(2)

	wakeA <- struct{}{}

	wg.Wait()
}
