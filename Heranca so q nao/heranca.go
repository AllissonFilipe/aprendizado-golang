package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa    //dessa maneira vc vai colocar os mesmos campos do type pessoa em estudante como herança
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"joao", "carlos", 22, 178}
	fmt.Println(p1)

	e1 := estudante{p1, "sistemas de informação", "joaquim nabuco"}
	fmt.Println(e1)
}
