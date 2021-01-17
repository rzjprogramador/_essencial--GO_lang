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

	// retirando uma fatia de um array com slice que a fatia da posicao 1 ate a 3 ---vai peagr a 0, 1 ,2 menos a 3
	slice2 := array1[0:3]
	fmt.Println(slice2)

	// Alterando item do array e refletindo no slice que esta pegando a fatia deste array alterado
	array1[0] = "Posicao alterada"
	fmt.Println(slice2)
}

/*
SLICE == FATIA DE UM ARRAY
ele aponta para um array

Slice vc nao espefica o tamanho , aceita adicionar mais valores,
Os tipos tem que ser os mesmos
*/

/*
Array tipos de dados e tamnho são fixos
vc tem que especificar o tamanho e os tipos que tem que ser do mesmo tipo
Vc nao consegue adcionar mais items acima do que esta declarado
Senao especificar o Tamnho ele vira um slice
Pouco usado só se precisar espeficiar o tamanho dele na app

FUNÇÕES DE COLEÇÕES
reflect.
TypeOf função que mostra o tipo de uma variavel

append  == acrescentar
acrescentar items ao slice
*/
