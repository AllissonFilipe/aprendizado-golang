package main

import "fmt"

type usuario struct {
	nome     string
	idade    int8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("teste")
	var u usuario
	u.nome = "Alisson"
	u.idade = 26
	fmt.Println(u)

	enderecoExemplo := endereco{"rua quarenta", 166}

	u2 := usuario{"carlos", 21, enderecoExemplo}
	fmt.Println(u2)

	u3 := usuario{idade: 27}
	fmt.Println(u3)
}
