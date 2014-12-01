package vlib

import (
	"bufio"
	"fmt"
	"go/scanner"
	"go/token"
	"log"
	"math"
	"math/cmplx"
	"os"
	"strconv"
	"strings"
	// "os"
	// "regexp"
	// "strconv"
	// "strings"
)

// type Complex complex128
var Origin3D Location3D

type Location3D struct {
	X, Y, Z float64
}

func WrapAngle(degree0to360 float64) (degreePlusMinus180 float64) {
	//degree0to360 = math.Mod(degree0to360, 360)
	if degree0to360 > 180 {
		degree := math.Mod(degree0to360, 180)
		degreePlusMinus180 = -180 + degree

	} else if degree0to360 < -180 {
		degree := math.Mod(degree0to360, 180)
		degreePlusMinus180 = 180 + degree

	}
	log.Println("Input Output", degree0to360, degreePlusMinus180)
	return degreePlusMinus180
	// fmt.Println("Origina ", degree)
	// if degree > 180 {
	// 	rem := math.Mod(degree, 180.0)
	// 	degree = -180 + rem

	// } else if degree < -180 {
	// 	rem := math.Mod(degree, 180.0)
	// 	//	fmt.Println("Remainder for ", degree, rem)
	// 	degree = 180 + rem
	// }
}

func (l *Location3D) XY() complex128 {
	return complex(l.X, l.Y)
}

func (l *Location3D) XZ() complex128 {
	return complex(l.Z, l.X)
}

func (l *Location3D) SetHeight(height float64) {
	l.Z = height
}

func (l *Location3D) Cmplx() complex128 {
	return complex(l.X, l.Y)
}

func (l *Location3D) Shift3D(delta Location3D) {
	l.X += delta.X
	l.Y += delta.Y
	l.Z += delta.Z
}

func (l *Location3D) Shift2D(deltaxy complex128) {
	l.Shift3D(FromCmplx(deltaxy))
}

func (l *Location3D) SetLoc(loc2D complex128, height float64) {
	*l = FromCmplx(loc2D)
	l.SetHeight(height)
}

func FromVectorC(loc2d VectorC, height float64) []Location3D {
	result := make([]Location3D, loc2d.Size())

	for indx, val := range loc2d {
		result[indx] = FromCmplx(val)
		result[indx].SetHeight(height)
	}
	return result
}

func LoadLocationsFromFile(fcsvname, separator string) []Location3D {
	var result []Location3D
	fd, err := os.Open(fcsvname)
	if err != nil {
		log.Fatal("LoadLocationsFromFile():", fcsvname, "error ", err)
		return result
	}

	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		linestr := scanner.Text()
		var x, y, z float64
		fmt.Sscanf(linestr, "%f,%f,%f", &x, &y, &z)
		// fmt.Println(x, y, z)
		result = append(result, Location3D{x, y, z + 10})

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

func ToVectorC(locs []Location3D) VectorC {
	result := NewVectorC(len(locs))
	for indx, val := range locs {
		result[indx] = val.Cmplx()
	}
	return result
}

func FromCmplx(val complex128) Location3D {
	return Location3D{real(val), imag(val), 0}
}

func (l *Location3D) FromCmplx(val complex128) {
	l.X, l.Y = real(val), imag(val)

}

func (l *Location3D) Length() float64 {
	sum := math.Pow(l.X, 2)
	sum += math.Pow(l.Y, 2)
	sum += math.Pow(l.Z, 2)
	return math.Sqrt(sum)
}

func (l *Location3D) DistanceFrom(src Location3D) float64 {

	sum := math.Pow(l.X-src.X, 2)
	sum += math.Pow(l.Y-src.Y, 2)
	sum += math.Pow(l.Z-src.Z, 2)
	return math.Sqrt(sum)
}

func ToDegree(phase float64) float64 {
	return phase * 180.0 / math.Pi
}
func ToRadian(degree float64) float64 {
	return degree * math.Pi / 180.0
}

func RelativeGeo(src, dest Location3D) (distance3d, thetaH, thetaV float64) {

	thetaH = cmplx.Phase(dest.XY() - src.XY())

	r := cmplx.Abs(dest.XY() - src.XY())
	z := dest.Z - src.Z
	thetaV = cmplx.Phase(complex(r, z))
	// thetaV = math.Acos(z / r)
	distance3d = dest.DistanceFrom(src)
	return distance3d, ToDegree(thetaH), ToDegree(thetaV)
}

//func init() {
//	fmt.Printf("\n%% === Vlib Initialized : - github.com/wiless ===\n") /// matlab or octave compatible dump
//}

//// Functions for Complex Vectors
type Complex complex128

func (c Complex) MarshalJSON() ([]byte, error) {
	str := fmt.Sprintf("\"%+f%+fi\"", real(c), imag(c))
	return []byte(str), nil
}

func (c Complex) String() string {
	str := fmt.Sprintf("%f%fi", real(c), imag(c))
	return str
}

func (c VectorC) MarshalJSON() ([]byte, error) {
	// ParseCVec
	var str []string
	for _, val := range c {
		str = append(str, fmt.Sprintf("%f%+fi", real(val), imag(val)))
	}
	result := "\"[" + strings.Join(str, ",") + "]\""
	// log.Print(" \n JSONING vector ", result)
	// str := fmt.Sprintf("\"%+g%+gi\"", real(c), imag(c))
	return []byte(result), nil
}

func (c *Complex) UnmarshalJSON(data []byte) error {
	var re, im float64
	str := string(data)
	fmt.Sscanf(str, "\"%f%fi\"", &re, &im)
	// fmt.Printf("\nB4 READ : %s %f %f", str, re, im)
	*c = Complex(complex(re, im))
	return nil
}

func (c *VectorC) UnmarshalJSON(databyte []byte) error {
	// ParseCVec
	*c = ParseCVec(string(databyte))
	return nil
	// var str []string
	// for _, val := range c {
	// 	str = append(str, fmt.Sprintf("%g%+gi", real(val), imag(val)))
	// }
	// result := "\"[" + strings.Join(str, ",") + "]\""
	// log.Print(" \n JSONING vector ", result)
	// // str := fmt.Sprintf("\"%+g%+gi\"", real(c), imag(c))
	// return []byte(result), nil
}

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

func (v VectorC) AddC(arg complex128) VectorC {
	result := NewVectorC(len(v))
	for i, val := range v {
		result[i] = val + arg
	}
	return result
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

func (v VectorC) String() string {
	var result string
	size := v.Size()
	for i := 0; i < size; i++ {
		result += fmt.Sprintf("%f ", v[i])
	}
	return result
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

func (v VectorC) Real() VectorF {
	result := NewVectorF(v.Size())
	for indx, val := range v {
		result[indx] = real(val)
	}
	return result
}
func (v VectorC) Imag() VectorF {
	result := NewVectorF(v.Size())
	for indx, val := range v {
		result[indx] = imag(val)
	}
	return result
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

func ParseCVec(str string) VectorC {
	var result VectorC
	// result = make([]complex128, 1)
	// src := []byte("[ 3.141+4i,1+0i,20+4.4i]")
	src := []byte(str)
	// fmt.Println("About to Parse ", str)
	// Initialize the scanner.
	var s scanner.Scanner
	fset := token.NewFileSet()                      // positions are relative to fset
	file := fset.AddFile("", fset.Base(), len(src)) // register input "file"
	s.Init(file, src, nil /* no error handler */, scanner.ScanComments)

	// Repeated calls to Scan yield the token sequence found in the input.
	var cnt int = 0
	var recentNumber complex128
	var realval, imagval float64
	var indx int = 0
	var recentsign float64 = 1
	for {
		_, tok, lit := s.Scan()
		if tok == token.EOF {
			// fmt.Print("breaking at ", pos)
			break

		}

		if tok == token.ADD {
			recentsign = 1
		}
		if tok == token.SUB {
			recentsign = -1
		}
		if tok == token.INT || tok == token.FLOAT || tok == token.IMAG {

			// INT    // 12345
			//      FLOAT  // 123.45
			//      IMAG   // 123.45i

			/// NEW COMPLEX NUMBER
			if cnt%2 == 0 {
				/// Real part of complex number
				imagval = 0
				if tok == token.INT {
					realval = 0
					tmp, err := strconv.ParseInt(lit, 10, 32)
					if err == nil {
						realval = recentsign * float64(tmp)
					}

				}
				if tok == token.FLOAT {
					realval, _ = strconv.ParseFloat(lit, 64)
					realval = recentsign * realval
				}
				if tok == token.IMAG {
					// PURE IMAGINARY NUMBER
					fmt.Sscanf(lit, "%fi", &imagval)
					// realval = fmt.strconv.ParseFloat(tok, 64)
					realval = 0
					recentNumber = complex(0, recentsign*imagval)
					// fmt.Println(" PURE IMAGINARY ", indx, recentNumber)
					result = append(result, recentNumber)
					indx++
					cnt++
				}

				//recentNumber = complex(realval, imagval)
			} else {

				if tok == token.IMAG {

					fmt.Sscanf(lit, "%fi", &imagval)
					imagval = recentsign * imagval
					recentNumber = complex(realval, imagval)
					// fmt.Println(" Finished FULL COMPLEX NUMBER ", indx, recentNumber)
					result = append(result, recentNumber)
					indx++
					// realval = fmt.strconv.ParseFloat(tok, 64)
				} else {

					recentNumber = complex(realval, 0)
					result = append(result, recentNumber)
					// fmt.Println(" REAL number ", indx, recentNumber)
					indx++
					cnt++

					if tok == token.INT {
						realval = 0
						tmp, err := strconv.ParseInt(lit, 10, 32)
						if err == nil {
							realval = recentsign * float64(tmp)
						}

					}
					if tok == token.FLOAT {
						realval, _ = strconv.ParseFloat(lit, 64)
						realval = recentsign * realval
					}
				}
			}

			cnt++
			// recentNumber = complex(0, lit)
			/// Imag part of complex number
		}

		// fmt.Printf("%s \t %s \t %q \n", fset.Position(pos), tok, lit)
	}

	// fmt.Printf("Parsed vector %v", result)
	return result
}
