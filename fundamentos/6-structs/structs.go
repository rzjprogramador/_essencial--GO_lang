package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	longradouro    string
	cep            string
	numeroEndereco uint8
}

func main() {
	var u usuario
	fmt.Println(u)

	// Instanciando struct usuario e recuperando poucos campos
	u.nome = "Reinaldo"
	u.idade = 43
	u.endereco.longradouro = "Rua do Reinaldo"
	fmt.Println(u)

	// forma 2 de popular Instancia -- requer recuperar todos campos
	// Criado o struct endereco e assinalado no struct usuario
	// para recupera-lo criar um variavel para receber os desse struct adicionado e preencher seus valores
	enderecoExemplo := endereco{"Rua do Programador", "08474000", 72}

	u2 := usuario{"Gustavo", 10, enderecoExemplo}
	fmt.Println(u2)

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
