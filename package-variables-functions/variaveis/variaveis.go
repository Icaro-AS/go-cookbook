package main

import "fmt"

var c, python, java bool

//As variáveis podem ter inicializadores
var i, j int = 1, 2

//declarando variáveis como uma lista de fun
func declarandoVariaveis() {
	var i int
	fmt.Println(i, c, python, java)
}

func main() {
	declarandoVariaveis()

	//No go você pode não explicitar o tipo, ele assume o tipo de acordo como valor passado
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

	/*Declaração de variáveis curtas podem ser utilizadas somente dentro de funções*/

	var i, j int = 1, 2
	k := 3
	d, julia, csharp := true, false, "no!"

	fmt.Println(i, j, k, d, julia, csharp)

}
