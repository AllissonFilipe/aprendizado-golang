package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"usuario": {
			"nome":      "Allisson",
			"Sobrenome": "Filipe",
		},
		"curso": {
			"nome":   "nabuco",
			"campus": "paulista",
		},
	}

	fmt.Println(usuario2)

	delete(usuario2, "nome")

	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "capricornio",
	}

	fmt.Println(usuario2)
}
