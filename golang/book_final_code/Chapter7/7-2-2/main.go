package main

import "fmt"

func main()  {
	menu := Menu("红烧肉",88)
	fmt.Println(menu)
}

func Menu(name string, price int) map[string]int {
	// 执行逻辑代码 返回结果
	m:=make(map[string]int)
	m[name]=price
	return m
}
