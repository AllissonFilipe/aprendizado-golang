package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) { // em go uma função pode ter mais de um retorno
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main() {
	fmt.Println(somar(1, 3))

	var f = func() {
		fmt.Println("funcao teste")
	}

	f()

	resultadoSoma, resultadoSubtracao := calculosMatematicos(3, 1)
	// resultadoSoma, _ := calculosMatematicos(3, 1) -> quando só quer o primeiro valor

	fmt.Println(resultadoSoma, resultadoSubtracao)
}
