package main

import "fmt"

func main() {
	retorno := func(texto string) string { //exemplo de função anônima
		return fmt.Sprintf("Texto Digitado -> %s", texto)
	}("exemplo")

	fmt.Println(retorno)
}
