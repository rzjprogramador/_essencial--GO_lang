package main

import "fmt"

func main() {

	//Declaração var
	var variavel1 string = "Variavel 1"

	/*
		# variavel var ::: para funcionar fora do escopo de modo global
		--- Variável declarada em um code block é undefined em outro
		--- Para variáveis com uma abrangência maior, package level scope, utilizamos `var`
		--- Funciona em qualquer lugar
	*/

	//------------------------------------

	// # Declaração com Inferencia de tipo o GO identifica o tipo da variavel conforme seu primeiro valor atribuido
	variavel2 := "Variavel 2"

	//------------------------------------

	// # Declaração estruturada
	var (
		variavel3 string = "foo"
		variavel4 string = "bar"
	)
	//------------------------------------

	//# Declaração multipla
	variavel5, variavel6 := "Variavel 5", "Variavel 6"

	// # Statement é uma declaração que faz alguma ação..no caso fez uma soma e atribuiu a var x
	x := 10 + 10
	//------------------------------------

	fmt.Println(x)
	fmt.Printf("valor da %v \n", variavel1)
	fmt.Printf("valor da %v \n", variavel2)
	fmt.Printf("valor da %v \n", variavel3)
	fmt.Printf("valor da %v \n", variavel4)
	fmt.Printf("valor da %v \n", variavel5)
	fmt.Printf("valor da %v \n", variavel6)

}

/*
- := parece uma marmota (gopher) ou o punisher.
- Uso:
    - Tipagem automática
    - Só pode repetir se houverem variáveis novas
    - != do assignment operator (operador de atribuição)
    - Só funciona dentro de codeblocks
- Terminologia:
    - keywords (palavras-chave) são termos reservados
    - operadores, operandos
    - statement (declaração, afirmação) → uma linha de código, uma instrução que forma uma ação, formada de expressões
    - expressão → qualquer coisa que "produz um resultado"
    - scope (abrangência)
        - package-level scope
- Lição principal:
    - := utilizado pra criar novas variáveis, dentro de code blocks
    - = para atribuir valores a variáveis já existentes
*/
