package mathbeta

import (
	"fmt"
	"log"
	"testing"
)

func TestMatrix(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Runtime error caught: %v", r)
		}
	}()
	m := NewMatrix(3, 3, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	//	m := RandMatrix(3, 3)
	fmt.Println("matrix:")
	m.Print()
	d, err := m.Determinant()
	if err == nil {
		fmt.Println("matrix determinant:", d)
	} else {
		t.Error("matrix should has a determinant")
		fmt.Println(err)
	}

	inverse, err := m.Inverse()
	if err == nil {
		t.Log("matrix should has no inversion")

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
