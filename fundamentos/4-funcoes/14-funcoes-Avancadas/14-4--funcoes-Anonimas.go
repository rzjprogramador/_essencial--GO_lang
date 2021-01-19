package main

import "fmt"

func main() {

	// funcao anonima com () no final dela, para assim que declarar ja executar
	func() {
		fmt.Println("Ola mundo")
	}()

	// Passando parametros na funcao anonima
	func(texto string) {
		fmt.Println(texto)
	}("Passando Parametro na funcao Anonima")
	
	
	// Colocando retorno na Funcao Anonima
	// guarda numa variavel a função -- especifica o tipo de retorno da funcao depois dos parametros --- recupera a variavel que guarda a função

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido --> %s", texto)
	}("Passando Parametro")
	fmt.Println(retorno)
}
