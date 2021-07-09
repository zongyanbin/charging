package main

import (
	"fmt"
	"time"
)

func Work()  {
	fmt.Println("我在工作")
}
func main() {
	fmt.Println("主协程开始运行")
	go Work()
	time.Sleep(1 * time.Second)  // 解决用channel
	fmt.Println("主协程运行结束")
}
