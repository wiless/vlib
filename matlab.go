package vlib

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strings"
)

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
	// if strings.HasSuffix(fname, ".m") {
	fname = strings.TrimSuffix(fname, ".m")
	// }

	m.filename = fname + ".m"
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

	// if strings.HasSuffix(fname, ".m") {
	fname = strings.TrimSuffix(fname, ".m")
	// }

	result := Matlab{}
	result.SetFile(fname)
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
	case "vlib.VectorI":
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
		if len(m.Keys) > 0 {
			log.Println("Appending Keys ", m.Keys)
			m.ExportStruct("Keys", m.Keys)
		}
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
