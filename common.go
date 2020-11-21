package vlib

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/cmplx"
	"os"
	"reflect"
	"sort"
	"strings"

	ms "github.com/mitchellh/mapstructure"
)

type Vector []int
type VectorF []float64
type VectorB []uint8
type VectorI []int
type VectorC []complex128
type VectorBool []bool

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

func GetEJtheta(degree float64) complex128 {
	return cmplx.Exp(complex(0.0, -degree*math.Pi/180.0))
}

func Radian(degree float64) float64 {
	return degree * math.Pi / 180.0
}

// Vector2D is a  XYer interface based on two VectorF
type Vector2D struct {
	X, Y VectorF
}

// XY returns X and Y value of the sample i
func (v Vector2D) XY(i int) (X, Y float64) {
	return v.X[i], v.Y[i]
}

// Len returns the length of the XY
func (v Vector2D) Len() int {
	return v.X.Len()
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

func Contains(array interface{}, elem interface{}) (found bool, index int) {
	// if len(array) == 0 {
	// 	return false, -1
	// }
	tOfArray := reflect.TypeOf(array)
	tOfelem := reflect.TypeOf(elem)
	if tOfArray.Kind() == reflect.Slice || tOfArray.Kind() == reflect.Array {
		// fmt.Printf("\n Argument 1 is %s of Types %s ", tOfArray.Kind(), tOfArray.Elem())
		if tOfArray.Elem() == tOfelem {
			avalue := reflect.ValueOf(array)
			evalue := reflect.ValueOf(elem)
			// fmt.Printf("\n Success : %s : %s : %v", tOfArray.Kind(), tOfelem.Kind(), avalue.Len())
			// result := false
			// myarray := reflect.ValueOf(array)
			for i := 0; i < avalue.Len(); i++ {

				var found bool
				switch tOfelem.Kind() {

				case reflect.String:
					found = avalue.Index(i).String() == evalue.String()

				case reflect.Int:
					found = avalue.Index(i).Int() == evalue.Int()

				case reflect.Float64:
					found = avalue.Index(i).Float() == evalue.Float()
				default:
					log.Panicln("vlib.Contains(): Unsupported ", tOfelem.Kind().String())
				}
				// fmt.Println("looking for ", evalue, " IN ", avalue.Index(i), avalue.Index(i).String() == evalue.String())
				if found {
					// fmt.Println("Found ")
					return true, i
				}
			}
			return false, -1

		} else {
			log.Panicf("Contains : MismatchType :  %v in %v ", elem, array)
		}
	}

	return false, -1
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

func LoadMapStructure(fname string, data interface{}) {

	r, ferr := os.Open(fname)
	if ferr != nil {
		log.Panicln("LoadMapStructure() : Unable to Open file ", fname)
	}

	// map[int]
	type Obj struct {
		ObjectID interface{}
		Object   map[string]interface{}
	}
	var objs []Obj

	dec := json.NewDecoder(r)
	derr := dec.Decode(&objs)
	if derr != nil {
		log.Panicln("LoadMapStructure():Unable to Decode json data", derr)

	}

	// fmt.Printf("\n Create Map [ %v ] %v  ", reflect.TypeOf(data).Key(), reflect.TypeOf(data).Elem())
	// vdata = reflect.MapOf(reflect.TypeOf(objs[0]))
	// fmt.Println("Try to Load array of %v ", reflect.TypeOf(data).Elem())
	vdata := reflect.ValueOf(data)
	// fmt.Println("%v ", vdata)
	for i := 0; i < len(objs); i++ {
		// fmt.Printf("\n%v : %v ", i, objs[i].Object)
		key := reflect.ValueOf(objs[i].ObjectID).Convert(reflect.TypeOf(data).Key())
		// var val reflect.Value
		val := reflect.Indirect(reflect.New(reflect.TypeOf(data).Elem()))
		derr := ms.Decode(objs[i].Object, val.Addr().Interface())
		if derr != nil {
			log.Panicln("LoadMapStructure() : Unable to Map to Structure")
		}
		// fmt.Println("dest obj ", val)
		// fmt.Println("Map2Str Error ", derr)
		// fmt.Printf("\n key %v MS : %v to %#v \n", key, objs[i].ObjectID, val.Interface())
		// fmt.Println("Setting to MAP")
		// rval := reflect.ValueOf(val)
		// oldval := vdata.MapIndex(key)
		// v = make(reflect.Value)0
		// fmt.Printf("\nOld value in Map %#v", val.Interface())

		vdata.SetMapIndex(key, reflect.Indirect(val))
	}
	// mdata := reflect.ValueOf(data)
	// var cnt int = 0
	// keys := mdata.MapKeys()

	// objs := make([]Obj, mdata.Len())

	// for _, val := range keys {
	// 	objs[cnt].ObjectID = val.Interface()
	// 	objs[cnt].Object = mdata.MapIndex(val).Interface()
	// 	//	fmt.Printf("\n Key  %v : Value %v", val.Int(), mdata.MapIndex(val))
	// 	cnt++
	// }
	// SaveStructure(objs, fname, formated...)
	// fmt.Println("\n", objs)
}

// func SaveMapStructure(data interface{}, fname, keyname, valname string, formated ...bool) {
// }
type Obj struct {
	ObjectID interface{}
	Object   interface{}
	keyName  string
	valname  string
}

func (o *Obj) init(k, v string) {
	o.keyName = k
	o.valname = v
}
func (o Obj) MarshalJSON() ([]byte, error) {
	bfr := bytes.NewBuffer(nil)
	enc := json.NewEncoder(bfr)

	bfr.WriteString("{")
	bfr.WriteString(fmt.Sprintf(`"%s":`, o.keyName))
	enc.Encode(o.ObjectID)
	bfr.WriteString(fmt.Sprintf(`,"%s":`, o.valname))

	enc.Encode(o.Object)

	bfr.WriteString("}")

	return bfr.Bytes(), nil
}

func (o *Obj) UnmarshalJSON(bfr []byte) error {

	dec := json.NewDecoder(bytes.NewReader(bfr))
	temp := make(map[string]interface{})
	derr := dec.Decode(&temp)
	// log.Println(derr)
	// log.Println("Decoded temp ", temp)

	return derr
}

func SaveMapStructure2(data interface{}, fname, keyname, valname string, formated ...bool) {
	// map[int]

	mdata := reflect.ValueOf(data)

	var cnt int = 0
	keys := mdata.MapKeys()

	objs := make([]Obj, mdata.Len())

	for _, val := range keys {
		objs[cnt].ObjectID = val.Interface()

		objs[cnt].Object = mdata.MapIndex(val).Interface()
		// fmt.Printf("\n Creating object %#v", objs[cnt].Object)
		objs[cnt].init(keyname, valname)
		//	fmt.Printf("\n Key  %v : Value %v", val.Int(), mdata.MapIndex(val))
		cnt++
	}
	SaveStructure(objs, fname, formated...)
	// fmt.Println("\n", objs)

}

func SaveMapStructure(data interface{}, fname string, formated ...bool) {
	SaveMapStructure2(data, fname, "ObjectID", "Object", formated...)
}

func LoadStructure(fname string, data interface{}) {
	dbytes, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Panicln("Error Reading File  ", fname, err)
	}
	// var newsystem deployment.DropSystem
	jerr := json.Unmarshal(dbytes, data)
	if jerr != nil {
		log.Panicln("Unable to UnMarshal ", jerr)
	}
}

func GetIntKeys(data interface{}) VectorI {
	if reflect.TypeOf(data).Kind() != reflect.Map {

	}
	vdata := reflect.ValueOf(data)
	result := NewVectorI(vdata.Len())
	cnt := 0
	for key, _ := range vdata.MapKeys() {
		result[cnt] = key
		cnt++
	}
	return result
}

func SaveStructure(data interface{}, fname string, formated ...bool) {

	if reflect.TypeOf(data).Kind() == reflect.Map {
		keyname := "key"
		elemName := "Value"

		// For the KeyType

		mtype := reflect.TypeOf(data)

		keytype := mtype.Key()
		valtype := mtype.Elem()
		if keytype.Kind() == reflect.Ptr {
			keytype = keytype.Elem()
		}

		if keytype.PkgPath() == "" {
			// Its a standard type
			keyname = "ID"
		} else {
			keyname = keytype.String()
		}
		// For the Value
		if valtype.Kind() == reflect.Ptr {
			valtype = valtype.Elem()
		}
		if valtype.PkgPath() == "" {
			// Its a standard type
			elemName = valtype.String()
			elemName = "Value" //;  + strings.TrimLeft(elemName, "*")
		} else {
			elemName = valtype.String()
			elemName = elemName[strings.LastIndex(elemName, ".")+1:]
			elemName = strings.TrimLeft(elemName, "*")
		}
		SaveMapStructure2(data, fname, keyname, elemName, formated...)
		return
	}

	var doFormat bool = true
	if len(formated) > 0 {
		doFormat = formated[0]
	}
	output, err := json.Marshal(data)

	if err != nil {
		log.Println("vlib:SaveStructure():Unable to Marshal it : Err ", err)
		return
	}
	fd, ferr := os.Create(fname)
	if ferr != nil {
		log.Println("vlib:SaveStructure():Unable to Create File  ", fname)
		return
	}

	if doFormat {
		fmt.Fprintf(fd, "%s", format(output))
		// return format(output)
	} else {
		fmt.Fprintf(fd, "%s", output)
		// return output
	}

	// fmt.Println("SUCCESS ==============================", fname)
}

func ModInt(number, N int) int {
	return int(math.Mod(float64(number), float64(N)))
	// return number / N
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
func Sinc(x float64) float64 {
	if x == 0 {
		return 1
	} else {
		return math.Sin(x) / x
	}

}

func SincF(x VectorF) VectorF {
	r := NewVectorF(x.Len())
	for i, v := range x {
		r[i] = Sinc(v)
	}
	return r
}

func ToDegree(radian float64) float64 {
	return radian * 180.0 / math.Pi
}
func ToRadian(degree float64) float64 {
	return degree * math.Pi / 180.0
}

func Sorted(data VectorF) (values VectorF, indx VectorI) {

	result := data.Clone()

	sort.Sort(sort.Reverse(sort.Float64Slice(result)))
	indx = NewVectorI(data.Len())
	for i, v := range data {
		indx[i] = result.FindSorted(v)

		// fmt.Printf("\n %v Found %f @ %d ", result, v, indx[i])
		// 		if indx[i] < len(data) && data[i] == x {
		// 	// x is present at data[i]
		// } else {
		// 	// x is not present in data,
		// 	// but i is the index where it would be inserted.
		// }

	}
	return VectorF(result), indx
}

func IterateF(input VectorF, myfunc func(float64) float64) VectorF {
	result := NewVectorF(input.Size())
	for indx, arg := range input {
		result[indx] = myfunc(arg)
	}
	return result

}

func IsTypeNumeric(t reflect.Type) bool {
	if t.Kind() >= reflect.Bool && t.Kind() <= reflect.Complex128 {
		return true
	}
	return false
}

func IsTypeString(t reflect.Type) bool {
	if t.Kind() == reflect.String {
		return true
	}
	return false
}

// Internal utility to convert a struct into array of strings
func StructNum2Strings(a interface{}) ([]string, error) {
	if reflect.TypeOf(a).Kind() != reflect.Struct {
		return nil, errors.New("Input Data not of Type Struct")
	}

	mvalue := reflect.ValueOf(a)
	mtype := reflect.TypeOf(a)
	var result []string = make([]string, 0, mtype.NumField())
	cnt := 0
	for i := 0; i < mtype.NumField(); i++ {
		// fmt.Println("Kind of field ", mtype.Field(i).Type.Kind())
		if mvalue.Field(i).CanInterface() && mtype.Field(i).Type.Kind() != reflect.Slice {

			if IsTypeNumeric(mtype.Field(i).Type) {
				result = append(result, fmt.Sprintf("%v", mvalue.Field(i).Interface()))
			}
			// } else if IsTypeString(mtype.Field(i).Type) {
			// 	result = append(result, fmt.Sprintf("%v", mvalue.Field(i).Interface()))
			// }

			cnt++
		}
	}
	return result, nil
}

// Internal utility to convert a struct into array of strings
func Struct2Strings(a interface{}) ([]string, error) {
	if reflect.TypeOf(a).Kind() != reflect.Struct {
		return nil, errors.New("Input Data not of Type Struct")
	}

	mvalue := reflect.ValueOf(a)
	mtype := reflect.TypeOf(a)
	var result []string = make([]string, 0, mtype.NumField())
	cnt := 0
	for i := 0; i < mtype.NumField(); i++ {
		// fmt.Println("Kind of field ", mtype.Field(i).Type.Kind())
		if mvalue.Field(i).CanInterface() && mtype.Field(i).Type.Kind() != reflect.Slice {
			result = append(result, fmt.Sprintf("%v", mvalue.Field(i).Interface()))
			cnt++
		}
	}
	return result, nil
}

// Internal utility to convert a struct into array of strings
func Struct2Header(a interface{}) ([]string, error) {
	if reflect.TypeOf(a).Kind() != reflect.Struct {
		return nil, errors.New("Input Data not of Type Struct")
	}

	mtype := reflect.TypeOf(a)
	var result []string = make([]string, 0, mtype.NumField())
	cnt := 0

	for i := 0; i < mtype.NumField(); i++ {
		if mtype.Field(i).PkgPath == "" && mtype.Field(i).Type.Kind() != reflect.Slice {
			// fmt.Println("Kind of field ", mtype.Field(i).Type.Kind())
			result = append(result, mtype.Field(i).Name)
			cnt++
		}
	}
	return result, nil
}

func DumpMap2CSV(fname string, arg interface{}) {

	if !(reflect.TypeOf(arg).Kind() == reflect.Map || reflect.TypeOf(arg).Kind() == reflect.Slice) {
		log.Println("Unable to Dump: Not Map or Struct interface")
		return
	}

	arrayData := reflect.ValueOf(arg)

	w, fer := os.Create(fname)
	if fer != nil {
		log.Print("Error Creating CSV file ", fer)
	}
	cwr := csv.NewWriter(w)
	// var record [4]string

	cwr.Comma = ','

	if reflect.TypeOf(arg).Kind() == reflect.Map {

		mapkeys := arrayData.MapKeys()
		once := true
		for _, key := range mapkeys {
			metric := arrayData.MapIndex(key).Interface()

			if once {
				headers, _ := Struct2Header(metric)
				w.WriteString(strings.Join(headers, ",") + "\n")
				once = false
			}
			data, _ := Struct2Strings(metric)
			cwr.Write(data)
		}
	}

	if reflect.TypeOf(arg).Kind() == reflect.Slice {
		tp := reflect.TypeOf(arg).Elem()

		var headers string
		for i := 0; i < tp.NumField(); i++ {
			headers += tp.FieldByIndex([]int{i}).Name + ","
		}
		w.WriteString(headers + "\n")

		for i := 0; i < arrayData.Len(); i++ {

			metric := arrayData.Index(i).Interface()
			data, _ := StructNum2Strings(metric)
			cwr.Write(data)
		}
	}
	cwr.Flush()
	w.Close()
}

// DumpMap2CSV2 dumps all the fields including numbers and strings to CSV
func DumpMap2CSV2(fname string, arg interface{}) {

	if !(reflect.TypeOf(arg).Kind() == reflect.Map || reflect.TypeOf(arg).Kind() == reflect.Slice) {
		log.Println("Unable to Dump: Not Map or Struct interface")
		return
	}

	arrayData := reflect.ValueOf(arg)

	w, fer := os.Create(fname)
	if fer != nil {
		log.Print("Error Creating CSV file ", fer)
	}
	cwr := csv.NewWriter(w)
	// var record [4]string

	cwr.Comma = '\t'

	if reflect.TypeOf(arg).Kind() == reflect.Map {

		mapkeys := arrayData.MapKeys()
		once := true
		for _, key := range mapkeys {
			metric := arrayData.MapIndex(key).Interface()

			if once {
				headers, _ := Struct2Header(metric)
				w.WriteString("% " + strings.Join(headers, "\t") + "\n")
				once = false
			}
			data, _ := Struct2Strings(metric)
			cwr.Write(data)
		}
	}

	if reflect.TypeOf(arg).Kind() == reflect.Slice {
		tp := reflect.TypeOf(arg).Elem()

		var headers string
		for i := 0; i < tp.NumField(); i++ {

			headers = headers + "\t" + tp.FieldByIndex([]int{i}).Name
		}
		w.WriteString("%" + headers + "\n")

		for i := 0; i < arrayData.Len(); i++ {

			metric := arrayData.Index(i).Interface()
			data, _ := Struct2Strings(metric)
			fmt.Println(data, metric)
			cwr.Write(data)
		}
	}
	cwr.Flush()
	w.Close()
}

func Log(vec VectorF) (result VectorF) {
	result.Resize(vec.Len())
	for i, val := range vec {
		if val != 0 {
			result[i] = math.Log(val)
		}
	}
	return result
}
