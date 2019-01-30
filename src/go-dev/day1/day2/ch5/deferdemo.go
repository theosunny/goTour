package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	defer defertest()()
	time.Sleep(3 * time.Second) // simulate slow
	fmt.Println("我是主方法")
}

func defertest() func() {
	start := time.Now()
	log.Println("调用defer开始", start)
	return func() {
		log.Println("调用defer结束", time.Since(start))
	}
}
