package main

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

func log(x float64) float64 {
	return 1. / (1. + math.Exp(-x))
}

func logPrime(x float64) float64 {
	return x * (1. - x)
}

func sigmoid(matrix *mat.Dense) *mat.Dense {
	return matrix
}

func sigmoidPrime(matrix *mat.Dense) *mat.Dense {
	return matrix
}

func relu(matrix *mat.Dense) *mat.Dense {
	return matrix
}
