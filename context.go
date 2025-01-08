package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx1 := context.Background()

	ctx2 := context.TODO()

	ctx3, cancel := context.WithCancel(context.Background())
	defer cancel()

	deadline := time.Now().Add(1 * time.Second)
	ctx4, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	ctx5, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	ctx6 := context.WithValue(context.Background(), "userID", 42)

	select {
	case <-ctx1.Done():
		println("ctx1 is done")
	case <-ctx2.Done():
		println("ctx2 is done")
	case <-ctx3.Done():
		println("ctx3 is done")
	case <-ctx4.Done():
		println("ctx4 is done")
	case <-ctx5.Done():
		println("ctx5 is done")
	case <-ctx6.Done():
		println("ctx6 is done")
	}

	fmt.Print("Context: ", ctx6.Value("userID"))
}
