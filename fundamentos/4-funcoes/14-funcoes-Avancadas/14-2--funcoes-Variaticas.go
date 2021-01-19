package main

import "fmt"

func soma(numeros ...int) int {
	fmt.Println(numeros)
	//retorna um slice - se é um slice posso interar sobre ele

	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func main() {
	totalDaSoma := soma(1, 3, 2, 4, 9, 8, 10)
	fmt.Println(totalDaSoma)
}

/*

FUNCÇÕES VARIATICAS
Podemos definir a quantidade de parametros que ela vai receber
... 3 pontinhos depois da variavel no parametro quer dizer N-numeros sem limite de parametros, pode passar nada tbm 
Para operações Os dados passados tem que ser do messmo tipo
para passar tipos diferentes veremos na proxima aula.

IMPORTANTE : Sempre que assinalar tipo de retorno no Parametro da função - Atribuir a uma nova variavel o resultado na invoção

Exemplo : SOMANDO VARIOS NUMEROS DE UMA COLEÇÃO


*/
