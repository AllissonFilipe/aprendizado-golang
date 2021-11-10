package main

import (
	"errors"
	"fmt"
)

func main() {
	// int, int8, int16, int32, int64, uint -> int sem sinal
	// float32, float64
	// string
	// bool
	// error
	var numero int16 = 100

	fmt.Println(numero)

	//alias
	// INT32 = RUNE
	// INT8 = BYTE
	var numero3 rune = 12345

	fmt.Println(numero3)

	char := 'B'       // convert em numero da tabela ASC
	var numeroT int32 // valor inicial 0

	fmt.Println(char)
	fmt.Println(numeroT)

	var erro error = errors.New("Erro interno")

	fmt.Println(erro)
}
