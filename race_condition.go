package main

import (
	"fmt"
	"time"
)

var globalSlice []int

func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			globalSlice = append(globalSlice, i)
		}(i)
	}

	fmt.Print("Global slice: ", globalSlice)
	time.Sleep(time.Second)
}
