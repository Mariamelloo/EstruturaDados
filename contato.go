package main

//"bufio"
//"os"

type Contato struct {
	Nome  string
	Email string
}

func (c Contato) Altera(novoEmail string) {
	c.Email = novoEmail
}

func main() {
	contato1 := Contato{
		Nome:  "Maria",
		Email: "mariamello@gmail.com",
	}

}
