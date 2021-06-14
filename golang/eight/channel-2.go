/**
close 关闭chan
确定没有数据写入了才可以关闭
 */

package main

import "fmt"

func main(){
	c := make(chan  int)
	go func() {
		for i :=0; i<5; i++ {
			c<-i
		}
		close(c)
	}()
/*
这段代码可以简写用 range 来迭代不断操作
	for  {
		// ok如果true表示channel没有关闭，如果为false表示channel已经关闭
		if data,ok :=<-c; ok{
			fmt.Println(data)
		}else{
			break
		}
	}
*/

// 可以使用range 来迭代操作channel
// 简化代码 channel 和 range
for data := range c{
		fmt.Println(data)
}
fmt.Println("MainFinished..")
}

// channel 与 select