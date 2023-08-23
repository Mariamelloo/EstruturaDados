package main

import "fmt"

var linguagem string
var soma bool

func main() {
	//declaração explicira do tipo
	var msg string
	var n1, n2 int

	//declaração implicita
	var n3, n4 = 15, 25

	//declaração curta
	msg2 := "olá, mundo"

	msg = "oi"
	n1 = 10
	n2 = 20

	fmt.Println(msg, n1, n2, n3, n4, msg2)

	const Pi = 3.14

	f2()
}

func f2() {
	linguagem = "python"
	fmt.Println(linguagem)
}

/*
inteiros

int8	-128 a 127
uint8	0 a 255
int16	-32.768 a 32.767
uint16	0 a 65.535
int32
uint32
int64
uint64

byte	uint8
rune	int32 caracteres unicode

int -> int32 ou in64, dependendo da arquitetura do computador
flutuantes

float32		precisão de 6 a 9 dígitos
float64		precisão de 15 a 17 digitos

é necessário especificar o tamanho do float
caso a declaração seja implicíta, é assumido float64
-------------

complexo

complex64
complex128

string
bool		true/ false
nil			ausencia de dados
*/
