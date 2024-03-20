package main

import (
	"fmt"
	"runtime"
)

var num int = 1

func main() {
	//for i := 1; i < 10000000; i++ {
	//	go runTimes(1)
	//}
	runtime.GOMAXPROCS(16) //设置cpu最大核心数
	fmt.Println(runtime.NumCPU())
	//for i:=1;i<=10;i++{
	//	fmt.Println("main", i, "你好!", 10-i)
	//	time.Sleep(time.Second*2)
	//}
}
func runTimes(times int) int {
	for i := 1; i <= times; i++ {
		fmt.Println("runtimes", i, "你好!", times-i)
		fmt.Println("num:", num)
	}
	num++
	return times
}
