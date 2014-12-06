package vlib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/cmplx"
	"os"
	"reflect"
)

type Vector []int
type VectorF []float64
type VectorB []uint8
type VectorI []int
type VectorC []complex128

// type VectorIface interface {
// }

// MarshalJSON()
// func (v *VectorF) MarshalJSON() ([]byte, error) {
// 	str := fmt.Sprintf("x=%f", v)
// 	return []byte(str), nil
// 	// return json.Marshal([]float64(v))
// }

type Point struct {
	X float64
	Y float64
}

type PointA []Point

func FromVecF(f VectorF) PointA {
	p := make(PointA, len(f))
	for indx, val := range f {
		p[indx].X = float64(indx)
		p[indx].Y = val
	}
	return p
}

func FromXYVecF(fx, fy VectorF) PointA {
	if len(fx) != len(fy) {
		return nil
	}
	p := make(PointA, len(fx))
	for indx, _ := range fx {
		p[indx].X = fx[indx]
		p[indx].Y = fy[indx]
	}
	return p
}

func FromVecCabs(f VectorC) PointA {
	// p := new(P/ointA)
	p := make(PointA, len(f))
	for i, val := range f {
		p[i].X = float64(i)
		p[i].Y = cmplx.Abs(val)
	}
	return p
}

func FromVecCreal(f VectorC) PointA {
	// p := new(PointA)
	p := make(PointA, len(f))
	for indx, val := range f {
		p[indx].X = float64(indx)
		p[indx].Y = real(val)
	}
	return p
}

func (p PointA) String() string {
	var result string
	Len := len(p)
	for indx, val := range p {
		result += fmt.Sprintf("[%f,%f]", val.X, val.Y)
		if indx != Len-1 {
			result += ","
		}
	}
	return result
}

// type Sequence struct {
// 	VectorI
// }

// func (s *Sequence) Create(str string) {
// 	s.VectorI = ToVectorI(str)
// }

func GenType(keyname string, val interface{}) map[string]interface{} {
	return map[string]interface{}{keyname: val}
}

func ToMap(in interface{}) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToMap only accepts structs; got %T", v)
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)
		if fi.Type.Name() == "complex128" {
			temp := v.Field(i).Interface().(complex128)
			// log.Println("Converting complex128 to Complex")
			// out[fi.Name] = fmt.Sprintf("%f%fi", real(temp), imag(temp))
			out[fi.Name] = fmt.Sprintf("%v", temp)
		} else {
			out[fi.Name] = v.Field(i).Interface()
		}

	}
	// x, _ := json.Marshal(out)
	// fmt.Printf("\n %v ", string(x))
	return out, nil
}

func SaveStructure(data interface{}, fname string, formated ...bool) {
	var doFormat bool = true
	if len(formated) > 0 {
		doFormat = formated[0]
	}
	output, err := json.Marshal(data)
	if err != nil {
		log.Println("Unable to Marshal it : Skip Saving to ", fname)
		log.Println("Unable to Marshal it : Err ", err)
		return
	}
	fd, ferr := os.Create(fname)
	if ferr != nil {
		log.Println("Unable to Create File  ", fname)
		return
	}

	if doFormat {
		fmt.Fprintf(fd, "%s", format(output))
		// return format(output)
	} else {
		fmt.Fprintf(fd, "%s", output)
		// return output
	}

}

func format(jbdata []byte) []byte {
	var data []byte
	buffer := bytes.NewBuffer(data)
	json.Indent(buffer, jbdata, "", "\t")
	return buffer.Bytes()
}

func Db(linearValue float64) float64 {
	return 10.0 * math.Log10(linearValue)
}

func InvDb(dBValue float64) float64 {
	return math.Pow(10, dBValue/10.0)
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

}

func ToDegree(radian float64) float64 {
	return radian * 180.0 / math.Pi
}
func ToRadian(degree float64) float64 {
	return degree * math.Pi / 180.0
}
