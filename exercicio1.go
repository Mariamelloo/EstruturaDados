package main

import "fmt"

func main() {
	e_primo()
}

func e_primo() {
	var n = 30
	x := 0
	for i := 2; i < n; i++ {
		if n%i == 0 {
			x++
			fmt.Println("Multiplo de", i)

			break
		} else {
			fmt.Println("O número é primo")
		}

	}

}
