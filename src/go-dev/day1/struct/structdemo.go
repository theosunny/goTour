package structdemo

import (
	"fmt"
	"math"
	"unsafe"
)

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
	//顺序形式
	var c3 = Circle{1, 2, 3}
	//指针的形式
	var pointer *Circle = &Circle{100, 100, 50}
	fmt.Printf("%+v\n", c)
	fmt.Printf("%+v\n", c2)
	fmt.Printf("%+v\n", c3)
	fmt.Printf("%+v\n", pointer)
}

//使用new创建零值的结构体 new() 函数返回的是指针类型。
func Newstruct() {
	var c *Circle = new(Circle)
	c.Radius = 1000
	fmt.Printf("%+v\n", c)
	//这种不太雅观
	var c1 Circle
	fmt.Printf("%+v\n", c1)
	//nil 结构体是指结构体指针变量没有指向一个实际存在的内存。
	// 这样的指针变量只会占用 1 个指针的存储空间，也就是一个机器字的内存大小。
	//而零值结构体是会实实在在占用内存空间的，只不过每个字段都是零值。
	// 如果结构体里面字段非常多，那么这个内存空间占用肯定也会很大。
	var c3 *Circle = nil
	fmt.Printf("%+v\n", c3)
}

//结构体的大小
func StructSize() {
	var c = Circle{Radius: 50}
	fmt.Println(unsafe.Sizeof(c))
}

//结构体之间可以相互赋值，它在本质上是一次浅拷贝操作，拷贝了结构体内部的所有字段。
// 结构体指针之间也可以相互赋值，它在本质上也是一次浅拷贝操作，不过它拷贝的仅仅是指针地址值，结构体的内容是共享的
func CopyStruct() {
	fmt.Println("-----结构体拷贝")
	var c1 = Circle{Radius: 50}
	var c2 = c1
	fmt.Printf("%+v\n", c1)
	fmt.Printf("%+v\n", c2)
	c1.Radius = 100
	fmt.Printf("%+v\n", c1)
	fmt.Printf("%+v\n", c2)

	var c3 *Circle = &Circle{Radius: 50}
	var c4 *Circle = c3
	fmt.Printf("%+v\n", c3)
	fmt.Printf("%+v\n", c4)
	c3.Radius = 100
	fmt.Printf("%+v\n", c3)
	fmt.Printf("%+v\n", c4)
}

//在函数里面修改结构体的状态不会影响到原有结构体的状态，
// 函数内部的逻辑并没有产生任何效果。通过指针传递就不一样。
func PassStruct() {
	fmt.Println("结构体传递")

	var c = Circle{0, 0, 100}
	expandByPointer(&c)
	fmt.Println(c)
	expandByValue(c)
	fmt.Println(c)
}

func expandByValue(c Circle) {
	c.Radius *= 2
}

func expandByPointer(c *Circle) {
	c.Radius *= 2
}

func (c Circle) Area() float64 {
	return math.Pi * float64(c.Radius) * float64(c.Radius)
}
func (c Circle) Circumference() float64 {
	return 2 * math.Pi * float64(c.Radius)
}

//结构体方法
func MethodStruct() {
	fmt.Println("结构体方法")
	var c = Circle{Radius: 50}
	fmt.Println(c.Area(), c.Circumference())
	// 指针变量调用方法形式上是一样的
	var pc = &c
	fmt.Println(pc.Area(), pc.Circumference())

}

type Point struct {
	x int
	y int
}

func (p Point) show() {
	fmt.Println(p.x, p.y)
}

type InnerCircle struct {
	loc    Point
	radius int
}

//内嵌结构体
func InnerStruct() {
	//内部结构体
	var c = InnerCircle{
		loc:    Point{100, 100},
		radius: 50,
	}
	fmt.Printf("%+v\n", c)
	fmt.Printf("%+v\n", c.loc)
	fmt.Printf("%d %d\n", c.loc.x, c.loc.y)
	c.loc.show()
}

//匿名内嵌结构体
type AnonymousCircle struct {
	Point
	radius int
}

//匿名内嵌结构体
func AnonymousStruct() {
	fmt.Println("匿名内嵌结构体")
	var c = AnonymousCircle{
		Point{100, 100},
		50,
	}
	fmt.Printf("%+v\n", c)
	fmt.Printf("%+v\n", c.Point)
	fmt.Printf("%d %d\n", c.x, c.y) // 继承了字段
	fmt.Printf("%d %d\n", c.Point.x, c.Point.y)
	c.show() // 继承了方法
	c.Point.show()
}
