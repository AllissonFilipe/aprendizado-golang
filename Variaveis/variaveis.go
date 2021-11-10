package main

import "fmt"

func main() {
	var variavel1 string = "teste"
	variavel2 := "teste2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "teste 3"
		variavel4 string = "teste 4"
	)

	fmt.Println(variavel3, variavel4)

	variavel1, variavel2 = variavel2, variavel1

	fmt.Println(variavel1, variavel2)
}
