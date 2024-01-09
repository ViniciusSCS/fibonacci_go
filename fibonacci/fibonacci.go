package fibonacci

func Fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return Fibonacci(posicao-2) + Fibonacci(posicao-1)
}
