package main

import (
	"AC3/contato"
	"AC3/operacoes"
	"fmt"
)

type Contato struct {
	Nome  string
	Email string
}

func main() {

	var contatos [5]Contato
	c1 := Contato{
		Nome:  "Maria",
		Email: "maria@gmail.com",
	}

	fmt.Println(contatos)
	contato.Alterar(1)
	operacoes.AlterarEmail("maria29@gmail.com")
	operacoes.AdicionaContato(c1, &contatos)
	operacoes.ExcluiContato(c1)

}
