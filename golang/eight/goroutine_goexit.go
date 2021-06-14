/**
goroutine 匿名函数
 退出用 runtime.Goexit()
 通信用 go channel
 */
package main

import (
	"fmt"
	"time"
)

func main(){
	// 用go创建承载一个形参为空，返回值为空的一个函数
/*	go func() {
		defer fmt.Println("A.defer...")
		func(){
			defer fmt.Println("B.defer...")
			// 退出当前 goroutine
			runtime.Goexit()
			fmt.Println("B")
		}()

		fmt.Println("A")
	}()*/

	//有参数 不能通过返回值拿到只能通过channel
	go func(a int, b int) bool {
		fmt.Println("a=",a,"b=",b)
		return true
	}(10,20)

	// 死循环
	for {
		time.Sleep(1 *time.Second)
	}
}
