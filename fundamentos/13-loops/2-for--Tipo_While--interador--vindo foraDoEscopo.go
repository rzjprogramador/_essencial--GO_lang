package main

import (
	"fmt"
	"time"
)

func main() {

	/* 

	// for(para) Com o valor interador declarado fora do escopo vira tipo WHILE (Enquanto)
	// Enquanto uma condicao for true --repete um bloco de codigo
	// Posso continuar usando a var do interador fora do escopo do for

	*/

	i := 0

	for i < 5 {
		i++
		fmt.Println("Incrementando o i")
		time.Sleep(time.Second)

	}
	fmt.Println("O valor de i ... é : ", i)

}
