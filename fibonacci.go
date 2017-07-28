package main

import "fmt"

func main() {
	var r = fibonacci(34)
	fmt.Println("fibonacci: ", r)
}

func fibonacci(i int) int {
	if i < 2 {
		return i
	}
	return fibonacci(i-2) + fibonacci(i-1)
}
