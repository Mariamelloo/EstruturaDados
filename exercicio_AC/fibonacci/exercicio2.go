package main

import "fmt"

func main() {
	fibo()
}
func fibo() {

	var posicao int
	fmt.Println("Digite um número:")
	fmt.Scanln(&posicao)

	n := 1
	n2 := 1

	if posicao < 0 {
		fmt.Println("O resultado é 0")
	} else if posicao == 1 || posicao == 2 {
		fmt.Println("O resultado é 1")
	} else {
		for i := 3; i <= posicao; i++ {
			n, n2 = n2, n+n2
		}
		fmt.Println("O resultado é", n2)
	}
}
