package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"_"`
	Saldo  int `json:"saldo"`
}

func main() {

	//Criando o Json e inserindo na variavel res
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}
	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		println(err)
	}

	//Crindo o json puro
	jsonPuro := []byte(`{"n":2, "s":200}`)
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		println(err)
	}
	println(contaX.Saldo)

}
