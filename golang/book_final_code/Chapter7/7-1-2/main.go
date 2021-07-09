package main

import "fmt"
// 实际参数和形式参数
func main()  {
	Fav("老王")
	Fav("老李")
}

func Fav(username string)  {
	fmt.Println(username+"喜欢吃生蚝")
}
