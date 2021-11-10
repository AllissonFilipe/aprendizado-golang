package main

import "fmt"

func soma(numeros ...int) int { //numero dinamico de parametros para a função
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func main() {
	somaTotal := soma(1, 4, 5, 8, 123, 578, 32)
	fmt.Println(somaTotal)
}
