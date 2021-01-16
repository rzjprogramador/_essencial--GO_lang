package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	// Recuperando pessoa
	p1 := pessoa{"Reinaldo", "Zacharias", 43, 180}
	fmt.Println(p1)

	// Recuperando estudante no preenchimento dos campos colocar tbm a variavel que traz os dados da Instancia da struct herdada
	e1 := estudante{p1, "Programação", "Faculdade da Vida"}
	fmt.Println(e1)

	//acessar de forma fracionada
	fmt.Println(e1.nome)
	fmt.Println(p1.idade)

}

/*
Conceito de Herança "É UM"
ex: Estudante "É UMA" pessoa

Para fazer a simulada herança ::  cria os structs e faz a ligacao somente com o nome do struct que vai herdar Sem passar o tipo do struct herdado
Struct não pe obrigaorio usar invocar no console a variavel ...

É UMA OPÇÃO SE VC TIVER STRUCTS MUITO GRANDES QUE PRECISE PASSAR MUITOS CAMPOS É TIPO UM MIXIN
*/
