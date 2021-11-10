package main

import "fmt"

func funcao1() {
	fmt.Println("Executando função 1")
}

func funcao2() {
	fmt.Println("Executando funcao2")
}

func main() {
	defer funcao1() // ele adia a execução da função para o final antes de terminar a execução no seu escopo
	funcao2()
}
