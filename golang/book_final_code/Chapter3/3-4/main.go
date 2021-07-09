package main

import "fmt"

// 数组都是从0开始
func main() {
	var list2 = [5]string{
		"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "蒜蓉粉丝扇贝",
	}
	fmt.Println(list2[2])
}