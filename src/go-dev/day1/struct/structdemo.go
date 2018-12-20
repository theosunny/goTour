package structdemo

import "fmt"

type Circle struct {
	x      int
	y      int
	Radius int
}

//结构体变量最常见的创建形式
func StructDemo() {
	var c Circle = Circle{
		x:      100,
		y:      100,
		Radius: 50,
	}
	fmt.Printf("%+v\n", c)
}
func StructDemo1() {
	var c Circle = Circle{
		Radius: 50,
	}
	var c2 Circle = Circle{}
	var c3 = Circle{1, 2, 3}
	fmt.Printf("%+v\n", c)
	fmt.Printf("%+v\n", c2)
	fmt.Printf("%+v\n", c3)
}
