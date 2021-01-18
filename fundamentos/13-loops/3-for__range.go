package main

import "fmt"

func main() {

	nomes := []string{"Reinaldo", "Gustavo", "Leonardo"}

	// Interando por Slice
	for _, nome := range nomes {
		//fmt.Println(indice, nome)
		fmt.Println(nome)
	}

	//Interando por string
	for i, letra := range "REINALDO" {
		fmt.Println(i, string(letra))
	}

	//interar por um map
	usuario := map[string]string{
		"nome":      "Renata",
		"sobrenome": "Santos",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

}

/*
RANGE == VARIEDADE

Intera por cada item do slice ou array

por padrao ja espera indice, nome
Senão quiser mostrar algum parametro colocar anderline ao assinalar oos params

Não tem como interar sobre structs

*/
