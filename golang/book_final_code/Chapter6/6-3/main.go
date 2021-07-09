package main

import "fmt"

func main() {
	var m map[string]int = map[string]int{"红烧肉": 88, "清蒸鱼": 98, "溜大虾": 128, "蒸螃蟹": 198, "鲍鱼粥": 68}
	fmt.Println(m)

	foodsMap := make(map[string]int)
	foodsMap["红烧肉"] = 88
	foodsMap["清蒸鱼"] = 98
	foodsMap["溜大虾"] = 128
	foodsMap["蒸螃蟹"] = 198
	foodsMap["鲍鱼粥"] = 68
	fmt.Println(foodsMap)

	// 字典可以有一个元素也可以有多个元素
	/**
	字典键有规定但是值可以放很多类型
	字典关联的值： 数组 字符串 切片 字典

	不用的代码就是删掉节省空间
	 */

	var m2 map[string]int = map[string]int{
		"红烧肉": 88,
		"清蒸鱼": 98,
		"溜大虾": 128,
		"蒸螃蟹": 198,
		"鲍鱼粥": 68,
	}
	fmt.Println(m2)
}
