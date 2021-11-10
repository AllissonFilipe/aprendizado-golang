package main

import "fmt"

func main() {
	if numero := 10; numero > 5 { // o if do GO se assemelha as demais linguagens
		fmt.Println("Numero é maior que 5") // o diferencial é o if init que permite criar uma variável no if
	} else { // mas lembrando que essa variavel só é visivel no escopo do if
		fmt.Println("Numero é menor ou igual a 5")
	}

	//SWITCH
	//se assemelha ao javascript, nao tem o brek e tem uma nova propriedade opcional
	//ela se chama fallthrough que faz o caso satisfeito reproduzir o proximo caso automaticamente
	//e a variavel nao precisa ser declarada so no switch, mas tambem no case, existem essas duas maneiras
}
