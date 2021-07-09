package main

import "fmt"

/**
返回单值：要结果
 */

func main()  {
	reuslt:=Desc("川菜","佛跳墙")
	fmt.Println(reuslt)
}

func Desc(foodType string,foodName string) string  {
	r:= "这是"+foodType+"的"+foodName
	// 很多业务逻辑
	return r
}