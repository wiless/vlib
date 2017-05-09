// Helps interface plotter XYers with vlib library routines
package vlib

type Table []VectorF

//XY to support plotters
func (v VectorF) XY(i int) (float64, float64) {
	return float64(i), v.Get(i)
}
