package main

import "fmt"

func main() {
	//ARRAYS TEM LIMITACAO DE TAMANHO E TIPO
	//SLICES NAO TEM LIMITACAO DE TAMANHO, MAS TEM DE TIPO
	fmt.Println("Arrays e Slices")

	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slice)

	slice = append(slice, 10)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "posicao alterada"
	fmt.Println(slice2) //O VALOR EXIBIDO NO SLICE 2 VAI SER O DA MODIFICACAO PQ O SLICE ELE PEGA A REFERENCIA DO ARRAY E NAO O VALOR

	//ARRAYS INTERNOS

	fmt.Println("-----------------------------")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)

	fmt.Println(slice3)
	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) //capacidade

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4)) //VERIFICAR O TAMANHO DO SLICE
	fmt.Println(cap(slice4)) //VERIFICAR A CAPACIDADE DO SLICE
}
