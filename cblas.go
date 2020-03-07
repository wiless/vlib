package vlib

import "gonum.org/v1/gonum/blas/cblas128"

// Dotu implements the cblas128 based operation
func Dotu(input1 VectorC, input2 VectorC) complex128 {
	
	v1:=cblas128.Vector{Data:input1,N:len(input1),Inc: 1}
	v2:=cblas128.Vector{Data:input2,N:len(input2),Inc: 1}
	
	v := cblas128.Dotu(v1,v2)

	return v
}