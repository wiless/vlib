package vlib

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	// "os"
	"regexp"
	"strconv"
	"strings"
)

//func init() {
//	// fmt.Printf("\n%% === Vlib Initialized : - github.com/wiless ===\n") /// matlab or octave compatible dump
//}

func ElemMult(in1, in2 Vector) Vector {
	size := len(in1)
	result := New(size)

	for i := 0; i < size; i++ {
		result[i] = in1[i] * in2[i]
	}

	return result
}

func (v VectorB) Size() int {
	return len(v)
}

func (v VectorB) ToString() string {
	var result string
	size := v.Size()
	for i := 0; i < size; i++ {
		result += fmt.Sprintf("%d", v[i])
	}
	return result
}

func (v VectorB) CountErrors(in1 VectorB) int {
	size := len(v)
	if size != len(in1) {
		return size
	}
	return (ElemAddB(v, in1)).NNZ()

}

func (v VectorB) IsEqual(in1 VectorB) bool {
	size := len(v)
	if size != len(in1) {
		return false
	}
	if (ElemAddB(v, in1)).NNZ() == 0 {
		return true
	} else {
		return false
	}

}
func (v VectorB) NNZ() int {
	size := len(v)
	var result int = 0

	for i := 0; i < size; i++ {
		if v[i] == 1 {
			result++
		}

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

// func Sum(VectorB) VectorB {

func InvDbF(in1 VectorF) VectorF {

	result := NewVectorF(len(in1))
	for i, val := range in1 {
		result[i] = InvDb(val)
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

func ElemDivF(in1, in2 VectorF) VectorF {

	size := len(in1)
	result := NewVectorF(size)

	for i := 0; i < size; i++ {
		result[i] = in1[i] / in2[i]
	}

	return result
}

func ElemMultF(in1, in2 VectorF) VectorF {
	size := len(in1)
	result := NewVectorF(size)

	for i := 0; i < size; i++ {
		result[i] = in1[i] * in2[i]
	}

	return result
}

func New(size int) Vector {
	return Vector(make([]int, size))
}

func NewVectorF(size int) VectorF {
	return VectorF(make([]float64, size))
}

func NewVectorI(size int) VectorI {
	return VectorI(make([]int, size))
}

func NewSegmentI(begin, size int) VectorI {
	var result VectorI
	result = make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = begin + i
	}
	return result
}

func NewVectorB(size int) VectorB {
	return VectorB(make([]uint8, size))
}

func (v Vector) Size() int {
	return len(v)
}

func (v VectorI) Size() int {
	return len(v)
}

func (v *Vector) Resize(size int) {
	// Only append etc length
	length := v.Size()
	extra := (size - length)
	if extra > 0 {
		tailvec := New(extra)
		*v = append(*v, tailvec...)
	}
	///copy(*v, Vector(make([]int, size)))
}

func (v VectorF) Clone() VectorF {
	result := NewVectorF(v.Size())
	copy(result, v)
	return result
}

func (v *VectorF) Resize(size int) {
	// Only append etc length
	length := len(*v)
	extra := (size - length)
	if extra > 0 {
		tailvec := NewVectorF(extra)
		*v = append(*v, tailvec...)
	}

	///copy(*v, Vector(make([]int, size)))
}

func (v *VectorI) Resize(size int) {
	// Only append etc length
	length := len(*v)
	extra := (size - length)
	if extra > 0 {
		tailvec := NewVectorI(extra)
		*v = append(*v, tailvec...)
	}

	///copy(*v, Vector(make([]int, size)))

}

func NewOnes(size int) (v Vector) {

	result := Vector(make([]int, size))

	for i := 0; i < size; i++ {
		result[i] = 1
	}

	return result
}
func NewOnesF(size int) (v VectorF) {
	result := VectorF(make([]float64, size))

	for i := 0; i < size; i++ {
		result[i] = 1
	}
	return result
}

func NewOnesB(size int) (v VectorB) {
	result := VectorB(make([]uint8, size))

	for i := 0; i < size; i++ {
		result[i] = 1
	}
	return result
}

func (v *Vector) Ones() {
	for i := 0; i < len(*v); i++ {
		(*v)[i] = 1
	}

}

func (v *VectorF) OnesF() {
	for i := 0; i < len(*v); i++ {
		(*v)[i] = 1
	}
}

func (v Vector) ScaleInt(factor int) Vector {

	VectorF := New(v.Size())
	for indx, val := range v {
		VectorF[indx] = val * factor
	}
	return VectorF
}

func (v VectorF) Scale(factor float64) VectorF {

	VectorF := NewVectorF(len(v))
	for indx, val := range v {
		VectorF[indx] = val * factor
	}
	return VectorF
}

func (v Vector) Scaleloat64(factor float64) VectorF {

	VectorF := NewVectorF(v.Size())
	for indx, val := range v {
		VectorF[indx] = float64(val) * factor
	}
	return VectorF
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

func ToVectorF(str string) VectorF {
	var v VectorF
	if strings.Contains(str, ":") {
		// fmt.Printf("Input String : %s ", str)
		result := strings.Split(str, ":")
		start, _ := strconv.ParseFloat(result[0], 64)
		// start, _ := strconv.ParseFloat(s, bitSize) (result[0])

		step := 1.0
		end := start
		var Len int
		switch len(result) {
		case 2:
			end, _ = strconv.ParseFloat(result[1], 64)
			Len = int(math.Floor(float64((end - start)) + 1))

		case 3:
			step, _ = strconv.ParseFloat(result[1], 64)
			end, _ = strconv.ParseFloat(result[2], 64)
			diffs := int(math.Abs(float64((end - start) / step)))

			if step < 0 {
				// tmp := start
				// start = end
				// end = tmp

			}

			Len = int(math.Floor(float64(diffs)) + 1)

		}

		if step < 0 {
			Len = Len

		}
		v.Resize(Len)
		// fmt.Printf("\n %v %v %v %v", start, step, end, Len)
		for k := range v {

			v[k] = start + float64(k)*step
		}

	}
	return v
}

func (v VectorF) SubV(rhs VectorF) VectorF {
	return Sub(v, rhs)
}

func (v VectorF) Sub(offset float64) VectorF {
	return v.Add(-offset)
}

func (v VectorF) Add(offset float64) VectorF {
	result := NewVectorF(len(v))
	for k := range v {
		result[k] = v[k] + offset
	}
	return result
}

func (v VectorF) Size() int {
	return len(v)
}
func (v VectorF) Len() int {
	return v.Size()
}
func (v VectorF) Less(i, j int) bool {
	return v[i] <= v[j]

}
func (v *VectorF) Swap(i, j int) {
	(*v)[i], (*v)[j] = (*v)[j], (*v)[i]
}

func (v VectorF) Find(x float64) int {
	result := -1
	for i, val := range v {
		if x == val {
			return i
		}
	}
	return result
}

/// Assumes descending ordered vector
func (v VectorF) FindSorted(x float64) int {
	result := -1
	length := v.Size()
	for i := (length - 1); i >= 0; i-- {

		val := v[i]
		if val > x {
			break
		}
		if val == x {
			return i
		}

	}
	return result
}

func (v VectorF) Get(indx int) float64 {
	if indx < 0 || indx >= v.Len() {
		log.Panicln("VectorF::Get() Index out of Bounds.. ")
	}
	return v[indx]
}

func (v VectorF) At(indx VectorI) VectorF {

	result := NewVectorF(v.Size())
	for i := 0; i < v.Len(); i++ {

		result[i] = v.Get(indx[i])
	}
	return result
}

func (v VectorI) Get(indx int) int {
	if indx < 0 || indx >= v.Size() {
		log.Panicln("VectorF::Get() Index out of Bounds.. ")
	}
	return v[indx]
}

func (v VectorI) At(indx VectorI) VectorI {

	result := NewVectorI(v.Size())
	for i := 0; i < v.Size(); i++ {

		result[i] = v.Get(indx[i])
	}
	return result
}

func (v *VectorF) PlusEqual(input VectorF) {
	if len(*v) != len(input) {
		log.Panicf("\n PlusEqual %d : Length Mismatch %d", v.Size(), input.Size())

	}
	cnt := v.Size()
	for k := 0; k < cnt; k++ {

		(*v)[k] = (*v)[k] + input[k]
	}

}

func (v *VectorF) AppendAtEnd(val ...float64) {
	// for i := 0; i < len(val); i++ {
	*v = append(*v, val...)
	// }

}

func (v *VectorI) AppendAtEnd(val ...int) {
	// for i := 0; i < len(val); i++ {
	*v = append(*v, val...)
	// }

}

func (v VectorF) IsEq(vals VectorF) bool {
	if v.Size() != vals.Size() {
		return false
	}

	var eps float64 = 1.0e-10

	for indx, val := range v {

		diff := val - vals[indx]
		errorval := math.Abs(diff)
		//log.Println("\n ELEMENT ", indx, "error ", errorval)
		if errorval > eps {
			// i := indx
			// log.Println("\nrow ", i, vals[i], "not match", val)
			//log.Println("\n ELEMENT ", indx, "error ", errorval)
			return false
		}

	}
	return true
}

func (v *VectorF) SetSubVec(pos int, vals VectorF) {
	loc := *v
	copy(loc[pos:], vals)
}

func (v VectorF) Insert(pos int, val float64) VectorF {
	result := NewVectorF(v.Size() + 1)
	copy(result[0:pos], v[0:pos])
	result[pos] = val
	copy(result[pos+1:], v[pos:])
	return result
}

func (v VectorI) Delete(pos int) VectorI {

	result := NewVectorI(v.Size())
	copy(result, v)
	copy(result[pos:], result[pos+1:])

	return result[:v.Size()-1]

}
func (v VectorF) Delete(pos int) VectorF {

	result := NewVectorF(v.Size())
	copy(result, v)
	copy(result[pos:], result[pos+1:])

	// result[v.Size()-1] = nil // or the zero value of T

	return result[:v.Size()-1]

}

func (v VectorF) AddVector(input VectorF) VectorF {
	result := NewVectorF(len(v))
	if v.Size() != input.Size() {
		log.Panicf("\nAddVector %d : Length Mismatch %d", v.Size(), input.Size())
	}
	for k := range v {
		result[k] = v[k] + input[k]
	}
	return result
}

func Dot(input1 VectorF, input2 VectorF) float64 {
	if input1.Size() != input2.Size() {
		log.Panicf("Dot: LHS (%d) RHS (%d) size mismatch", input1.Size(), input2.Size())
	}
	temp := ElemMultF(input1, input2)
	var result float64 = 0.0
	for _, val := range temp {
		result += val
	}
	return result
}

func Flip(input VectorF) VectorF {

	// var result Vec
	size := len(input)
	result := NewVectorF(size)

	/// short loop method-1
	for indx, val := range input {
		result[size-indx-1] = val
	}
	/// short loop method-2
	// copy(result, input)

	// for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
	// 	result[i], result[j] = result[j], result[i]
	// }
	return result
}

func Sum(v VectorF) float64 {
	var result float64
	for _, val := range v {
		result += val
	}
	return result

}

func Min(v VectorF) float64 {
	var result float64
	if v.Size() < 1 {
		return math.NaN()
	}
	result = v[0]
	for _, val := range v {
		result = math.Min(result, val)
	}
	return result
}

func Max(v VectorF) float64 {
	var result float64
	if v.Size() < 1 {
		return math.NaN()
	}
	result = v[0]
	for _, val := range v {
		result = math.Max(result, val)
	}
	return result
}

func (v VectorF) ShiftAndScale(shift, scale float64) VectorF {

	// v = v.Add(shift)
	// result := v.Scale(factor)

	result := NewVectorF(v.Size())
	for i := 0; i < v.Size(); i++ {
		result[i] = (v[i] + shift) * scale
	}
	return result
}
func (v VectorF) ScaleAndShift(shift, scale float64) VectorF {

	// v = v.Add(shift)
	// result := v.Scale(factor)

	result := NewVectorF(v.Size())
	for i := 0; i < v.Size(); i++ {
		result[i] = v[i]*scale + shift
	}
	return result
}

func (v *VectorF) Zeros() {
	v.Fill(0)
}

func (v *VectorF) Ones() {
	v.Fill(1)
}

func (v *VectorF) Fill(val float64) {
	for i := 0; i < v.Size(); i++ {
		(*v)[i] = val
	}
}

func MeanAndVariance(v VectorF) (mean, variance float64) {
	mean = Sum(v) / float64(v.Size())
	variance = 0
	for i := 0; i < v.Size(); i++ {
		variance += math.Pow(v[i], 2.0)
	}
	variance = variance / float64(v.Size()-1)

	return mean, variance
}

func Mean(v VectorF) float64 {

	return Sum(v) / float64(v.Size())

}

/// returns Euclidean Norm of the vector
func Variance(v VectorF) float64 {
	var result float64 = 0
	mean := Mean(v)
	for i := 0; i < v.Size(); i++ {
		result += math.Pow(v[i]-mean, 2.0)
	}

	return result / float64(v.Size()-1)

}

/// returns the sum of square of the elements in the vector
func Energy(v VectorF) float64 {
	var result float64 = 0
	for i := 0; i < v.Size(); i++ {
		result += math.Pow(v[i], 2.0)
	}

	return result
}

/// returns 2nd Norm of the vector (\sum(x[i]))^(1/2)
func Norm2(v VectorF) float64 {
	var result float64 = 0
	for i := 0; i < v.Size(); i++ {
		result += math.Pow(v[i], 2.0)
	}

	return math.Sqrt(result)

}

/// returns Euclidean Norm of the vector
func Norm(v VectorF) float64 {
	var result float64 = 0
	for i := 0; i < v.Size(); i++ {
		result += math.Pow(v[i], 2.0)
	}

	return math.Sqrt(result / float64(v.Size()))

}

func Add(A, B VectorF) VectorF {
	result := NewVectorF(A.Size())
	for i := 0; i < A.Size(); i++ {
		result[i] = A[i] + B[i]
	}
	return result
}

func Sub(A, B VectorF) VectorF {
	if A.Size() != B.Size() {
		log.Panicf("Sub: LHS (%d) and RHS (%d) size mismatch", A.Size(), B.Size())
	}
	result := NewVectorF(A.Size())
	for i := 0; i < A.Size(); i++ {
		result[i] = A[i] - B[i]
	}
	return result
}

/// Normalizes with 0 mean, and unit variance
func (v VectorF) Normalize() (result VectorF, mean, factor float64) {

	mean, variance := MeanAndVariance(v)
	factor = 1.0 / math.Sqrt(variance)
	result = v.ShiftAndScale(-mean, factor)

	// v = v.Sub(mean)
	// result = v.Scale(factor)

	return result, mean, factor
}

func ToVectorI(str string) VectorI {
	var v VectorI
	if strings.Contains(str, ":") {
		// fmt.Printf("Input String : %s ", str)
		result := strings.Split(str, ":")
		start, _ := strconv.ParseFloat(result[0], 64)
		// start, _ := strconv.ParseFloat(s, bitSize) (result[0])

		step := 1.0
		end := start
		var Len int
		switch len(result) {
		case 2:
			end, _ = strconv.ParseFloat(result[1], 64)
			Len = int(math.Floor(float64((end - start)) + 1))

		case 3:
			step, _ = strconv.ParseFloat(result[1], 64)
			end, _ = strconv.ParseFloat(result[2], 64)
			diffs := int(math.Abs(float64((end - start) / step)))

			if step < 0 {
				// tmp := start
				// start = end
				// end = tmp

			}

			Len = int(math.Floor(float64(diffs)) + 1)

		}

		if step < 0 {
			Len = Len

		}
		v.Resize(Len)
		// fmt.Printf("\n %v %v %v %v", start, step, end, Len)
		for k := range v {

			v[k] = int(start + float64(k)*step)
		}

	}
	return v

}

func (c *VectorI) MarshalJSON() ([]byte, error) {
	//// ParseCVec
	// var intarray []int
	// intarray = []int(*c)
	// res, err := json.Marshal(intarray)
	var str string
	for _, val := range *c {
		str += fmt.Sprintf("%d,", val)
	}

	str = "[" + strings.TrimSuffix(str, ",") + "]"
	// fmt.Println("Marshal : VectorI", str)
	return []byte(str), nil
	// 	// var str []string
	// 	// for _, val := range c {
	// 	// 	str = append(str, fmt.Sprintf("%f%+fi", real(val), imag(val)))
	// 	// }
	// 	// result := "\"[" + strings.Join(str, ",") + "]\""
	// 	// log.Print(" \n JSONING vector ", result)
	// 	// str := fmt.Sprintf("\"%+g%+gi\"", real(c), imag(c))
	// 	// return []byte(result), nil
}

// func (c *VectorI) UnmarshalJSON(databyte []byte) error {
// 	// ParseCVec
// 	var floatarray []float64
// 	err := json.Unmarshal(databyte, floatarray)
// 	fmt.Println("Unmarshal : VectorI", err, floatarray)

// 	// *c = ParseCVec(string(databyte))
// 	return err

// }

func (c *VectorI) Decode(databyte []byte) error {
	// ParseCVec
	var floatarray []float64
	err := json.Unmarshal(databyte, floatarray)
	fmt.Println("Unmarshal : VectorI", err, floatarray)

	// *c = ParseCVec(string(databyte))
	return err

}
