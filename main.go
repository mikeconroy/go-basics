package main

import (
	"fmt"

	"github.com/GoesToEleven/puppy"
)

const pi = 3.14

var a int = 123

func main() {
	x := "Hello"
	fmt.Println("Hello World")
	fmt.Println("PI:", pi)
	fmt.Println("a:", a)
	fmt.Println("x:", x)
	fmt.Println(puppy.Bark())

}
