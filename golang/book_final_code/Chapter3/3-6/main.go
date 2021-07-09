package main

import "fmt"
/**
数组循环 for range
 */
func main() {
	menu := [5]string{
		"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "蒜蓉粉丝扇贝",
	}

	for index, item := range menu {
		// 打印下标和名字 赋值给desc变量
		desc := fmt.Sprintf("%d-%s", index, item)
		fmt.Println(desc)
	}
}
