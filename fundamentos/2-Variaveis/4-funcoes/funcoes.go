package main

import "fmt"

// FUNCAO PARA SOMAR 2 NUMEROS
func soma(num1 int8, num2 int8) int8 {
	return num1 + num2
}

// RECUOPERANDO NA FUNCAO MAIN APRA ORDAR O PROGRAMA
func main() {
	somar := soma(10, 20)
	fmt.Println(somar)
}

/*
Funcao é como uam receita de bolo recebem instrucoes para executar algo
e podem ter como ingredientes parametros que se passados vc recebe e na invocação tem que devolver

- Funções são tipos
- Podem ser guardadas em um variavel
- Parametros podem ser funcoes tambem
- O retorno de uma função pode ser uma funcao tambem


1º - #ENTRADA -- ASSINALA
func nomeFuncao(parm tipo)tipoRetornoFuncao {
	return instrucao --com ou sem param
}
2º #PROCESSA --
DENTRO DA func main para executar o Programa
varResultadodaFuncao := infere o tipo - funcaoCriada(devolveParam,)

3º SAIDA --
--Mostra o resultada printando : varResultado

//=====================
Narrando funcao ACIMA :
Funcao somar que vai receber num1 do tipo int8 e num2 do tipo int8 porque vao ser numeros de pequenos bytes
o retorno da funcao sera do tipo int8 - entao {
	vai retornar num1 + num2
}
Criar a funcao main para ordar o programa e dentro dela instruir :
criar a varResultado chamada soma para receber com inferencia de tipos com o atribuidor curto := guardar a funcao soma(10 ,20)
somando 10 com 20
e mostrando na tela com o Print



*/
