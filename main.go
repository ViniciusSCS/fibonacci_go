package main

import (
	"fibonacci/fibonacci"
	"fmt"
)

func main() {
	fmt.Println("Fibonacci")

	posicao := uint(12)

	for i := uint(0); i <= posicao; i++ {
		fmt.Println(fibonacci.Fibonacci(i))
	}
}
