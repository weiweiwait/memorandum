package main

import (
	"fmt"
	"time"
)

func main() {
	go runTimes(10)
	for i := 1; i <= 10; i++ {
		fmt.Println("main", i, "你好", 10-i)
		time.Sleep(time.Second * 2)
	}
}
func runTimes(times int) int {
	for i := 1; i <= times; i++ {
		fmt.Println("runTimes", i, "你好!", times-i)
		time.Sleep(time.Second)
	}
	return times
}
