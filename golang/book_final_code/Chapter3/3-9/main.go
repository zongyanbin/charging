package main

import "fmt"

func main()  {
	menu := [5]string{
		"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "蒜蓉粉丝扇贝",
	}
	menu[0]="猪蹄子"
	// invalid array index 99 (out of bounds for 5-element array)
	// 数组下标越界 menu[99]
	fmt.Println(menu[3])
}
