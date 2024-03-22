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

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado ", c.Nome)
}

func main() {

	hugo := Cliente{
		Nome:  "Hugo Alexandro Souza",
		Idade: 39,
		Ativo: true,
	}
	hugo.Ativo = false
	hugo.Desativar()

	fmt.Println(hugo.Ativo)
	fmt.Println(hugo.Nome)
	fmt.Println(hugo.Idade)
}
