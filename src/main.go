package main

import "fmt"
import "math/rand"
import "gonum.org/v1/gonum/mat"

func fmtMatrix(name string, matrix *mat.Dense) {
	fMatrix := mat.Formatted(matrix, mat.Prefix("    "), mat.Squeeze())

	fmt.Printf(name + " = %v", fMatrix)
}

func randMatrix(x int, y int) *mat.Dense {
	data := make([]float64, x * y)

	for i := range data {
		data[i] = rand.NormFloat64()
	}

	return mat.NewDense(x, y, data)
}

// Hyperparameters
var n_inputs int = 5
var n_hidden int = 3
var n_output int = 2

func main() {
	b := mat.NewDense(3, 5, nil)
	w := randMatrix(6, 4)

	fmt.Println(w)
	fmtMatrix("b", b)
	//fmtMatrix("w", w)
}
