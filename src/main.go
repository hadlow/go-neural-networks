package main

import "fmt"
import "gonum.org/v1/gonum/mat"

func pMatrix(name string, matrix *mat.Dense) {
	fMatrix := mat.Formatted(matrix, mat.Prefix("    "), mat.Squeeze())

	fmt.Printf(name + " = %v", fMatrix)
}

func main() {
	zero := mat.NewDense(3, 5, nil)

	pMatrix("b", zero)
}
