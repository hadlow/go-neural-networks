package main

import "fmt"
import "math"
import "math/rand"
import "gonum.org/v1/gonum/mat"

func fmtMatrix(name string, matrix *mat.Dense) {
	fMatrix := mat.Formatted(matrix, mat.Prefix("    "), mat.Squeeze())

	fmt.Printf(name + " = %v\n", fMatrix)
}

func randMatrix(x int, y int) *mat.Dense {
	data := make([]float64, x * y)

	for i := range data {
		data[i] = rand.NormFloat64()
	}

	return mat.NewDense(x, y, data)
}

func log(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

func relu(matrix *mat.Dense) *mat.Dense {
	return matrix
}

func cz(x *mat.Dense, w *mat.Dense, b *mat.Dense) *mat.Dense {
	wx := mat.NewDense()
}

// Hyperparameters
var n_inputs int = 5
var n_hidden int = 3
var n_output int = 2
var n_samples int = 3

var data []float64 = []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
var x *mat.Dense = mat.NewDense(n_inputs, n_samples, data)

func main() {
	// First layer
	b := mat.NewDense(n_hidden, 1, nil)
	w := randMatrix(n_inputs, n_hidden)

	fmtMatrix("b", b)
	fmtMatrix("w", w)

	// z = (w * x) + b
	z := cz(x, w, b)

	fmtMatrix("z", z)
	fmtMatrix("w", relu(w))
}
