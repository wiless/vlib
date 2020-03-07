// Package vlib gonum.go friendly C and R vectors
// Based on gonum/examples - This example shows how simple user types can be constructed to
// implement basic vector functionality within the mat package.
package vlib

import (
	"fmt"

	"gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/gonum/mat"
)

// ExampleUserVectors
func ExampleUserVectors() {
	// Perform the cross product of [1 2 3 4] and [1 2 3].
	r := R{1, 2, 3, 4}
	c := C{1, 2, 3}

	var m mat.Dense
	m.Mul(c, r)

	fmt.Println(mat.Formatted(&m))

	// Output:
	//
	// ⎡ 1   2   3   4⎤
	// ⎢ 2   4   6   8⎥
	// ⎣ 3   6   9  12⎦
}

// R is a user-defined R vector.
type R []float64

// Dims, At and T minimally satisfy the mat.Matrix interface.
func (v R) Dims() (r, c int)    { return 1, len(v) }
func (v R) At(_, j int) float64 { return v[j] }
func (v R) T() mat.Matrix       { return C(v) }

// RawVector allows fast path computation with the vector.
func (v R) RawVector() blas64.Vector {
	return blas64.Vector{N: len(v), Data: v, Inc: 1}
}

// C is a user-defined C vector.
type C []float64

// Dims, At and T minimally satisfy the mat.Matrix interface.
func (v C) Dims() (r, c int)    { return len(v), 1 }
func (v C) At(i, _ int) float64 { return v[i] }
func (v C) T() mat.Matrix       { return R(v) }

// RawVector allows fast path computation with the vector.
func (v C) RawVector() blas64.Vector {
	return blas64.Vector{N: len(v), Data: v, Inc: 1}
}
