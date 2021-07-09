package main

import "fmt"

/**
多条件判断
 */
func main() {
	s := []string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "蒜蓉粉丝扇贝"}
	for _,item :=range s{
		//if item=="红烧肉" && item=="清蒸鱼"{
		//	fmt.Println("送一个猪蹄子")
		//}else {
		//	fmt.Println(item)
		//}

		if item=="红烧肉" || item=="清蒸鱼"{
			fmt.Println(item,"送一个猪蹄子")
		}else {
			fmt.Println(item)
		}


	}
}