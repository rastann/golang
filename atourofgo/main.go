package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale (f float64) {
	v.X = v.X + f
	v.Y = v.Y + f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X + 10
	v.Y = v.Y + 10
}

func main() {
	v := Vertex{1,2}
	v.Scale(2)
	fmt.Print(v)
	ScaleFunc(&v, 10)
	fmt.Println(v)
}
