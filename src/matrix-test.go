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
	d, err := m.Determinant()
	if err == nil {
		fmt.Println("matrix determinant:", d)
	} else {
		fmt.Println(err)
	}

	inverse, err := m.Inverse()
	if err == nil {
		inverse.Print()
		multiplication, err := m.Multiply(inverse)
		if err == nil {
			multiplication.Print()
		}

		multiplication, err = inverse.Multiply(m)
		if err == nil {
			multiplication.Print()
		}
	} else {
		fmt.Println(err)
	}

	//	mathbeta.Ones(3, 2).Print()
	//	m.Transpose().Print()

	//	m = mathbeta.RandMatrix(5, 3)
	//	fmt.Println("rand matrix")
	//	m.Print()
}
