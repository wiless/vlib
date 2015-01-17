package vlib_test

import (
	"fmt"
	"github.com/wiless/vlib"
	"testing"
)

func init() {
	fmt.Println("Testing initialized")
}

func TestNewEyeF(t *testing.T) {
	output := vlib.NewEyeF(5)
	exOutput := vlib.MatrixF{
		{1, 0, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 0},
		{0, 0, 0, 0, 1},
	}

	// fmt.Print(output.Elems())

	if !output.IsEq(exOutput) {
		t.Error("NewEyeF : Did not return the correct size ")
	}

}

func TestReShape(t *testing.T) {
	input := vlib.VectorF{0, 1, 2, 3, 4, 5}
	output := vlib.ReShape(input, 2, 3)
	exOutput := vlib.MatrixF{
		{0, 2, 4},
		{1, 3, 5},
	}
	// fmt.Print("x", output)
	if !output.IsEq(exOutput) {
		t.Error("ReShape : 2x3 Reshape did not match")
	}

}

func TestElems(t *testing.T) {
	exOutput := vlib.VectorF{0, 1, 2, 3, 4, 5}
	input := vlib.MatrixF{
		{0, 2, 4},
		{1, 3, 5},
	}
	output := input.Elems()

	// fmt.Printf("\nElems(%v) should return %v,\n but returned %v", input, exOutput, output)
	if !output.IsEq(exOutput) {
		t.Errorf("\nElems(%v) should return %v,\n but returned %v", input, exOutput, output)
	}

}

func TestT(t *testing.T) {
	fname := "T"
	input := vlib.MatrixF{
		{0, 2, 4},
		{1, 3, 5},
	}
	exOutput := vlib.MatrixF{
		{0, 1},
		{2, 3},
		{4, 5},
	}

	output := input.T()

	// fmt.Printf("\nElems(%v) should return %v,\n but returned %v", input, exOutput, output)
	if !output.IsEq(exOutput) {
		t.Errorf("\n%s(%v) should return %v,\n but returned %v", fname, input, exOutput, output)
	}

}

func TestMatchDim(t *testing.T) {
	fname := "MatchDim"
	input := vlib.MatrixF{
		{0, 2, 4},
		{1, 3, 5},
	}
	input1 := vlib.MatrixF{
		{0, 2, 4},
		{1, 3, 5},
	}
	exOutput := true
	output := vlib.MatchDim(&input, &input1)
	if output != exOutput {
		t.Errorf("\n%s(%v,%v) should return %v, but returned %v", fname, input, input1, exOutput, output)
	}
}

func TestMatchResize(t *testing.T) {
	fname := "Resize"

	input := vlib.MatrixF{
		{0, 2, 4},
		{1, 3, 5},
	}
	exOutput := vlib.MatrixF{
		{0, 2, 4, 0},
		{1, 3, 5, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	output := vlib.Resize(&input, 4, 4)

	if !vlib.MatchDim(&exOutput, output) {
		t.Errorf("\n%s(%v) should return %v, but returned %v", fname, input, exOutput, output)
	}
}
