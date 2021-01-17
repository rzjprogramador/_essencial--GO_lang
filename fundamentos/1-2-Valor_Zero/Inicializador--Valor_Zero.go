package main

import (
	"fmt"
)

// DECLARAÇÃO DE VALOR ZERO -- dos TIPOS PRIMITIVOS
var a int
var b float64
var c string
var d bool

func main() {

	fmt.Printf("%T inicializa com valor : %v \n", a, a)
	fmt.Printf("%T inicializa com valor : %v \n", b, b)
	fmt.Printf("%T inicializa com valor : vazio %v \n", c, c)
	fmt.Printf("%T inicializa com valor : %v \n", d, d)
}

/*

# VALOR ZERO INICIALIZAÇÃO DE VARIAVEIS
TIPOS PRIMITIVOS

- Declaração vs. inicialização vs. atribuição de valor. Variáveis: caixas postais.
- O que é valor zero?
- Os zeros:
    - ints: 0
    - floats: 0.0
    - booleans: false
    - strings: ""
    - pointers, functions, interfaces, slices, channels, maps: nil
- Use := sempre que possível.
- Use var para package-level scope.

[https://www.youtube.com/watch?v=Ruwp2xLD_AI&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=12]


*/
