package main

import "fmt"

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	escrever("Ola mundo", 10, 20, 30, 100)
}

/*

FUNCÇÕES VARIATICAS
PASSANDO COMO APRAM STRING E ...N-NUMEROS
ex: ESCREVENDO UMA FRASE VARIAS VEZES CONCATENADA COM NUMEROS DIFERENTES


*/
