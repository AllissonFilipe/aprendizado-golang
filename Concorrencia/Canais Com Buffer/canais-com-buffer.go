package main

import "fmt"

func main() {
	canal := make(chan string, 2) //defino a capacidade do canal

	canal <- "Olá Mundo!"
	canal <- "Prgramando em Go!"

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
