package contato

import (
	"AC3/operacoes"
	"bufio"
	"fmt"
	"os"
)

type Contato struct {
	Nome  string
	Email string
}

func Alterar() {
	var contatos [5]Contato
	var nome, email, opcao string

	leitor := bufio.NewReader(os.Stdin)

	fmt.Println("Lista de contatos!")
	for {
		fmt.Print("Digite (1) para adicionar, (2) para remover ou qualquer outra coisa para sair: ")
		fmt.Scanln(&opcao)

		switch opcao {
		case "1":
			fmt.Print("Informe o seu nome: ")
			nome, _ = leitor.ReadString('\n')

			fmt.Print("Informe o seu email: ")
			fmt.Scanln(&email)

			contatos = operacoes.AdicionaContato(Contato{Nome: nome, Email: email}, contatos)
		case "2":
			contatos = operacoes.ExcluiContato(contatos)
		default:
			return
		}

		fmt.Println(contatos)
	}
}
