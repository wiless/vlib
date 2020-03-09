package vlib

import (
	"fmt"
	"math"
	"math/rand"
)

func RandBitsF(size int) VectorF {
	var result = make([]float64, size)
	for i := 0; i < size; i++ {
		if rand.Float32() >= 0.5 {
			result[i] = 1.0
		}
	}
	return result
}

func RandI(size int, maxvalue int) []int {
	var result = make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = rand.Intn(maxvalue)

	}
	return result
}

func Randsrc(size int, maxvalue int) VectorI {
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
		result[i] = uint8(rand.Int63n(2))
	}
	return result
}

type keypair struct {
	begin  uint8
	length int
}

var splAsciitable [3]keypair = [3]keypair{{48, 10}, {65, 26}, {97, 26}}

// RandString returns N numbers or alphabets
func RandString(N int) VectorB {
	// start,end
	// 48,10 0-9
	// 65,27 A-Z
	// 97,26 a-z

	result := NewVectorB(N)
	for i := 0; i < N; i++ {
		key := rand.Intn(3)
		result[i] = splAsciitable[key].begin + uint8(rand.Intn(splAsciitable[key].length))
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
	// if Mean != 0 && variance != 1 {
	var StdDev float64 = math.Sqrt(variance)
	result = complex128(complex(rand.NormFloat64()*StdDev+Mean, rand.NormFloat64()*StdDev+Mean))

	// } else {
	// 	result = complex128(complex(rand.NormFloat64()*StdDev, rand.NormFloat64()*StdDev))
	// }
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

// Generates Uniformly distributed Float Vector of size size
// Values between 0 to 1 uses rand.Float64()
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
	for i := 0; i < cols; i++ {
		input := RandNFVec(rows, variance)
		fmt.Println(input, result)
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
