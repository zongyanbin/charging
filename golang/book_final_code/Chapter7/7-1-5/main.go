package main

import "fmt"

// 切片自带引用传递
func main()  {
	todayFoods :=[]string{
		"卤肉饭",
		"西红柿鸡蛋饭",
		"肥牛饭",
	}
	ShowFoods(todayFoods)
	fmt.Println(todayFoods)
}

func ShowFoods(foods []string){
	foods[0] = "蛋炒饭"
	fmt.Println(foods)
}

