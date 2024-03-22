package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func main() {

	endereco := Endereco{
		Logradouro: "Praca Assis Chateaubriand",
		Numero:     54,
		Cidade:     "São José dos Campos",
		Estado:     "São Paulo",
	}

	hugo := Cliente{
		Nome:  "Hugo Alexandro Souza",
		Idade: 39,
		Ativo: true,
	}
	hugo.Ativo = false

	fmt.Println(hugo.Ativo)
	fmt.Println(hugo.Nome)
	fmt.Println(hugo.Idade)
	fmt.Println(endereco)
}
