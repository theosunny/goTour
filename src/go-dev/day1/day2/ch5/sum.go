package main

import (
	"fmt"
	"os"
)

func main() {
	vals := []int{1, 22}
	fmt.Println(f(1, 2, 34, 5))
	fmt.Println(f1(vals...))
	linenum, name := 12, "count"
	errorf(linenum, "undefined: %s", name) // "Line 12: undefined: count"
}

//可变参数
func f(x ...int) int {
	var total int
	i := len(x)
	for j := 0; j < i; j++ {
		total += x[j]
	}

	return total
}
func f1(x ...int) int {
	var total int
	i := len(x)
	for j := 0; j < i; j++ {
		total += x[j]
	}

	return total
}

func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}
