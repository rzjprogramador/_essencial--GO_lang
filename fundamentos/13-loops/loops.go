package main

import (
	"fmt"
	"time"
)

func main() {

	for j := 0; j < 6; j += 2 {
		fmt.Println("Incrementando o J que tem o valor", j)
		time.Sleep(time.Second)
	}

	/*
		Declarando o valor da var interadora no escopo do for
		Obs essa variavel declarada no for só fica visivel ai neste escopo
		não da para acessa-la fora do escopo do for

		// incrementando de 1 em 1 usar no encrementador : ++
		// incrementando de 2 em 2 usar no encrementador : += 2



	*/

	/*
		// Com o valor interador declarado antes
		i := 0

		for i < 5 {
			i++
			fmt.Println("Incrementando o i")
			time.Sleep(time.Second)

		}
		fmt.Println("O valor de i ... é : ", i)
		/*
			Igual o while :
			Enquanto uma condicao for true --repete um bloco de codigo



	*/

}
