package main

import "fmt"

func main()  {

	// 数组 第一种
	var foodList [5]string
	var priceList [5]int
	fmt.Println(foodList)
	fmt.Println(priceList)

	// 数组 第二种
	var foodList2 = [5]string{
		"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "蒜蓉粉丝扇贝",
	}
	fmt.Println(foodList2)

	// 数组 第三种
	foodList3 :=[5]string{
		"排骨","牛肉","红烧鱼","鸡爪子","猪蹄子",
	}
	fmt.Println(foodList3)

	// 类型大小必须一样才可以赋值
	var foodList4 [5]string
	foodList=foodList4
}
