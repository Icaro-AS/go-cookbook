package main

/*Os tipos basicos do Go são

bool
string
int
int8
int16
int32
int64
uint
uint8
uint16
uint32
uint64
uintptr
byte // pseudônimo para uint8
rune // pseudônimo para int32 // representa um ponto de código Unicode
float32
float64
complex64
complex128
*/

import (
	"fmt"
	"math"
	"math/cmplx"
)

const Pi = 3.14

var (
	ToBe   bool       = false
	MaxInt uint       = 1<<64 - 1
	Z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {

	//Descobre tipo e valor das variáveis
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", Z, Z)

	//Valores zero
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	//Convertendo tipos
	var x, y = 3, 4
	var numberf float64 = math.Sqrt(float64(x*x + y*y))
	fmt.Printf("Valor = %T\n", numberf)
	var z uint = uint(numberf)
	fmt.Println(numberf, x, y, z)

	//constantes
	const World = "Meu mundo"
	fmt.Println("Hello", World)
	fmt.Println(Pi)

	const Truth = true

	fmt.Println("Go Rules?", Truth)

}
