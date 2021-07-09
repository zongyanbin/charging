package main

import "fmt"

/**
无缓冲：发送和接收同时准备好
	手递手
*/

func Eat(foodName string, c chan bool) {
	fmt.Println("我正在吃" + foodName)
	c <- true
}

func main() {
	fmt.Println("主协程开始运行")
	c := make(chan bool)
	go Eat("生蚝", c)
	<-c
	fmt.Println("主协程运行结束")
}
