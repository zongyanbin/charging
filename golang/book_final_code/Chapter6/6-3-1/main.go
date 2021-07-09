package main

import "fmt"

func main() {
	var m map[string]int = map[string]int{"红烧肉": 88, "清蒸鱼": 98, "溜大虾": 128, "蒸螃蟹": 198, "鲍鱼粥": 68}
	price,ok:=m["西红柿鸡蛋"]
	if !ok{
		fmt.Println("没有找到对应的键值对")
		// 返回
	}
	fmt.Println(price)
}
