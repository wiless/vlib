package vlib

import (
	"fmt"
	"reflect"
)

type Vector []int
type VectorF []float64
type VectorB []uint8
type VectorI []int
type VectorC []complex128

type VectorIface interface {
	//	String() string
}

type Sequence struct {
	VectorI
}

func (s *Sequence) Create(str string) {
	s.VectorI = ToVectorI(str)
}

// type VectorIface interface{
// 	GetSize() int
// 	SetSize(int)
// }

type GIntVector struct {
	VectorIface
	size int
}

type GDoubleVector struct {
	VectorIface
	size int
}

type GComplexVector struct {
	VectorIface
	size int
}

func (v *GIntVector) SetSize(size int) {
	v.VectorIface = make([]int, size)
	v.size = size
}

func (v *GDoubleVector) SetSize(size int) {
	v.VectorIface = make([]float64, size)
	v.size = size
}

func (v *GComplexVector) SetSize(size int) {
	v.VectorIface = make([]complex128, size)
	v.size = size
}

func (v GDoubleVector) String() string {
	return toString(v.VectorIface)
	//return fmt.Sprintf("%v", v.VectorIface)

}

func (v GIntVector) String() string {
	return toString(v.VectorIface)
	//return fmt.Sprintf("%v", v.VectorIface)

}
func toString(v VectorIface) string {

	return fmt.Sprintf("%v", v)
}
func GetSize(v VectorIface) int64 {
	var size int64

	switch reflect.TypeOf(v).String() {
	// case reflect.Array:
	// 	fmt.Printf("its array", reflect.ValueOf(v).Len())
	// case reflect.Slice:
	// 	//size = reflect.ValueOf(v).Len()
	// 	fmt.Printf("Its Slice %v", reflect.ValueOf(v).Len())
	case "vlib.GIntVector", "vlib.GDoubleVector", "vlib.GComplexVector":
		//str := reflect.TypeOf(v).String()
		size = reflect.ValueOf(v).FieldByName("size").Int()
		//fmt.Printf("\n Its of Type %s", str)
		//fmt.Printf("\n Actual Object is %d", size)

	default:
		fmt.Printf("Unknown Type")
		size = -1

		// case GDoubleVector:
		// size = reflect.ValueOf(v).FieldByName("size").Int()
		// fmt.Printf("Its of Type %v", reflect.TypeOf(v))
		// fmt.Printf("Actual Object is %d", size)

	}

	return size
}

// func ElemMult(in1, in2 VectorIface) Vector {
// 	size := len(in1)
// 	result := New(size)

// 	for i := 0; i < size; i++ {
// 		result[i] = in1[i] * in2[i]
// 	}

// 	return result
// }
