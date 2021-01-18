package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	default:
		return "Invalido caiu no Default"
	}
}

func main() {
	dia := diaDaSemana(2)
	fmt.Println(dia)

}

/*

Switch == Interruptor
CONDICIONAIS LIMITADAS ::
Usabilidade : Quando vc sabe as opções que poderam ser usadas
Use: Atraves de uma opção escolhida --Devolve um resultado

------------------------------------------------
# MONTAGEM DA FUNCAO SWITCH -- INTERRUPTOR QUE FORNCERA AS OPÇÕES COM SEUS PARAMETROS

func <nomeDaFuncao>(<parametroQueTrataráOsDados> <tipoDoParam>) <tipoDoRetorno>
	switch>interruptor <parametroQueTrataráOsDados> {
	case>caso param1:
		return "Retorne esta opção 1"

	case>caso param2:
		return "Retorne esta opção 2"

	default:
		return "Nenhuma das opções --CAIU NO DEFAULT"
	}
------------------------------------------------	

------------------------------------------------
Montagem da Invocação que recuperara a funcao :::

varReceptora := <funcaoCaso(<opcaodaFuncao>)>
	mostrar(varReceptora)
-------------------------------------------------
*/
