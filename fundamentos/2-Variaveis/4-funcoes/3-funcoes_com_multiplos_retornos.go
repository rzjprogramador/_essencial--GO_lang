package main

import "fmt"

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	resultSoma, resultSub := calculosMatematicos(10, 15)
	fmt.Println(resultSoma, resultSub)

	// IGNORANDO ou ESCOLHENDO PARAMETROS A SEREM USADOS >>
	//  _, resultSub := calculosMatematicos(10, 15)
	// fmt.Println(resultSub)

}

/*
SOMAR E SUBTRAIR 2 NUMEROS AO MESMO TEMPO
FAZER 2 OPERAÇÕES NOS MESMOS PARAMETROS

criada funcao calculos (passados 2 parametros com seus tipos) passados a quantidade em tipos do retorno que se espera da funcao
instruindo o que se espera da funcao a soma := n1 + n2 ....... e  a subtracao := n1 - n2
definindo o retorno das variaveis soma e subtracao

Rodando na funcao main esta funcao para ter os resultados :::
criando 2 variaveis para receber por atribuição curta a funcao calculos que traz nos aprametros os numeros que serao operados
printar na tela as variaveis que receberam essas atribuições da funcao

- SE NAO QUISER USAR UM PARAMETRO Q VEM NA FUNÇÃO COLOCAR _ Handerline para ignorar um parametro
-Obs se for ignorar cuidado com a ordem dos fatores a Ordem de cada parametro é alinhada com o tipo que ele vai retornar
- Ignore com _ exatamente o parametro certo que nao quer usar no retorno
*/
