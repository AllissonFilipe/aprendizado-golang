package main

import "fmt"

func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return //com retorno nomeado nao precisa criar a variavel dentro do método e nem especidicar um retonro pq ele ja sabe o que tem q retornar
}

func main() {
	soma, subtracao := calculosMatematicos(10, 20)
	fmt.Println(soma, subtracao)
}
