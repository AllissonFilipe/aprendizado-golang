package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("incrementando i")
		time.Sleep(time.Second)
	}

	for j := 0; j < 10; j++ {
		fmt.Println("incrementando i")
		time.Sleep(time.Second)
	}

	nomes := [3]string{"João", "Davi", "Lucas"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for indice, letra := range "STRING" {
		fmt.Println(indice, string(letra)) //sem usar o metodo "string" ele retorna os números da tabela ASC
	}

	usuario := map[string]string{
		"nome":      "Allisson",
		"sobrenome": "Filipe",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	for {
		fmt.Println("Loop Infinito")
	}
}
