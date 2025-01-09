package main

import (
	"fmt"
	"sync"
)

func case3(channels ...chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, channel := range channels {
		go func() {
			defer wg.Done()
			for val := range channel {
				out <- val
			}
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	ch4 := make(chan int)

	go func() {
		ch1 <- 1
		close(ch1)
	}()
	go func() {
		ch2 <- 2
		close(ch2)
	}()
	go func() {
		ch3 <- 3
		close(ch3)
	}()
	go func() {
		ch4 <- 4
		close(ch4)
	}()

	chs := case3(ch1, ch2, ch3, ch4)
	for c := range chs {
		fmt.Println(c)
	}

}
