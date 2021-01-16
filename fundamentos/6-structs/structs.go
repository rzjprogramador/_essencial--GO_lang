package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func main() {
	var u usuario
	fmt.Println(u)

	// forma 1 de Popular uma instancia do struct usuario
	u.nome = "Reinaldo"
	u.idade = 43
	fmt.Println(u)

	// forma 2 de popular Instancia + recomendada

	usuario2 := usuario{"Gustavo", 10}
	fmt.Println(usuario2)

	// Recuperar somente os campos que desejar :: coloca-se só a chave e o valor dos campos que desejar
	usuario3 := usuario{nome: "Leonardo"}
	fmt.Println(usuario3)

}

/*
STRUCTS == ESTRUTURAS

Como se fosse uma classe
è uma coleção de campos
cada campo tem um nome e um tipo

*/
