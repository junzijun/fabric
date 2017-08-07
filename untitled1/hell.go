package main

import "math"

importC(
  "fmt"
  "math"
)
type Vertex struct{
	X,Y float64
}
func (v *Vertex) Scale(f float64){
	v.X = v.X *f
	v.Y = v.Y *f
}
func (v *Vertex) Abs() float64{
	return math.Sqrt(v.X*v+v.Y *)
}