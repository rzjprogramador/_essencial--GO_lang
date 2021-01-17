package main

import (
	"fmt"
	"reflect"
)

func main() {
	slice := []int{10, 15, 11, 28, 03}
	fmt.Println(slice)

	// Analisando qual o tipo deste slice
	fmt.Println(reflect.TypeOf(slice))
	//resp: []int  --é um slice de int

	slice = append(slice, 123)
	fmt.Println(slice)

	//--------------------------
	array1 := [5]string{"Posicao1", "Posicao2", "Posicao3", "Posicao4", "Posicao5"}
	fmt.Println(array1)

	// retirando uma fatia de um array com slice que a fatia da posicao 1 ate a 3 ---vai pegar a 0, 1 ,2 menos a 3
	slice2 := array1[0:3]
	fmt.Println(slice2)

	// Alterando item do array e refletindo no slice que esta pegando a fatia deste array alterado
	array1[0] = "Posicao alterada"
	fmt.Println(slice2)

	//============================================
	/*ARRAYS INTERNOS <<>> COMO É CONSTRUIDO O SLICE :
	nomeSlice := make([])executar um array -tipo, quantidade-inicial, capacidade

	VC NEM PRECISA PASSA A CAPACIDADE ELE JA ADEQUA PROPORCIONALMENTE PRA VC
	DE ACORDO COM O QUE VC PASSA NA QUANTIDADE
	O GO ADEQUA A QUANTIDADE COM O QUE VC PASSA E DOBRA A CAPACIDADE PROPORCIONAL SE BASEANDO NA QUANTIDADE Q VC PASSOU
	ex: se vc passou QUE VAI USAR 6 POSICOES O GO TE DA 6 POSIÇÕES E 12 DE CAPACIDADE
	ELE CRIA UM NOVO SLICE E DOBRA A CAPACIDADE DO ARRAY PARA VC
	POR ISSO QUE O SLICE NUNCA ACABA A SUA CAPACIDADE

	OBS

	*/

	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)

	fmt.Println(len(slice3)) // len == length -- total usado
	fmt.Println(cap(slice3)) // cap == capacidade

	//------------------------------
	fmt.Println("=============================")
	slice4 := make([]float32, 5) // aqui defini que tera 5 posicoes
	fmt.Println(slice4)          // aqui mostrou as 5 posicoes com valor inicial 0

	slice4 = append(slice4, 9) // adicionei com append( no slice4 um item que é o numero 200)
	fmt.Println(slice4)        // mostrou o slice com uma posicao a mais e o que foi adicionado

	fmt.Println(len(slice4)) // mostra o tamanho usado no slice
	fmt.Println(cap(slice4)) // mostra a capacidade agora do slice que foi dobrada proporcionalmente

}

/*
SLICE == FATIA DE UM ARRAY
ele aponta para um array

Slice vc nao espefica o tamanho , ele aceita adicionar mais valores,
Os tipos tem que ser os mesmos
*/

/*
ARRAY DEPRECIADO NO GO -- PROCURE NAO UTILIZAR --É LIMITADO EM TIPOS E QUANTIDADE DE CAPACIDADE
SUPERADO POR SLICE

Sabendo mais sobre arrays ===
Array tipos de dados e tamanho são fixos
vc tem que especificar o tamanho e os tipos que tem que ser do mesmo tipo
Vc nao consegue adcionar mais items acima do que esta declarado
Senao especificar o Tamnho ele vira um slice
Pouco usado só se precisar espeficiar o tamanho dele na app

*/
