package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	s := strings.Map(func(r rune) rune { return r + 1 }, "HAL-9000")
	log.SetPrefix("date:")
	log.Println(s)
	f := squares()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"
}
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
