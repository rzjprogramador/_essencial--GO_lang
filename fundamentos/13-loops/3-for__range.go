package main

import (
	"fmt"
	"time"
)

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

	// Loop Infinito se não colocar nada ou colocar true vai rodar infinito
	for false {
		fmt.Println("Executando Infinitamente")
		time.Sleep(time.Second)
	}

}

/*
RANGE == VARIEDADE

Intera por cada item do slice ou array

por padrao ja espera indice, valor
Senão quiser mostrar algum parametro colocar anderline ao assinalar oos params

Não tem como interar sobre structs

*/
