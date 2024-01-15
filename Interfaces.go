package main

import (
	"fmt" 
	"math"
)

//Interfaces are collections of method signatures. A type "implements" an interface if it has all of the methods of the given interface defined on it.
type shape interface {
	area() float64

}

type rect struct {
		width, height float64
}
func (r rect) area() float64 {
		return r.width * r.height
}

type circle struct {
		radius float64
}
func (c circle) area() float64 {
		return math.Pi * c.radius * c.radius
}

func getArea(sh shape) float64 {
	return sh.area()
}



func main() {
	r:= rect{
		width:5.0,
		height:2.0,
	}
	fmt.Printf("Area: %.2f\n", getArea(r))
}	
