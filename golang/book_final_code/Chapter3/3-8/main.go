package main

import "fmt"

func main()  {
	menu := [5]string{
		"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "蒜蓉粉丝扇贝",
	}
	menu[0]="猪蹄子"
	fmt.Println(menu)
}
