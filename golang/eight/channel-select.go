
/**
select 多路监控 channel状态
 */

package main

import "fmt"

func fibonacii(c,quit chan int)  {
	x,y := 1,1
	for  {
		select {
		case c <-x:
			// 如果c可写，则该case就会进来
			x = y
			y = x + y
		case <-quit: // 可读就会退出
			fmt.Println("quit")
			return
		//
		}
	}
}
func main()  {

	// select 一般都会和for搭配 可以多路监控channel状态
	c := make(chan int) // 定义管道
	quit := make(chan int) // 定义退出管道

	 // sub go
	go func() {
		for i :=0; i<6; i++{
			// 只要可读 方法里case 就可以写
			fmt.Println(<-c) // 去c里读数据有数据就打印出来，如果没有数据就会阻塞
		}

		quit<- 0  // 循环结束后 发送0 表示退出
	}()

	// main go  监控两个channel状态 谁先触发就先处理谁的业务
	fibonacii(c,quit)
}