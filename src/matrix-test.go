package main

import (
	"fmt"
	"log"
	"mathbeta"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Runtime error caught: %v", r)
		}
	}()
	m := mathbeta.NewMatrix(3, 3, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	//	m := mathbeta.RandMatrix(3, 3)
	fmt.Println("matrix:")
	m.Print()
	fmt.Println("matrix determinant:", m.Determinant())
	inverse := m.Inverse()
	inverse.Print()

	m.Multiply(inverse).Print()
	inverse.Multiply(m).Print()

	//	mathbeta.Ones(3, 2).Print()
	//	m.Transpose().Print()

	//	m = mathbeta.RandMatrix(5, 3)
	//	fmt.Println("rand matrix")
	//	m.Print()
}
