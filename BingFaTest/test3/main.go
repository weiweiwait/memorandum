package main

import (
	"fmt"
	"time"
)

var (
	testMap = make(map[int]int, 10)
)

func testNum(num int) {
	res := 1
	for i := 1; i <= num; i++ {
		res *= 1
	}
	testMap[num] = res
}
func main() {
	start := time.Now()
	for i := 1; i < 200; i++ {
		go testNum(i)
	}
	time.Sleep(time.Second * 5)
	for key, val := range testMap {
		fmt.Println("数字%v 对应的接触是 %v\n", key, val)
	}
	end := time.Since(start)
	fmt.Println(end)
}
