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

func (m MatrixF) Size() (rows, cols int) {
	return len(m), len(m[0])
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

func (m MatrixF) String() string {
	var str string

	rows := len(m)
	str = "["
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

func (m MatrixF) InsertColumn(pos int, input VectorF) MatrixF {
	rows, cols := m.Size()

	if input.Size() != rows {
		log.Panicf("InsertColumn Rows %d <>Input %d", rows, input.Size())
	}
	cols++
	result := NewMatrixF(rows, cols)
	for j := 0; j < rows; j++ {
		result[j] = m[j].Insert(pos, input[j])
	}
	return result
}
