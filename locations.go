package vlib

import (
	"math"
)

// Array of Location3D, supports plotter XYer, XYZer
type VectorPos3D []Location3D

type Location3D struct {
	X, Y, Z float64
}

func (v VectorPos3D) Len() int {
	return len(v)
}

func (v VectorPos3D) XY(i int) (x, y float64) {
	if i < len(v) {
		return v[i].X, v[i].Y
	} else {
		// return math.MaxFloat64, math.MaxFloat64
		return math.NaN(), math.NaN()
	}
}

func (v VectorPos3D) X() VectorF {
	r := NewVectorF(v.Len())
	for i := 0; i < v.Len(); i++ {
		r[i] = v[i].X
	}
	return r
}

func (v VectorPos3D) Y() VectorF {
	r := NewVectorF(v.Len())
	for i := 0; i < v.Len(); i++ {
		r[i] = v[i].Y
	}
	return r
}
func (v VectorPos3D) Z() VectorF {
	r := NewVectorF(v.Len())
	for i := 0; i < v.Len(); i++ {
		r[i] = v[i].Z
	}
	return r
}
func (l *Location3D) ToSpherical() (r, thetaH, thetaV float64) {

	return 0, 0, 0
}

// FromSpherical converts the r,thetaH and thetaV (all in degree to Cartesian)
func (l *Location3D) FromSpherical(r, thetaH, thetaV float64) {
	thetaH = ToRadian(thetaH)
	thetaV = ToRadian(thetaV)
	l.X = r * math.Sin(thetaH) * math.Cos(thetaV)
	l.Y = r * math.Sin(thetaH) * math.Sin(thetaV)
	l.Z = r * math.Cos(thetaH)
}

func (l *Location3D) Float64() []float64 {
	return []float64{l.X, l.Y, l.Z}
}

func (l *Location3D) Float32() []float32 {
	return []float32{float32(l.X), float32(l.Y), float32(l.Z)}
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

func (l Location3D) Cmplx() complex128 {
	return complex(l.X, l.Y)
}

func (l Location3D) Scale3D(factor float64) Location3D {
	l.X *= factor
	l.Y *= factor
	l.Z *= factor
	return l
}

func (l Location3D) Scale(factor float64) Location3D {
	l.X *= factor
	l.Y *= factor
	// l.Z = factor
	return l
}

func (l Location3D) Shift3D(delta Location3D) Location3D {
	l.X += delta.X
	l.Y += delta.Y
	l.Z += delta.Z
	return l
}

func (l *Location3D) Shift2D(deltaxy complex128) {
	l.Shift3D(FromCmplx(deltaxy))
}

func (l *Location3D) SetLoc(loc2D complex128, height float64) {
	*l = FromCmplx(loc2D)
	l.SetHeight(height)
}
func (l *Location3D) SetXY(x, y float64) {
	l.X, l.Y = x, y
}

func (l *Location3D) SetXYZ(x, y, z float64) {
	l.X, l.Y, l.Z = x, y, z
}

func FromVectorC(loc2d VectorC, height float64) VectorPos3D {
	result := make([]Location3D, loc2d.Size())

	for indx, val := range loc2d {
		result[indx] = FromCmplx(val)
		result[indx].SetHeight(height)
	}
	return result
}
