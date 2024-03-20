package main

import "fmt"

var intChan chan int

func main() {
	intChan = make(chan int, 10)
	intChan <- 1
	fmt.Println(intChan)
	fmt.Println("intChan的值是:%v,地址是%v\n", intChan, &intChan)
	fmt.Println("intChan的大小是:%v,容量是%v\n", len(intChan), cap(intChan))
}
