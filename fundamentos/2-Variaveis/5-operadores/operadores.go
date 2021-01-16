package main

import "fmt"

func main() {

	// ARITMETICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	// Só pode fazer operacoes com tipos iguais
	var numero1 int16 = 10
	var numero2 int16 = 25
	var resultado = numero1 + numero2
	fmt.Println(resultado)

	// ATRIBUIÇÃO
	var variavel1 string = "String1"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	// RELACIONAIS
	fmt.Println(1 == 2) // exatamente igual
	fmt.Println(1 > 2)  // maior
	fmt.Println(1 >= 2) // maior ou igual
	fmt.Println(1 < 2)  // menor
	fmt.Println(1 <= 2) // menor ou igual
	fmt.Println(1 != 2) // diferente

	// LOGICOS
	verdadeiro, falso := true, false

	// -- && é o "E" então uma E a outra condição tem que ser true para dar true
	fmt.Println(verdadeiro && verdadeiro)
	fmt.Println(verdadeiro && falso)

	// -- || é o OU ou OR --quer dizer que Basta uma condição ou outra sendo true ja retorna true
	fmt.Println(verdadeiro || falso)
	fmt.Println(falso || falso)

	// -- ! é o NEGAÇÃO ele inverte o valor booleano da condição , se era false vira true
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	// UNARIOS ++ --
	// Só agem em uma variavel por vez , quando vc quer encrementar ou decrementar o valor de uma var vc usa ele
	
	// encrementando em 1
	idade := 10
	idade++
	fmt.Println(idade)

	// encrementando em 15
	numero := 10
	numero++
	numero += 15 // numero = numero + 15
	fmt.Println(numero)
	// resposta: 26 ..Porque era 10 + 1 + 15 resultado 26

	// Decrementar diminuir
	numero--
	fmt.Println(numero)

	// Encrementando com multiplicacao
	numero *= 3 // numero = numero * 3
	fmt.Println(numero)

	// Unario com Divisao
	numero /= 3
	fmt.Println(numero)

	// Unario com Resto da Divisao
	numero %= 3
	fmt.Println(numero)

}

	/*
	
	 NÃO É POSSIVEL EM GO ::
	 > --numero ou ++numero ===== Pré-fixado decrementar antes 
	 > Operador ternario tbm nao funciona ainda estamos em 2021 
	
	
	 */
