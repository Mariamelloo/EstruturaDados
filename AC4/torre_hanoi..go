package main

import "fmt"

func main() {
    hanoi(3, "A", "B", "C")
}

func hanoi(n int, origem, destino, trabalho string) {
    if n>0{
        hanoi(n-1, origem, trabalho, destino)
        fmt.Printf("Mova o disco %d de %s para %s\n", n, origem, destino)
        hanoi(n-1, trabalho, destino, origem)
    }
    

}


