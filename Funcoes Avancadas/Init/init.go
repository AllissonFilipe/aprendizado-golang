package main

import "fmt"

func init() {
	fmt.Println("Executando a função init") //a função init ela sempre vai ser executada primeiro que a função main
}

func main() {
	fmt.Println("Executando a função main")
}
