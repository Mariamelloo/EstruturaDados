package operacoes

import (
	"fmt"
	"strings"
)

type Contato struct {
	Nome  string
	Email string
}

func AlterarEmail(c *Contato, novoEmail string) {
	c.Email = novoEmail
	editaEmail(&novoEmail)
}

func editaEmail(email *string) {
	*email = strings.ReplaceAll(*email, "maria.mello@gmail.com", "maria.mello2910@gmail.com")
}

func AdicionaContato(c Contato, a *[5]Contato) {
	for ind, contato := range a {
		if (contato == Contato{}) {
			a[ind] = c
			break
		}
	}

}

func ExcluiContato(a *[5]Contato) {
	if (a[0] == Contato{}) {
		fmt.Println("Lista de contatos vazia!")
		return
	}

	for ind, contato := range a {
		if (contato == Contato{}) {
			a[ind-1] = Contato{}
		}
	}

}
