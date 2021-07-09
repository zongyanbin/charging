package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum2 := 0
	for {
		sum2++
		if sum2 > 100 {
			fmt.Println("break")
			break
		}
	}


	sum3:=0
Out:  // out 跳转到某种场景
	for  {
		sum3++
		if sum3>10{
			fmt.Println("我要退出")
			break Out
		}
	}
	fmt.Println("此处结束")
}
