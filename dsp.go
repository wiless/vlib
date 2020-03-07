package vlib

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func printM(str string, m mat.Matrix) {
	// fmt.Printf("\n%s=\n %f\n", str, m)
	fmt.Printf("\n%s=%m\n", str, mat.Formatted(m, mat.Prefix(" "), mat.Squeeze()))
}

func Conv(input, filt []float64, trunc bool) []float64 {
	M := len(input)
	L := len(filt)
	N := L + M - 1
	if trunc {
		N = M
	}
	result := NewVectorF(N)

	for i := 0; i < N; i++ {
		result[i] = 0 // set to zero before sum

		for j := 0; j < L; j++ {
			tau := i - j
			if tau < M && tau >= 0 {
				result[i] += input[i-j] * filt[j] // convolve: multiply and accumulate
			}

			// }
		}
	}
	return result
}

//Conv2 Implements 2D convolution Vertical then Horizontal
func Conv2(X mat.Matrix, filt []float64, trunc bool) *mat.Dense {
	R, C := X.Dims()
	L := len(filt)
	NR := L + R - 1
	NC := L + C - 1
	if trunc {
		NR = R
		NC = C
	}
	Y := mat.NewDense(NR, NC, nil)
	ROWS, COLS := X.Dims()
	for c := 0; c < COLS; c++ {

		col := mat.Col(nil, c, X)
		output := Conv(col, filt, trunc)
		Y.SetCol(c, output)
	}
	ROWS, COLS = Y.Dims()
	for r := 0; r < ROWS; r++ {

		row := mat.Row(nil, r, Y)
		output := Conv(row[0:R], filt, trunc)
		Y.SetRow(r, output)
	}
	return Y
}
