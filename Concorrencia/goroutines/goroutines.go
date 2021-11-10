package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Olá Mundo") // o go routine permite que o codigo continue sem ao mesmo ter sido finalizado a execução da função, executando os demais codigos de maneira concorrente
	escrever("Programando em Go!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
