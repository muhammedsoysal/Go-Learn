package interfaces

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type rectangel struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rectangel) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func school(s shape) {
	fmt.Println("Şekilin Alanı: ", s.area())
}

func Demo1() {
	r := rectangel{10, 6}

	school(r)

	c := circle{10}
	school(c)

}
