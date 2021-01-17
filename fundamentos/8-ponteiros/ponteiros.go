package main

import "fmt"

func main() {
	var var1 int = 10
	var var2 int = var1
	fmt.Println(var1, var2)
	var1++
	fmt.Println(var1, var2)
	// Aqui var1 ta com valor : 11 e var2 ta com valor 10
	/*
		## ATRIBUINDO VALORES POR COPIA ::
			No exemplo acima :
			a var1 começa com 10
			atribuimos este valor a var2 que agora var2 vale 10 tabm
			Depois encrementamos ++1 a var1 ...
			entao var1 agora vale 11 e a var2 continua valendo 10
			a var atribuida recebeu uma copia e continua com o valor da copia de inicio ate que alguma coisa seja feita com ela propria
			o que for feito com a original nao mudara nada pra copia = append(o que for feito com a original nao mudara nada pra copia

			E daqui para baixo elas nao tem mais vinculo se mudar em uma vai mudar só nesta var mexida , apesar de uma ser atribuida a outra

	*/
	var2++
	fmt.Println(var1, var2)

	/*

		## PONTEIRO É UMA REFERENCIA DE MEMORIA
		PONTEIROS
		É uma variavel que vai guardar dentro dela um endereço de memoria

		Quando vc cria um ponteiro e passa uma var para ele vc nao ta passando o valor da variavel
		vc ta passando só o endereço de memoria
		- Ponteiro é a Porta Identificada da gaveta ou Recipiente que guarda a variavel
		--Variavel atribuida é chave e o valor

		entao Ponteiro é a gaveta ... vc diz a variavel ta nesta gaveta aqui

	*/

	var var3 int = 100
	var ponteiro *int
	fmt.Println(var3, ponteiro)
	/*
		ponteiro vc assinala no antes do tipo com *asteristico
		O valor Zero de inicialização d eum ponteiro é <nil> <nulo>

	*/

	var var4 = 100
	ponteiro = &var4
	fmt.Println(var4, ponteiro) // Resp : 0xc000016138

	// Para conseguir atribuir valor a um ponteiro coloca & no inicio da var que vai fazer esta atribuição
	// Assim o ponteiro retornará o endereço de memoria que esta contido esta var com este valor atribuido guardado
	// DESREFERENCIAÇÃO -- DESFAZENDO A REFERENCIA : Para poder ver o valor guardado neste ponteiro colocar o *asteristico na frente do ponteiro na invocação

	fmt.Println(var4, *ponteiro) // Resp: 100 100

}

/*
	Usabilidade : Um Ponteiro não muda seu endereço de memoria
	pode-se mudar o valor la dentro mas o endereço sempre sera o mesmo
	e isso sera util em aplicações que necessita de algo imutavel
	Vc precisa usar um valor que esta na sua memoria e precisa uma mesma informação para varios lugares, ir modificando ele ao longo da aplicação sem ficar passando uma copia

*/
