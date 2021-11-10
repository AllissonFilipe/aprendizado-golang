package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup // fazer agrupamentos de go routines
	waitGroup.Add(2)

	go func() {
		escrever("Olá Mundo")
		waitGroup.Done() // tira um go do contador
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done()
	}()

	waitGroup.Wait() // fale pro nosso main esperar a contagem chegar em 0
	// eles nao sao tao utilizados, os que são mais utilizados sao os canais
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
