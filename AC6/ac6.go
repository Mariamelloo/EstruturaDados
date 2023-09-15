package main

import (
	"fmt"
)

const M = 5

func main() {
	var lista [M]int
	var n = 0

	insereOrd(&lista, &n, 4)
	fmt.Println(lista)

	insereOrd(&lista, &n, 2)
	fmt.Println(lista)

	insereOrd(&lista, &n, 12)
	fmt.Println(lista)

	insereOrd(&lista, &n, 6)
	fmt.Println(lista)

	insereOrd(&lista, &n, 17)
	fmt.Println(lista)

	insereOrd(&lista, &n, 1)
	fmt.Println(lista)
}

func insereOrd(v *[M]int, n *int, novoValor int) {

	if *n >= M {
		fmt.Println("Overflow")
		return
	} else {
		if confere_repeticao(v, *n, novoValor) == -1 {
			fmt.Println("Tentando inserir ", novoValor)
			fmt.Println(novoValor, "não encontrado, inserindo na lista")
			i := *n
			for i > 0 && v[i-1] > novoValor {
				v[i] = v[i-1]
				i--
			}
			v[i] = novoValor
			*n++
		} else {
			fmt.Println("O numero já se encontra na lista")
		}
	}
}

func confere_repeticao(v *[M]int, n int, x int) int {
	for i := 0; i < n; i++ {
		if v[i] == x {
			return i
		}
	}

	return -1
}
