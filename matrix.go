package vlib

import (
	"fmt"
	"log"

	// "math"
	// // "os"
	// "strconv"
	// "strings"
)

type Matrix []Vector
type MatrixF []VectorF
type MatrixC []VectorC

type Matrixer interface {
	NRows() int
	NCols() int
	Size() (int, int)
	AppendNRows(int)
	AppendNCols(int)
	// GetRow(int) VectorIface
	// SetRow(int, VectorIface)
	// SetCol(int, VectorIface)

	// GetRow(int) VectorIface
}

// type Dimension struct {
// 	Rows int
// 	Cols int
// }

// type GMatrixF struct {
// 	MatrixF
// 	Dimension
// }

func ReShape(v VectorF, rows, cols int) MatrixF {
	m := NewMatrixF(rows, cols)
	cnt := 0
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			m[j][i] = v[cnt]
			cnt++
		}
	}
	return m
}

func (m MatrixF) Elems() VectorF {

	rows, cols := m.Size()
	result := NewVectorF(rows * cols)
	cnt := 0
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			result[cnt] = m[j][i]
			cnt++
		}
	}
	return result
}

func NewOnesMatF(rows, cols int) MatrixF {

	result := MatrixF(make([]VectorF, rows))
	for i := 0; i < rows; i++ {
		result[i] = NewOnesF(cols)
	}
	return result
}

func (m MatrixF) IsEq(val MatrixF) bool {
	if m.NRows() != val.NRows() || m.NCols() != val.NCols() {
		return false
	}
	for i := 0; i < m.NRows(); i++ {
		if !m[i].IsEq(val[i]) {

			return false
		}
	}
	return true
}

func (m MatrixC) IsEq(val MatrixC) bool {
	if m.NRows() != val.NRows() || m.NCols() != val.NCols() {
		return false
	}
	for i := 0; i < m.NRows(); i++ {
		if !m[i].IsEq(val[i]) {
			log.Println("\nrow ", i, m[i], "not match", val[i])
			return false
		}
	}
	return true
}

func (m MatrixF) NRows() (rows int) {

	if len(m) == 0 {
		return 0
	}
	return len(m)
}

func (m MatrixF) NCols() (cols int) {
	if len(m) == 0 {
		return 0
	}
	return len(m[0])
}
func (m MatrixF) Size() (rows, cols int) {
	return m.NRows(), m.NCols()
}

func NewEyeF(rows int) MatrixF {

	dvec := NewOnesF(rows)
	return NewDiagMatF(dvec)
}

func NewMatrixF(rows, cols int) MatrixF {

	result := MatrixF(make([]VectorF, rows))
	for i := 0; i < rows; i++ {
		result[i] = NewVectorF(cols)
	}
	return result
}
func NewMatrixC(rows, cols int) MatrixC {

	result := MatrixC(make([]VectorC, rows))
	for i := 0; i < rows; i++ {
		result[i] = NewVectorC(cols)
	}
	return result
}

func CreateMatrixF(matrix MatrixF) MatrixF {
	rows := matrix.NRows()
	result := NewMatrixF(matrix.NRows(), matrix.NCols())
	for i := 0; i < rows; i++ {
		copy(result[i], matrix[i])
	}
	return result
}

func (m MatrixF) String() string {
	var str string

	rows := len(m)
	// str = "=["
	str = "["
	for i := 0; i < rows; i++ {
		str += fmt.Sprintf("%f", m[i])
		if i != rows-1 {
			str += ",\n"
		}
	}
	str += "]"
	return str
}

func (m MatrixC) String() string {
	var str string

	rows := len(m)
	// str = "=["
	str = "["
	for i := 0; i < rows; i++ {
		str += fmt.Sprintf("%f", m[i])
		if i != rows-1 {
			str += ",\n"
		}
	}
	str += "]"
	return str
}

func (m MatrixF) Get(row, col int) float64 {
	var elemval float64
	elemval = m[row][col]
	return elemval
}

func (m MatrixF) GetRow(row int) VectorF {
	// var elemrow []float64
	return m[row]
	// return elemrow
}

func (m MatrixC) GetRow(row int) VectorC {
	// var elemrow []float64
	return m[row]
	// return elemrow
}

func (m MatrixF) GetCol(col int) VectorF {
	// var elemrow []float64
	cols := len(m[0])
	rows := len(m)
	resultvector := make([]float64, rows)

	if col < cols {
		for i := 0; i < rows; i++ {
			resultvector[i] = m[i][col]
		}
	} else {
		log.Fatalf("MatrixF Index out of bound %d of %d", col, cols)
	}
	return resultvector
	// return elemrow
}

func (m MatrixF) GetColRange(begin, end int) MatrixF {
	// col := 9
	// x:=vlib.
	str := fmt.Sprintf("%d:%d", begin, end)
	indx := ToVectorI(str)
	result := m.GetCols(indx...)
	return result

}

func (m MatrixF) GetCols(col ...int) MatrixF {
	// var elemrow []float64
	rows := len(m)
	cols := len(col)

	result := NewMatrixF(rows, cols)
	if col != nil {

		for j := 0; j < cols; j++ {
			column := col[j]

			for row := 0; row < rows; row++ {
				result[row][j] = m[row][column]
			}
		}

	} else {
		log.Fatalf("MatrixF Index out of bound %d of %d", col, cols)
	}
	return result

}

func (m MatrixF) GetSubMatF(srow, scol, rows, cols int) MatrixF {
	result := NewMatrixF(rows, cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[i][j] = m[i+srow][j+scol]
		}
	}

	return result
}

func (m MatrixF) GetDiagF() VectorF {
	var result VectorF
	for i := 0; i < len(m); i++ {
		result[i] = m.Get(i, i)
	}
	return result

}

func NewDiagMatF(input VectorF) MatrixF {
	rows := len(input)
	result := NewMatrixF(rows, rows)
	for i := 0; i < rows; i++ {
		result[i][i] = input[i]
	}
	return result

}

func (m *MatrixF) SetSubMatF(srow, scol int, input MatrixF) {

	rows := len(input)
	cols := len(input[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			(*m)[i+srow][j+scol] = input[i][j]
		}
	}

}

func (m *MatrixC) T() MatrixC {
	rows, cols := m.Size()
	result := NewMatrixC(cols, rows)

	for i := 0; i < rows; i++ {
		colVector := m.GetRow(i)
		result.SetCol(i, colVector)
	}
	return result
}

// func T(m Matrixer) Matrixer {
// 	rows, cols := m.Size()
// 	var result Matrixer
// 	Resize(result, cols, rows)

// 	for i := 0; i < rows; i++ {
// 		colVector := m.GetRow(i)
// 		result.SetCol(i, colVector)
// 	}
// 	return result
// }

/// Transposes the matrix
func (m MatrixF) T() MatrixF {
	rows, cols := m.Size()
	result := NewMatrixF(cols, rows)
	var colVector VectorF
	for i := 0; i < rows; i++ {
		colVector = m.GetRow(i)
		result.SetCol(i, colVector)
	}
	return result
}

func (m *MatrixF) SetRow(srow int, rowVector VectorF) {

	// rows := len(input)
	// cols := len(rowVector)
	(*m)[srow] = rowVector
	// for i := 0; i < rows; i++ {
	// for j := 0; j < cols; j++ {
	// (*m)[srow][j] = rowVector[j]
	// }
	// }

}

func (m *MatrixF) SetCol(scol int, colVector VectorF) {

	rows := len(colVector)

	for j := 0; j < rows; j++ {
		(*m)[j][scol] = colVector[j]
	}
}

func (m *MatrixC) SetCol(scol int, colVector VectorC) {

	rows := len(colVector)

	for j := 0; j < rows; j++ {
		(*m)[j][scol] = colVector[j]
	}
}

func (m MatrixF) DeleteColumn(col int) MatrixF {
	rows, cols := m.Size()
	result := NewMatrixF(rows, cols)
	if col >= cols {
		fmt.Printf("%% DeleteColumn: Index Out of Bound")
		return make([]VectorF, 0)
	}
	for i := 0; i < rows; i++ {
		result[i] = m.GetRow(i).Delete(col)
	}
	return result
}

func (m MatrixF) InsertColumnVector(pos int, input VectorF) MatrixF {
	rows, cols := m.Size()

	if input.Size() != rows {
		log.Panicf("\nInsertColumn Rows %d <>Input %d", rows, input.Size())
	}
	cols++
	result := NewMatrixF(rows, cols)
	for j := 0; j < rows; j++ {
		result[j] = m[j].Insert(pos, input[j])
	}
	return result
}

func (m MatrixF) InsertColumn(pos int, val float64) MatrixF {
	result := m.Insert(pos)
	rows := result.NRows()
	for i := 0; i < rows; i++ {
		result[i][pos] = val
	}
	return result
}

func (m MatrixF) InsertOnes(pos int) MatrixF {

	result := m.InsertColumn(pos, 1)
	return result
}

func (m MatrixF) Insert(pos int) MatrixF {
	rows, cols := m.Size()
	cols++
	result := NewMatrixF(rows, cols)
	for j := 0; j < rows; j++ {
		result[j] = m[j].Insert(pos, 0)
	}
	return result
}

func (m *MatrixF) AppendNRows(n int) {
	v := NewVectorF(m.NCols())
	for i := 0; i < n; i++ {
		*m = append(*m, v)
	}
}

func (m *MatrixF) AppendNCols(n int) {
	v := NewVectorF(m.NRows())
	for i := 0; i < n; i++ {
		m.AppendColumn(v)
	}
}

func (m *MatrixC) AppendNRows(n int) {
	v := NewVectorC(m.NCols())
	for i := 0; i < n; i++ {
		*m = append(*m, v)
	}
}

func (m *MatrixC) AppendNCols(n int) {
	v := NewVectorC(m.NRows())
	for i := 0; i < n; i++ {
		m.AppendColumn(v)
	}
}

func (m *MatrixF) AppendRow(v VectorF) {
	/// Fill the columns of matrix with data from  source column
	if m.NCols() != len(v) {
		log.Panicln("AppendRow() : %s mismatch cols(Matrix)=%d", v, m.NCols())
	}
	*m = append(*m, v)
}

func (m *MatrixF) AppendColumn(colvec VectorF) {
	targetRows := colvec.Size()
	rows, cols := m.Size()

	/// Fill the columns of matrix with data from  source column
	minRows := rows
	if targetRows < minRows {
		minRows = targetRows
	}

	for i := 0; i < minRows; i++ {
		(*m)[i].Resize(cols + 1)
		(*m)[i][cols] = colvec[i]
	}

	extraRows := targetRows - rows
	for i := 0; i < extraRows; i++ {

		*m = append(*m, NewVectorF(cols+1))

		(*m)[minRows+i][cols] = colvec[minRows+i]
		// fmt.Printf("\n Adding %vth Row Col valu %v ", minRows+i, colvec[minRows+i])
	}

	/// if targetRows>rows , add  additional rows zeros(targetRows-rows,cols)

}

func (m *MatrixC) AppendColumn(colvec VectorC) {
	targetRows := colvec.Size()
	rows, cols := m.Size()

	/// Fill the columns of matrix with data from  source column
	minRows := rows
	if targetRows < minRows {
		minRows = targetRows
	}

	for i := 0; i < minRows; i++ {
		(*m)[i].Resize(cols + 1)
		(*m)[i][cols] = colvec[i]
	}

	extraRows := targetRows - rows
	for i := 0; i < extraRows; i++ {

		*m = append(*m, NewVectorC(cols+1))

		(*m)[minRows+i][cols] = colvec[minRows+i]
		// fmt.Printf("\n Adding %vth Row Col valu %v ", minRows+i, colvec[minRows+i])
	}

	/// if targetRows>rows , add  additional rows zeros(targetRows-rows,cols)

}

type genericIface Matrixer

func MatchDim(v, m Matrixer) bool {
	// fmt.Printf("\n Type of matchdim %v \n", reflect.TypeOf(m).String())
	// fmt.Printf("\n (%v,%v) ==  (%v,%v ??", v.NRows(), v.NCols(), m.NRows(), m.NCols())
	//if reflect.TypeOf(v) == reflect.TypeOf(m) {
	if v.NRows() != m.NRows() || v.NCols() != m.NCols() {
		return false
	}

	//}/

	return true
}

func Resizer(v Matrixer, rows, cols int) {

	r, c := (v).NRows(), (v).NCols()

	if cols < c || rows < r {
		log.Panicln("Trucation currently not enabled")
	}

	if rows > r {
		v.AppendNRows(rows - r)

		if cols > c {
			v.AppendNCols(cols - c)
		}

	}

}

/// Resizes the input matrix by either padding zeros or trucating rows/cols as per need
func Resize(v Matrixer, rows, cols int) Matrixer {
	var result Matrixer

	r, c := v.Size()

	if cols < c || rows < r {
		log.Panicln("Trucation currently not enabled")
	}

	if rows > r {
		result = v
		result.AppendNRows(rows - r)

		if cols > c {
			result.AppendNCols(cols - c)
		}
		// if cols < c {
		// 	result.DelCols(c - cols)
		// }
	}

	// if rows <= r {

	// }
	return result
	//result.Resize(v.NRows(), v.NCols())

	return result
}

func ToMatrixC(re MatrixF) MatrixC {

	result := NewMatrixC(re.NRows(), re.NCols())
	for indx, _ := range re {
		result[indx] = ToVectorC(re.GetRow(indx))
	}
	return result
}

func ToMatrixC2(re, im MatrixF) MatrixC {

	result := NewMatrixC(re.NRows(), re.NCols())
	for indx, _ := range re {
		result[indx] = ToVectorC2(re.GetRow(indx), im.GetRow(indx))
	}
	return result
}

func (m MatrixC) NRows() (rows int) {

	if len(m) == 0 {
		return 0
	}
	return len(m)
}

func (m MatrixC) NCols() (cols int) {
	if len(m) == 0 {
		return 0
	}
	return len(m[0])
}
func (m MatrixC) Size() (rows, cols int) {
	return m.NRows(), m.NCols()
}

func (m *MatrixC) Scale(val float64) {
	newval := complex(val, 0)
	m.ScaleC(newval)
}

func (m *MatrixC) ScaleC(val complex128) {
	for i := 0; i < m.NRows(); i++ {
		for j := 0; j < m.NCols(); j++ {
			(*m)[i][j] *= val
		}
	}
}

// output.Minus(exOutput)
func (m MatrixC) Minus(val MatrixC) MatrixC {

	if !MatchDim(&m, &val) {
		return MatrixC{}
	}
	result := NewMatrixC(m.NRows(), m.NCols())
	for i := 0; i < m.NRows(); i++ {
		for j := 0; j < m.NCols(); j++ {
			result[i][j] = m[i][j] - val[i][j]
		}
	}
	return result
}

// output.Minus(exOutput)
func (m MatrixF) Minus(val MatrixF) MatrixF {

	if !MatchDim(&m, &val) {
		return MatrixF{}
	}
	result := NewMatrixF(m.NRows(), m.NCols())
	for i := 0; i < m.NRows(); i++ {
		for j := 0; j < m.NCols(); j++ {
			result[i][j] = m[i][j] - val[i][j]
		}
	}
	return result
}
