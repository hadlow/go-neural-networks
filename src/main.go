package main

import "gonum.org/v1/gonum/mat"

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
	w1 := randMatrix(n_inputs, n_hidden)
	b1 := mat.NewDense(1, n_hidden, nil)

	// Second layer
	w2 := randMatrix(n_hidden, n_output)
	b2 := mat.NewDense(1, n_output, nil)

	fmtMatrix("b", b)
	fmtMatrix("w", w)

	// z = (w * x) + b
	z := cz(x, w, b)

	fmtMatrix("z", z)
	fmtMatrix("w", relu(w))
}
