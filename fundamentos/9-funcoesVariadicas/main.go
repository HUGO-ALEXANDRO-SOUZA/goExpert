package main

import "fmt"

func main() {
	fmt.Println(sum(1, 10, 25, 84, 54, 84, 66, 21, 74))
}

func sum(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}
	return total
}
