package main

import "fmt"

func main() {
	e_primo()

}

func e_primo() {
	var n = 10
	x := 0

	if n <= 1 {
		fmt.Println("O número não é primo")
		return
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			x++
			fmt.Println("Múltiplo de", i)
		}
	}

	if x > 0 {
		fmt.Println("O número não é primo")
	} else {
		fmt.Println("O número é primo")
	}
}
