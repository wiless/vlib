// Package vlib provides some trivial functions for vector of int,float64,complex128 and bits. Each corresponding vector is extended
// from the standard array of data types. Hence it can be type-casted anytime and interface with other libraries
package vlib

import (

	// "os"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func (v VectorB) Size() int {
	return len(v)
}
func (v VectorB) Len() int {
	return v.Size()
}

func (v *VectorB) Resize(size int) {
	// Only append etc length
	length := len(*v)
	extra := (size - length)
	if extra > 0 {
		tailvec := NewVectorB(extra)
		*v = append(*v, tailvec...)
	}

	//copy(*v, Vector(make([]int, size)))

}

func (v VectorB) Get(indx int) uint8 {
	if indx < 0 || indx >= v.Len() {
		log.Panicln("VectorF::Get() Index out of Bounds.. ")
	}
	return v[indx]
}
func (v VectorB) At(indx VectorI) VectorB {

	result := NewVectorB(v.Size())
	for i := 0; i < v.Len(); i++ {

		result[i] = v.Get(indx[i])
	}
	return result
}

// Does elementwise XOR addition between vectors
func ElemAddB(in1, in2 VectorB) VectorB {
	size := len(in1)
	result := NewVectorB(size)

	for i := 0; i < size; i++ {
		// bool(in1[i])
		if in1[i] == 1 && in2[i] == 1 {
			result[i] = 0
		} else if in1[i] == 1 || in2[i] == 1 {
			result[i] = 1
		}

	}

	return result
}

// Does elementwise Multiplication (AND operator) addition between vectors
func ElemMultB(in1, in2 VectorB) VectorB {
	size := len(in1)
	result := NewVectorB(size)

	for i := 0; i < size; i++ {
		if in1[i] == 1 && in2[i] == 1 {
			result[i] = 1
		} else {
			result[i] = 0
		}

	}

	return result
}

func NewVectorB(size int) VectorB {
	return VectorB(make([]uint8, size))
}

func NewOnesB(size int) (v VectorB) {
	result := VectorB(make([]uint8, size))

	for i := 0; i < size; i++ {
		result[i] = 1
	}
	return result
}

func ToVectorB(str string) VectorB {

	str = strings.TrimSpace(str)
	var exp string = "[.|;&,: ]+"
	regx, _ := regexp.Compile(exp)
	bitstrlist := regx.Split(str, -1)
	result := NewVectorB(len(bitstrlist))

	for cnt, bitstr := range bitstrlist {
		// if bitstr != "" {
		bit, _ := strconv.ParseUint(bitstr, 10, 1)
		result[cnt] = uint8(bit)
		// fmt.Printf("\n %d , %s , %d => %d", cnt, bitstr, bit, result[cnt])
		// }

	}

	return result
}
