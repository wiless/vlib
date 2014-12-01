package vlib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

// MarshalJSON()
// func (v *VectorF) MarshalJSON() ([]byte, error) {
// 	str := fmt.Sprintf("x=%f", v)
// 	return []byte(str), nil
// 	// return json.Marshal([]float64(v))
// }

type VectorIface interface {
	//	String() string
}

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

type Matlab struct {
	filename    string
	datfilename string
	file        *os.File
	datfile     *os.File

	Silent   bool
	writer   io.Writer
	cmdqueue []string
	Keys     []string
	Encoder  *json.Encoder
	flushed  bool
	Json     bool
}

func (m *Matlab) SetWriter(w io.Writer) {
	m.writer = w
	m.file = nil
}

func (m *Matlab) SetFile(fname string) {
	m.filename = fname
	m.datfilename = fname + ".dat"
	/// Matlab File
	fd, err := os.Create(m.filename)
	if err != nil {
		log.Printf("Matlab Error : Creating %s : %v"+m.filename, err)
		return
	}
	log.Println("Matlab Opened : " + m.filename)
	m.file = fd

	/// Obj Data File
	fdjson, errjson := os.Create(m.datfilename)
	if errjson != nil {
		log.Printf("Matlab Error : Creating %s : %v"+m.datfilename, errjson)
		return
	}
	log.Println("Matlab Opened JSON : " + m.datfilename)
	m.datfile = fdjson

	m.Encoder = json.NewEncoder(m.datfile)
	m.file.WriteString("\n% =========== File Auto generated ===========\n ")
}
func (m *Matlab) SetDefaults() {
	m.Silent = true
}

func NewMatlab(fname string) *Matlab {
	result := Matlab{}
	result.SetFile(fname + ".m")
	// result.filename = fname
	result.Silent = true
	// fd, err := os.Create(result.filename)
	// // result.cmdqueue = make([]string)
	// if err != nil {
	// 	return nil
	// }
	// log.Println("Matlab Opened : " + result.filename)
	// result.file = fd
	// result.Encoder = json.NewEncoder(fd)

	return &result
}

func (m *Matlab) Close() error {

	if m.file == nil {
		return nil
	}
	log.Println("Matlab Closing : " + m.filename)
	m.Flush()
	return m.file.Close()
}
func (m *Matlab) Name() string {
	return m.filename
}

func (m *Matlab) Export(varname string, data interface{}) {

	// if m.Json {
	// 	m.ExportStruct(varname, data)
	// 	return
	// }
	var str string
	typestr := reflect.TypeOf(data).String()
	if !m.Silent {
		log.Printf("Matlab::Export(): %s(%s)", varname, reflect.TypeOf(data).String())
	}

	switch typestr {
	case "int":
		str = fmt.Sprintf("%s = %d;", varname, data)
	case "vlib.VectorF":
		str = fmt.Sprintf("%s = %v;", varname, data)
	case "[]string":
		str = fmt.Sprintf("%s = %v;", varname, data)
	default:
		str = fmt.Sprintf("%s = %f;", varname, data)
	}

	// if typestr == "int" {
	// 	str = fmt.Sprintf("%s = %d;", varname, data)
	// } else {
	// 	str = fmt.Sprintf("%s = %f;", varname, data)
	// }

	if !m.Silent {
		fmt.Printf("\n %s : %s", m.filename, str)
	}

	// os.Stdout.WriteString("\n ====== STD OUTPUT  WRITING  \n" + str)
	// log.Printf("Current status %v", m.file)
	if m.file != nil {
		m.file.WriteString("\n" + str)
	} else if m.writer != nil {

		// log.Printf("Wrote something in the Writer Object")
		databyte := []byte("\n" + str)
		m.writer.Write(databyte)
	}

}

func (m *Matlab) Command(cmd string) {
	if !m.Silent {
		fmt.Printf("\n %s : %s", m.filename, cmd)
	}

	// fmt.Printf("Current status %v", m.file)
	if m.file != nil {
		m.file.WriteString("\n" + cmd)
	} else if m.writer != nil {
		// fmt.Printf("Wrote something in the Writer Object")
		databyte := []byte("\n" + cmd)
		m.writer.Write(databyte)
	}

}

func (m *Matlab) Q(cmd string) {

	m.cmdqueue = append(m.cmdqueue, cmd)
	m.flushed = false
}

func (m *Matlab) Flush() {
	if !m.flushed {
		for i := 0; i < len(m.cmdqueue); i++ {
			m.Command(m.cmdqueue[i])
		}
		log.Println("Appending Keys ", m.Keys)
		m.ExportStruct("Keys", m.Keys)
	}
	m.flushed = true
}

func (m *Matlab) AddText(pos complex128, text string) string {
	x, y := real(pos), imag(pos)
	cmd := fmt.Sprintf("text(%f,%f,'%s')", x, y, text)
	return cmd
}

func (m *Matlab) ExportStruct(keyname string, val interface{}) {
	mapstruct := GenType(keyname, val)
	m.Keys = append(m.Keys, keyname)
	if m.Encoder != nil {

		if reflect.TypeOf(val).Kind() == reflect.Ptr {
			fmt.Printf("\n ExportStruct %v Type of %v", mapstruct, reflect.TypeOf(val).Kind())

		}

		err := m.Encoder.Encode(mapstruct)
		m.flushed = false
		if err != nil {
			log.Println("Error Writing ", err)
		}
	}

}

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
