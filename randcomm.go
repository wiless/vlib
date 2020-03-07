package vlib

import (
	"math/rand"
)

// RandBPSK Generates a VectorC from BPSK symbols
func RandBPSK(samples int, variance float64) VectorC {
	bpsksymbols := []complex128{complex(1, 0), complex(-1, 0)}
	result := NewVectorC(samples)
	for i := 0; i < samples; i++ {
		result[i] = bpsksymbols[rand.Intn(2)]
	}
	return result

}

// RandBPSK Generates a VectorC from BPSK symbols
func RandPI2BPSK(samples int, variance float64) VectorC {
	bpsksymbols := []complex128{complex(0, 1), complex(0, -1)}
	result := NewVectorC(samples)
	for i := 0; i < samples; i++ {
		result[i] = bpsksymbols[rand.Intn(2)]
	}
	return result

}

// RandBPSK Generates a VectorC from BPSK symbols
func RandQPSK(samples int, variance float64) VectorC {
	qpsksymbols := []complex128{complex(0.707, 0.707), complex(-0.707, 0.707), complex(0.707, -0.707), complex(-0.707, -0.707)}
	result := NewVectorC(samples)
	for i := 0; i < samples; i++ {
		result[i] = qpsksymbols[rand.Int31n(4)]
	}
	return result

}

// // RandBPSK Generates a VectorC from BPSK symbols
// func RandQAM16(samples int, variance float64) VectorC {
//	qam16symbols := []complex128{..}
// 	result := NewVectorC(samples)
// 	for i := 0; i < samples; i++ {
//result[i] = qam16symbols[rand.Int31n(16)]
// 	}
// 	return result

// }
