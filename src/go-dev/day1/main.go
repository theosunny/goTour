package main

import (
	"fmt"
	"go-dev/day1/string"
	"go-dev/day1/struct"
)

func main() {
	fmt.Println("hello")
	stringex.StringDemo1()
	stringex.StringDemo2()
	fmt.Println("-----结构体")
	structdemo.StructDemo()
	structdemo.StructDemo1()
	structdemo.Newstruct()
	structdemo.CopyStruct()
	structdemo.PassStruct()
	structdemo.MethodStruct()
	structdemo.AnonymousStruct()
}
