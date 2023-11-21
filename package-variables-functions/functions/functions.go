package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

//A função add também pode ser escritas dessa forma

func add2(x, y int) int {
	return x + y
}

//As funções também podem ter múltiplos retornos
func swap(x, y string) (string, string) {
	return y, x
}

func split(sum float32) (x, y float32) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(add2(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))
}
