package main

import (
	"fmt"
)

// Функция с дженериками
func Print[T any](value T) {
	fmt.Println(value)
}

func main() {
	Print(42)          // int
	Print("Hello Go!") // string
	Print(3.14)        // float64
}
