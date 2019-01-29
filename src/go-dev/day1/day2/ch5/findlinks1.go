package main

import "fmt"

func main() {
	fmt.Println(CaculateGs())
	fmt.Println(Ex(7))
	fmt.Println(512 / 7)
}

/**
计算高斯1+100
*/
func CaculateGs() int {
	var res int
	for i := 1; i <= 100; i++ {
		res += i
	}
	return res
}

//一个人赶着鸭子去每个村庄卖，每经过一个村子卖去所赶鸭子的一半又一只。
//这样他经过了七个村子后还剩两只鸭子，问他出发时共赶多少只鸭子？经过每个村子卖出多少只鸭子？
var number = 2
var x = 7

func Ex(i int) int {
	if i == 0 {
		return number
	} else {
		number = (number + 1) * 2
		x = (number / 2) + 1
		fmt.Printf("经过第%d村卖了%d\n", i, x)
		return Ex(i - 1)
	}

}
