package main

import "fmt"

func main() {
	total := func() int {
		return sum(1, 3, 45, 6, 34, 654, 654, 34534, 443, 546, 22, 323, 2332) * 2
	}()
	fmt.Println(total)
}

func sum(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}
	return total
}
