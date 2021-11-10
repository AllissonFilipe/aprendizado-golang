package main

import "fmt"

func main() {
	var var1 int = 8
	var var2 int = var1

	fmt.Println(var1, var2)

	var1++
	fmt.Println(var1, var2)

	var var3 int
	var ponteiro *int

	var3 = 10
	ponteiro = &var3 //nao pegando o valor e sim a referencia da memória

	fmt.Println(var3, ponteiro) // Ele vai printar o endereço de memória

	var3 = 150
	fmt.Println(var3, *ponteiro) // Desreferenciacao pra exibir o valor
}
