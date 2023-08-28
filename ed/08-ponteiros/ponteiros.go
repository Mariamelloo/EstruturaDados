package main

import (
	"fmt"
	"strings"
)

type Pessoa struct {
	Nome  string
	Email string
}

func (p *Pessoa) AlteraEmail(novoEmail string) {
	p.Email = novoEmail
}

func adicionaPessoa(p Pessoa, a *[5]Pessoa) {
	for ind, pessoa := range a {
		if (pessoa == Pessoa{}) {
			a[ind] = pessoa
			break
		}

	}
}

func main() {

	var pessoas [5]Pessoa
	p1 := Pessoa{
		Nome:  "maria",
		Email: "maria@gmail.com",
	}

	fmt.Println(pessoas)
	adicionaPessoa(p1, &pessoas)
	fmt.Println(pessoas)

	x := 5
	y := x
	fmt.Println(x, y) // 5 5

	x = 6
	fmt.Println(x, y) //6 5

	z := &x //z é um ponteiro que aponta para x
	fmt.Println(z)
	fmt.Println(x, *z) // 6 6

	var w *int
	fmt.Println(w)

	//Ponteiros de ponteiros
	a := &z
	fmt.Println(a)
	fmt.Println(*a)
	fmt.Println(**a) // 6

	mensagem := "Olá, mundo!"
	alterarMensagem(&mensagem)
}

func alterarMensagem(msg *string) {
	// strings.ReplaceAll(stringOriginal string, textoProcurar string, textoSubstituir string)
	// "mundo" -> "turma"
	*msg = strings.ReplaceAll(*msg, "mundo", "turma")

}
