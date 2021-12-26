package main

import "fmt"

func main() {
	for i := 0; i <= 7; i++ {
		fmt.Print(FibonacciRecursion(i))
	}
}

func FibonacciRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}
