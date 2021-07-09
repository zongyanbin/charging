package main

import "fmt"

func main() {
	menu := [5]string{
		"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "蒜蓉粉丝扇贝",
	}
	msg := "您点的菜：" + menu[1]
	fmt.Println("您点的菜是：" + msg)
}
