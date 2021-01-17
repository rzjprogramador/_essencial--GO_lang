/*
FUNÇÕES DE COLEÇÕES
======================================

reflect.TypeOf
função que mostra o tipo de uma variavel
Usando ::
// Analisando qual o tipo deste slice
	fmt.Println(reflect.TypeOf(slice))
	//resp: []int  --é um slice de int

	====================================
append  == acrescentar
acrescentar items ao slice

Usando :
slice = append(slice, 123)
	fmt.Println(slice)

	====================================

// %v  para saber valor , SEPARAR ASPAS DA VAR COM VIRGULA e coloca a variavel depois das aspas
// %T para mostrar o tipo e a var que quer saber o tipo depois das aspas tbm
// iMPORTANTE ::  SEPARAR ASPAS DA MENSAGEM COM AS VARIAVEIS UTILIZADAS COM VIRGULA
// \n -- para pular linha.

==============================
TIPOS ::
NUMERO     == int (int8, int16, int32, int64) de 8 em 8 bytes
DECIMAL    == float (float8, float16, float32, float64) de 8 em 8 bytes
TEXTO      == string
CARACTERE  == string (Obs; Não tem char caractere unico no GO)




*/