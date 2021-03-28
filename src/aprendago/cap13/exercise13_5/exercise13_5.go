package main

// Aulas sobre interface e polimorfismo:
// https://www.youtube.com/watch?v=2zTENBJTlD0&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=92
// https://www.youtube.com/watch?v=3pbmLfcgN2s&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=102

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

func (s square) area() {
	result := s.side * s.side
	fmt.Println("Area of square:", result)
}

type circle struct {
	ray float64
}

func (c circle) area() {
	result := math.Pi * 2 * c.ray
	fmt.Println("Area of circle:", result)
}

type info interface {
	area()
}

func measure(i info) {
	i.area()
}

func main() {

	x := square{
		side: 15.0,
	}

	y := circle{
		ray: 5.25,
	}

	//x.area()
	//y.area()

	measure(x)
	measure(y)
}
