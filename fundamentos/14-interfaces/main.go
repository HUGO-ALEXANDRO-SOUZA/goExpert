package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Pessoa interface {
	Desativar()
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {
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

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {

	hugo := Cliente{
		Nome:  "Hugo Alexandro Souza",
		Idade: 39,
		Ativo: true,
	}
	minhaEmpresa := Empresa{}

	Desativacao(minhaEmpresa)
	Desativacao(hugo)
}
