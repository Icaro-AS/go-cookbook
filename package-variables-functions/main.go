package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello World!!!")
}

func hora() {
	fmt.Println("Olha a hora!!!")
	fmt.Println("A hora é:", time.Now())
}

func main() {
	hello()
	hora()
}
