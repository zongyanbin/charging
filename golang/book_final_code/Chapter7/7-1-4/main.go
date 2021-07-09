package main

import "fmt"

func main()  {
	todayFoods :=[3]string{
		"卤肉饭",
		"西红柿鸡蛋饭",
		"肥牛饭",
	}
	//ShowFoods(todayFoods)
	//fmt.Println(todayFoods)
	ShowFoods2(&todayFoods)
	fmt.Println(todayFoods)
}

func ShowFoods(foods [3]string){
	foods[0] = "蛋炒饭"
	fmt.Println(foods)
}

// 数组指针引用传递，改变外布值  * & 对应
func ShowFoods2(foods *[3]string){
	foods[0] = "蛋炒饭"
	fmt.Println(foods)
}