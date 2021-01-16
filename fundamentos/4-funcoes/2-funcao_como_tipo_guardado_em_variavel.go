package main

import "fmt"

func main() {

	// Assinalando variavel como do tipo função
	// Passando parametro

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}
	resultado := f("Sou o texto da Funcao F com Retono")
	fmt.Println(resultado)

}
