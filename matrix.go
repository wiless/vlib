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

type Dimension struct {
	Rows int
	Cols int
}

type GMatrixF struct {
	MatrixF
	Dimension
}

func NewOnesMatF(rows, cols int) MatrixF {

	result := MatrixF(make([]VectorF, rows))
	for i := 0; i < rows; i++ {
		result[i] = NewOnesF(cols)
	}
	return result
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
	str = "=["
	for i := 0; i < rows; i++ {
		str += fmt.Sprintf("%f", m[i])
		if i != rows-1 {
			str += "\n"
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
