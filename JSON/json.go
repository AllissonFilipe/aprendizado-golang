package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	//MARSHAL
	c := cachorro{"rex", "vira-lata", 2}
	fmt.Println(c)

	cachorroEmJSON, erro := json.Marshal(c) //transformar struct ou map em json
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(cachorroEmJSON)
	fmt.Println(bytes.NewBuffer(cachorroEmJSON))

	c2 := map[string]string{
		"nome": "rock",
		"raca": "vira-lata",
	}

	cachorroEmJSON2, erro2 := json.Marshal(c2) //transformar struct ou map em json
	if erro2 != nil {
		log.Fatal(erro2)
	}
	fmt.Println(cachorroEmJSON2)
	fmt.Println(bytes.NewBuffer(cachorroEmJSON2))
}
