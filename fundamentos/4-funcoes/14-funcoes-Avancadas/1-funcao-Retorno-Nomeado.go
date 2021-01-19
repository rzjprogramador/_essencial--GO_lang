package main

import "fmt"

func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	soma, subtracao := calculosMatematicos(10, 20)
	fmt.Println(soma, subtracao)
}

/*
RETORNO NOMEADO
Ao assinalar o retorno da funcao em variaveis que serao usadas na msm com seus tipos que vaoretornar
Usabilidade :: Fazer mais que uma operação na mesma função ..neste caso fez soma e subtracao nos mesmos valores

# ::: Montagem :::

func <nomedafuncao> (<dadoParam,>) [[O QUE VAI RETORNAR]](<varDadoParam <tipo>>) {
	varDadoParam1 = dadoParam + dadoParam2
	varDadoParam2 = dadoParam + dadoParam2
}
[[RECUPERAR NA FUNCAO MAIN]]
varDadoParam1, varDadoParam2 := nomedafuncao(valor, valor2)
mostrar(varDadoParam1, varDadoParam2)
*/
