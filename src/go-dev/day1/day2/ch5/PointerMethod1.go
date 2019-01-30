package main

import "fmt"

type Point1 struct {
	X float64
	Y float64
}

func (p *Point1) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	point := &Point1{1, 2}
	point1 := Point1{1, 2}
	point1.ScaleBy(1)
	fmt.Println(point1)
	point.ScaleBy(12)
	fmt.Println(*point)
}
