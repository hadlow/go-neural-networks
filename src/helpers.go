package main

import (
	"fmt"
	"math/rand"

	"gonum.org/v1/gonum/mat"
)

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
