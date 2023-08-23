package main

import (
	"bufio" //leitura de texto
	"fmt"
	"os"
)

func main() {
	var x, z int
	var y float64

	fmt.Println("Informe dois números:")
	fmt.Scanln(&x, &y)
	fmt.Println(x)
	fmt.Println(z)

	fmt.Println("Informe um float:")
	fmt.Scanln(&y)
	fmt.Println(y)

	//lendo frases inteiras
	leitor := bufio.NewReader(os.Stdin)
	fmt.Println("Escreva um texto:")
	msg, _ := leitor.ReadString('\n')

	fmt.Println(msg)

}
