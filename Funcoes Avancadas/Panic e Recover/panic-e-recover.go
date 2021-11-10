package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil { //recover ele recupera a execução do programa
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A Média é exatamente 6 !") //O panic faz parar a execução do programa, ele chama todas as funções que tem o defer, ele mata a execução
}

func main() {

}
