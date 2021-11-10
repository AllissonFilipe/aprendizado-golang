package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() { // como criar um método em GO
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados\n", u.nome)
}

func (u *usuario) fazerAniversario() { //como alterar os dados no struct
	u.idade++ //pq com o ponteiro vc nao ta alterando uma copia, mas sim a estrutura
}

func main() {
	usuario1 := usuario{"Allisson", 26}
	usuario1.salvar()

	usuario2 := usuario{"David", 30}
	usuario2.salvar()

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}
