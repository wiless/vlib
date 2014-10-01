package vlib

import (
	"fmt"
	"log"
	"math"
	"math/cmplx"
	// "os"
	// "regexp"
	// "strconv"
	// "strings"
)

func init() {
	fmt.Printf("\n%% === Vlib Initialized : - github.com/wiless ===\n") /// matlab or octave compatible dump
}

//// Functions for Complex Vectors

func Conj(in1 VectorC) VectorC {
	result := NewVectorC(in1.Size())
	for i := 0; i < in1.Size(); i++ {
		result[i] = cmplx.Conj(in1[i])
	}
	return result
}

//////// Methods over the Complex Vector
func ElemMultC(in1, in2 VectorC) VectorC {
	size := len(in1)
	result := NewVectorC(size)

	for i := 0; i < size; i++ {
		result[i] = in1[i] * in2[i]
	}

	return result
}

func (v *VectorC) SetVectorF(input VectorF) {
	v.Resize(input.Size())
	for i := 0; i < v.Size(); i++ {
		(*v)[i] = complex(input[i], 0)
	}
}

func NewVectorC(size int) VectorC {
	return VectorC(make([]complex128, size))
}

func (v *VectorC) Resize(size int) {
	// Only append etc length
	length := len(*v)
	extra := (size - length)
	if extra > 0 {
		tailvec := NewVectorC(extra)
		*v = append(*v, tailvec...)
	}

	///copy(*v, Vector(make([]int, size)))

}

func NewOnesC(size int) (v VectorC) {
	result := NewVectorC(size)
	for i := 0; i < size; i++ {
		result[i] = 1 + 0i
	}
	return result
}

func (v *VectorC) OnesF() {
	for i := 0; i < len(*v); i++ {
		(*v)[i] = 1
	}
}

func (v VectorC) ScaleC(factor complex128) VectorC {

	result := NewVectorC(len(v))
	for indx, val := range v {
		result[indx] = val * factor
	}
	return result
}
func (v VectorC) Scale(factor float64) VectorC {

	return v.ScaleC(complex(factor, 0))
}

func (v VectorC) Size() int {
	return len(v)
}

func (v *VectorC) PlusEqual(input VectorC) {
	if len(*v) != len(input) {
		log.Panicf("\n PlusEqual %d : Length Mismatch %d", v.Size(), input.Size())

	}
	cnt := v.Size()
	for k := 0; k < cnt; k++ {

		(*v)[k] = (*v)[k] + input[k]
	}

}

func (v *VectorC) Shift(input VectorC) {
	if len(*v) != len(input) {
		log.Panicf("\n PlusEqual %d : Length Mismatch %d", v.Size(), input.Size())

	}
	cnt := v.Size()
	for k := 0; k < cnt; k++ {

		(*v)[k] = (*v)[k] + input[k]
	}

}

func (v VectorC) Insert(pos int, val complex128) VectorC {
	result := NewVectorC(v.Size() + 1)
	copy(result[0:pos], v[0:pos])
	result[pos] = val
	copy(result[pos+1:], v[pos:])
	return result
}

func (v VectorC) Delete(pos int) VectorC {

	result := NewVectorC(v.Size())
	copy(result, v)
	copy(result[pos:], result[pos+1:])

	// result[v.Size()-1] = nil // or the zero value of T

	return result[:v.Size()-1]

}

func (v VectorC) AddVector(input VectorC) VectorC {
	result := NewVectorC(len(v))
	if v.Size() != input.Size() {
		log.Panicf("\nAddVector %d : Length Mismatch %d", v.Size(), input.Size())
	}
	for k := range v {
		result[k] = v[k] + input[k]
	}
	return result
}

func GoDotC(input1 VectorC, input2 VectorC, splitN int) complex128 {

	if input1.Size() != input2.Size() {
		log.Panicf("Dot: LHS (%d) RHS (%d) size mismatch", input1.Size(), input2.Size())
	}
	sublen := input1.Size() / splitN
	outCH := make(chan complex128, splitN)

	for i := 0; i < splitN; i++ {
		in1 := input1[i*sublen : sublen*(i+1)]
		in2 := input2[i*sublen : sublen*(i+1)]
		// log.Printf("\n Start %d Splitting into %d of each length %d", i*sublen, splitN, sublen)
		go func(in1, in2 VectorC, outch chan complex128) {
			temp := ElemMultC(in1, in2)
			var result complex128 = 0.0
			for _, val := range temp {
				result += val
			}
			outCH <- result
		}(in1, in2, outCH)
	}
	var sum complex128 = 0
	for i := 0; i < splitN; i++ {
		sum += <-outCH
	}

	return sum
}

func DotC(input1 VectorC, input2 VectorC) complex128 {
	if input1.Size() != input2.Size() {
		log.Panicf("Dot: LHS (%d) RHS (%d) size mismatch", input1.Size(), input2.Size())
	}
	temp := ElemMultC(input1, input2)
	var result complex128 = 0.0
	for _, val := range temp {
		result += val
	}
	return result
}

func ElemAddCmplx(in1, in2 []complex128) []complex128 {
	size := len(in1)
	result := make([]complex128, size)

	for i := 0; i < size; i++ {
		// bool(in1[i])
		result[i] = in1[i] + in2[i]
	}

	return result
}

func SumC(v VectorC) complex128 {
	var result complex128
	for _, val := range v {
		result += val
	}
	return result

}

func (v VectorC) ShiftAndScale(shift, scale complex128) VectorC {

	// v = v.Add(shift)
	// result := v.Scale(factor)

	result := NewVectorC(v.Size())
	for i := 0; i < v.Size(); i++ {
		result[i] = (v[i] + shift) * scale
	}
	return result
}
func (v VectorC) ScaleAndShift(shift, scale complex128) VectorC {

	// v = v.Add(shift)
	// result := v.Scale(factor)

	result := NewVectorC(v.Size())
	for i := 0; i < v.Size(); i++ {
		result[i] = v[i]*scale + shift
	}
	return result
}

func (v *VectorC) Zeros() {
	v.Fill(0)
}

func (v *VectorC) Ones() {
	v.Fill(1)
}

func (v *VectorC) Fill(val complex128) {
	for i := 0; i < v.Size(); i++ {
		(*v)[i] = val
	}
}

func MeanAndVarianceC(v VectorC) (mean complex128, variance float64) {

	mean = SumC(v) / complex(float64(v.Size()), 0)
	variance = 0
	for i := 0; i < v.Size(); i++ {
		variance += math.Pow(cmplx.Abs(v[i]), 2.0)
	}
	variance = variance / float64(v.Size()-1)

	return mean, variance
}

func MeanC(v VectorC) complex128 {

	return SumC(v) / complex(float64(v.Size()), 0)

}

/// returns Euclidean Norm of the vector
func VarianceC(v VectorC) float64 {
	var result float64 = 0
	mean := MeanC(v)
	for i := 0; i < v.Size(); i++ {
		result += math.Pow(cmplx.Abs(v[i]-mean), 2.0)
	}

	return result / float64(v.Size()-1)

}

/// returns the sum of square of the elements in the vector
func EnergyC(v VectorC) float64 {
	var result float64 = 0
	for i := 0; i < v.Size(); i++ {
		result += math.Pow(cmplx.Abs(v[i]), 2.0)
	}

	return result
}

/// returns 2nd Norm of the vector (\sum(x[i]))^(1/2)
func Norm2C(v VectorC) float64 {
	var result float64 = 0
	for i := 0; i < v.Size(); i++ {
		result += math.Pow(cmplx.Abs(v[i]), 2.0)
	}

	return math.Sqrt(result)

}

/// returns Euclidean Norm of the vector
func NormC(v VectorF) float64 {
	var result float64 = 0
	for i := 0; i < v.Size(); i++ {
		result += math.Pow(v[i], 2.0)
	}

	return math.Sqrt(result / float64(v.Size()))

}

func AddC(A, B VectorC) VectorC {
	result := NewVectorC(A.Size())
	for i := 0; i < A.Size(); i++ {
		result[i] = A[i] + B[i]
	}
	return result
}

func SubC(A, B VectorC) VectorC {
	if A.Size() != B.Size() {
		log.Panicf("Sub: LHS (%d) and RHS (%d) size mismatch", A.Size(), B.Size())
	}
	result := NewVectorC(A.Size())
	for i := 0; i < A.Size(); i++ {
		result[i] = A[i] - B[i]
	}
	return result
}

func (v VectorC) ToUnitEnergy() (result VectorC, factor float64) {

	temp := ElemMultC(v, v)
	factor = 1.0 / math.Sqrt(cmplx.Abs(SumC(temp)))
	result = v.Scale(factor)

	return result, factor
}

/// Normalizes with 0 mean, and unit variance
func (v VectorC) Normalize() (result VectorC, mean complex128, factor float64) {

	mean, variance := MeanAndVarianceC(v)
	factor = 1.0 / math.Sqrt(variance)
	result = v.ShiftAndScale(-mean, complex(factor, 0))

	// v = v.Sub(mean)
	// result = v.Scale(factor)

	return result, mean, factor
}

/// Input element is pushed to end of the vector and first element is removed
func (v VectorC) ShiftLeft(val complex128) VectorC {
	N := v.Size()
	result := v.Insert(0, val)
	return result.Delete(N)

}
