package main

import "fmt"

func main() {
	s := []string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "鲍鱼粥"}

	for _,item:=range s{
		if "清蒸鱼" ==item{
			fmt.Println("今日免费送此菜")
		}else {
			fmt.Println(item)
		}
	}
}
