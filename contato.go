package main

import "fmt"


type Contato struct {
	Nome  string
	Email string
}

func (c * Contato) Altera(novoEmail string){
	c.Email = novoEmail
}

func adicionaContato(contato, a [4]Contato) [4]Contato{
	contatos[0] = c
	for _, c := range [4]Contato{
		if c == ""{
			contatos[] = contato
			break
		}
	}
}

func main(){
	contato1 := Contato{
		Nome:	"Maria",
		Email:	"mariamello@gmail.com",
	}

	fmt.Println(contato1)

	var contatos [4]Contato
	fmt.Println(contatos)

}

