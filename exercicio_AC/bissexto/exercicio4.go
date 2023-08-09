package main

import "fmt"

func main() {
	e_bissexto()
}

func e_bissexto() {
	var ano int
	fmt.Println("Digite um ano:")
	fmt.Scanln(&ano)

	if (ano%4 == 0 && ano%100 != 0) || ano%400 == 0 {
		fmt.Println("É um ano bissexto")
	} else {
		fmt.Println("Não é um ano bissexto")
	}
}
