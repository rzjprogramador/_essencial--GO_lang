/*
# TIPOS ::
NUMERO     == int (int8, int16, int32, int64) de 8 em 8 bytes
DECIMAL    == float (float8, float16, float32, float64) de 8 em 8 bytes
TEXTO      == string
CARACTERE  == string (Obs; Não tem char caractere unico no GO)

==============================

# FUNÇÕES DE COLEÇÕES

> reflect.TypeOf(<varSlice>) ::: mostra o tipo de uma variavel
> append(<varSlice>)  ::: acrescentar items ao slice

# FUNÇÕES DE TEMPO
.......................................
TIME
time.Sleep(time.Second) // Faz dar uma pausa de segundos tipo Durmir 1 segundo
.......................................


	====================================
# FUNÇÕES CRIADORAS
função make -- aloca um espaço na memoria para criar algo
DECLARAÇÃO COMO É CONSTRUIDO O SLICE :	nomeSlice := make([])<<tipo>>, <quantidade-inicial>, <capacidade>

==============================

// %v  para saber valor , SEPARAR ASPAS DA VAR COM VIRGULA e coloca a variavel depois das aspas
// %T para mostrar o tipo e a var que quer saber o tipo depois das aspas tbm
// iMPORTANTE ::  SEPARAR ASPAS DA MENSAGEM COM AS VARIAVEIS UTILIZADAS COM VIRGULA
// \n -- para pular linha.

==============================


*/