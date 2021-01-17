package main

import "fmt"

func main() {
	fmt.Println("======== ESTRUTURAS DE CONTROLE ========")

	// CONDICIONAL SIMPLES
	numero := 6

	if numero >= 15 {
		fmt.Println("Maior que 15")
	} else if numero < 5 {
		fmt.Println("Numero é menor que 5")
	} else {
		fmt.Println("ENTRE 0 e 5 ...O Digitado foi : ", numero)
	}

	//IF INIT
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Numero é maior que Zero , Porque aqui numero é : ", numero)
	}

	/*
		PROTOTIPO IF INIT -- Inicializar var dentro do if
		declaradaNovaVar := atribuindoValor; declaradaNovaVar > 0 {...faça}

		ob: a variavel inicializada no if só serve apra o escopo do if é criada e usada somente la dentro

	*/
}
