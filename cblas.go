package vlib

import (
	"log"

	"gonum.org/v1/gonum/blas/cblas128"
)

// Dotu implements the cblas128 based operation
func Dotu(input1 VectorC, input2 VectorC) complex128 {

	v1 := cblas128.Vector{Data: input1, N: len(input1), Inc: 1}
	v2 := cblas128.Vector{Data: input2, N: len(input2), Inc: 1}

	v := cblas128.Dotu(v1, v2)

	return v
}

// MulC implements the cblas128 based operation
func MulC(input1 MatrixC, input2 MatrixC) MatrixC {
	if input1.NCols() != input2.NRows() {
		log.Panicln("Matrix Dimension Mismatch to Multiply")
	}
	result := NewMatrixC(input1.NRows(), input2.NCols())

	for r := 0; r < input1.NRows(); r++ {
		for c := 0; c < input2.NCols(); c++ {
			result[r][c] = Dotu(input1[r], input2.GetCol(c))
		}
	}
	return result
}

// Mul implements the cblas128 based operation
func Mul(input1 MatrixF, input2 MatrixF) MatrixF {
	if input1.NCols() != input2.NRows() {
		log.Panicln("Matrix Dimension Mismatch to Multiply")
	}
	result := NewMatrixF(input1.NRows(), input2.NCols())

	for r := 0; r < input1.NRows(); r++ {
		for c := 0; c < input2.NCols(); c++ {

			result[r][c] = Sum(ElemMultF(input1[r], input2.GetCol(c)))
		}
	}
	return result
}
