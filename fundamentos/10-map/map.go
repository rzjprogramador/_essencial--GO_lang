package main

import "fmt"

func main() {

	// Map de strings

	usuario := map[string]string{
		"nome":      "Reinaldo",
		"sobrenome": "Zacharias",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	//-----------------------
	// Map de int e string
	usuario2 := map[int]string{
		1: "identificador",
		2: "qualquerCoisa",
	}

	fmt.Println(usuario2)
	fmt.Println(usuario2[1])

	/-----------------------
	// Map aninhado
	// è um := map que a chave é uma string > e o valor é um map onde a chave é uma string e o valor pe uma string  
	usuarioAninhado := map[string]map[string]string



}

/*
# FUNÇÃO MAP
<nomedomap> := map[<tipoDaChave>]<tipoDoValor> {
	"chave": "valor"
}
// RECUPERAR
fmt.Println(<nomeMap>)
fmt.Println(<>nomeMap["<chave para acessa o valor>"])

//--------------------------------
Map é para estrutura de dados semelhante ao struct
A diferença é que a chave e o valor tem que ser tudo do mesmo tipo



*/
