package main

import "fmt"

func main() {
	semana()
}

func semana() {
	var dia int
	fmt.Println("Digite um número:")
	fmt.Scanln(&dia)

	switch dia {
	case 1:
		fmt.Println("Domingo")
	case 2:
		fmt.Println("Seegunda-feira")
	case 3:
		fmt.Println("Terça-feira")
	case 4:
		fmt.Println("Quarta-feira")
	case 5:
		fmt.Println("Quinta-feira")
	case 6:
		fmt.Println("Sexta-feira")
	case 7:
		fmt.Println("Sábado")
	default:
		fmt.Println("Valor inválido")
	}
}
