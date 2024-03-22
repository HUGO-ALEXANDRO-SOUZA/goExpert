package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {

	hugo := Cliente{
		Nome:  "Hugo Alexandro Souza",
		Idade: 39,
		Ativo: true,
	}
	hugo.Ativo = false

	fmt.Println(hugo.Ativo)
	fmt.Println(hugo.Nome)
	fmt.Println(hugo.Idade)
}
