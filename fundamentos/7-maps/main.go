package main

import "fmt"

func main() {

	salarios := map[string]int{"Hugo": 30000, "Mariana": 3000, "Luara": 10000}

	//fmt.Println(salarios["Hugo"])
	delete(salarios, "Hugo")
	salarios["Hugo"] = 5000
	//fmt.Println(salarios["Hu"])

	//sal := make(map[string]int)
	//sal1 := map[string]int{}
	//sal1["Hugo"] = 1000

	for nome, salario := range salarios {
		fmt.Printf("O salário de %s é %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salário é %d\n", salario)
	}

}
