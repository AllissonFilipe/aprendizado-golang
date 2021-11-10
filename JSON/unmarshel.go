package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"` // se eu quiser nao incluir campos só é colocar "-"" na definicao da string de json
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	//CONVERTENDO JSON PARA STRUCT OU MAP
	cachorroEmJSON := `{"nome":"rock","raca":"vira-lata","idade":2}`
	var c cachorro

	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)

	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c2)
}
