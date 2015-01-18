package vlib

import (
	"math"
	"math/rand"
)

func Randsrc(size int, maxvalue int) []int {
	var result = make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = rand.Intn(maxvalue)
	}
	return result
}

// RandB generates a vector of N random bits
func RandB(N int) VectorB {
	result := NewVectorB(N)
	for i := 0; i < N; i++ {
		result[i] = uint8(rand.Intn(2))
	}
	return result
}

// RandReadableChars returns N printable random characters char=32 to 126
func RandReadableChars(N int) VectorB {
	/// 32 to 126
	result := NewVectorB(N)
	var startChar byte = 32
	for i := 0; i < N; i++ {
		result[i] = startChar + uint8(rand.Intn(94))
	}
	return result
}

func RandChars(size int) []uint8 {
	var result = make([]uint8, size)

	for i := 0; i < size; i++ {
		result[i] = uint8(rand.Intn(256))

	}
	return result
}

// Generates Normal (Gaussian) distributed complex number
func RandNC(variance float64) complex128 {

	var result complex128
	var Mean float64 = 0
	if Mean != 0 && variance != 1 {
		var StdDev float64 = math.Sqrt(variance)
		result = complex128(complex(rand.NormFloat64()*StdDev+Mean, rand.NormFloat64()*StdDev+Mean))

	} else {
		result = complex128(complex(rand.NormFloat64(), rand.NormFloat64()))
	}
	return result
}

// RandUC generates Uniformly  distributed complex number, Both real and imaginary part are uniformly distributed
func RandUC(variance float64) complex128 {
	var result complex128
	var Mean float64 = 0
	if Mean != 0 && variance != 1 {
		var StdDev float64 = math.Sqrt(variance)
		result = complex128(complex(rand.Float64()*StdDev+Mean, rand.Float64()*StdDev+Mean))

	} else {
		result = complex128(complex(rand.Float64(), rand.Float64()))
	}
	return result
}

// Generates Uniformly  distributed complex number
// Both real and imaginary part are uniformly distributed
func RandUFVec(size int) VectorF {

	result := NewVectorF(size)
	for i := 0; i < size; i++ {
		result[i] = rand.Float64()
	}

	return result
}

// Generates a complex vector with values uniformly distributed
func RandUCVec(samples int, variance float64) VectorC {

	result := NewVectorC(samples)
	for i := 0; i < samples; i++ {
		result[i] = RandUC(variance)
	}
	return result

}

//Generates a complex vector with values Normally distributed
func RandNCVec(samples int, variance float64) VectorC {

	result := NewVectorC(samples)
	for i := 0; i < samples; i++ {
		result[i] = RandNC(variance)
	}
	return result

}

// RandNF generates  random float values with distribution  N(0,var) - Guassian/Normal
func RandNF(variance float64) float64 {
	var result float64
	var Mean float64 = 0
	if Mean != 0 && variance != 1 {
		var StdDev float64 = math.Sqrt(variance)
		result = rand.NormFloat64()*StdDev + Mean

	} else {
		result = rand.NormFloat64()
	}
	return result
}

func RandNFVec(samples int, variance float64) VectorF {

	result := NewVectorF(samples)
	var Mean float64 = 0
	var StdDev float64 = math.Sqrt(variance)

	if Mean != 0 && variance != 1 {
		for i := 0; i < samples; i++ {
			result[i] = rand.NormFloat64()*StdDev + Mean
		}
	} else {

		for i := 0; i < samples; i++ {
			result[i] = rand.NormFloat64()
		}
	}
	return result
}

func RandUMatrix(rows, cols int) MatrixF {
	result := NewMatrixF(rows, cols)
	for i := 0; i < rows; i++ {
		input := RandUFVec(rows)

		result.SetCol(i, input)
	}

	return result
}

func RandNMatrix(rows, cols int, variance float64) MatrixF {
	result := NewMatrixF(rows, cols)
	for i := 0; i < rows; i++ {
		input := RandNFVec(rows, variance)

		result.SetCol(i, input)
	}

	return result
}

func RandNMatrixC(rows, cols int, variance float64) MatrixC {
	result := NewMatrixC(rows, cols)
	for i := 0; i < rows; i++ {
		input := RandUCVec(rows, variance)
		result[i] = input
	}

	return result
}
