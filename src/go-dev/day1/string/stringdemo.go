package stringex

import "fmt"

func StringDemo1() {
	var s = "嘻哈china"
	for i := 0; i < len(s); i++ {
		fmt.Printf("% x", s[i])
	}
}

func StringDemo2() {
	var s = "嘻哈china"
	for codepoint, runeValue := range s {
		fmt.Printf("%d %d", codepoint, int32(runeValue))
	}
}
func StringDemo3() {
	var s1 = "helo world"
	var b = []byte(s1)
	var s2 = string(b)
	fmt.Println()
	fmt.Println(b)
	fmt.Println(s2)
}
