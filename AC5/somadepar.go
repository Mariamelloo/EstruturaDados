package main

import "fmt"

func soma_dos_pares(array []int, alvo int) (int, int) {
	num1, num2 := 0, len(array)-1

	for num1 < num2 {
		soma := array[num1] + array[num2]

		if soma == alvo {
		return array[num1], array[num2]
		} else if soma < alvo {
		num1++
		} else {
		num2--
		}
	}

	return -1, -1
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6}
	alvo := 9

	num_par1, num_par2 := soma_dos_pares(array, alvo)

	if num_par1 != -1 && num_par2 != -1 {
	fmt.Printf("Par encontrado: (%d, %d)\n", num_par1, num_par2)
	} else {
	fmt.Println("Nenhum par encontrado.")
	}
}