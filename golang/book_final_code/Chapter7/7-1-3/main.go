package main

import "fmt"

func main()  {
	Desc("卤菜", "葱烧海参")
	Desc("川菜","麻婆豆腐")
}

// 方法类型必须和参数类型一致
func Desc(foodType string,foodName string)  {
	fmt.Println("这是" + foodType + "的" + foodName)
}