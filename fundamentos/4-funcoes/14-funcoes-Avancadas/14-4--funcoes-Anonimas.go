package main

import "fmt"

func main() {

	// funcao anonima com () no final dela, para assim qeu declarar ja executar
	func() {
		fmt.Println("Ola mundo")
	}()

	// Passando parametros na funcao anonima
	func(texto string) {
		fmt.Println(texto)
	}("Passando Parametro na funcao Anonima")
	
	// Colocando retorno na Funcao Anonima
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido --> %s", texto)
	}("Passando Parametro")
	fmt.Println(retorno)
}
