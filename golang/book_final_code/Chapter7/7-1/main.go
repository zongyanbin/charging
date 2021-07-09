package main

import "fmt"

func main()  {
	//func 函数名 (参数列表) (返回列表){
	// 函数体
	//}
	add1:=add(1,5)
	add2:=add2(3,4)
	fmt.Println("add1",add1)
	fmt.Println("add2",add2)

}

func add(x,y int) int {
	return x+y
}

func add2(x,y int) (result int) {
	result= x+y
	return
}
