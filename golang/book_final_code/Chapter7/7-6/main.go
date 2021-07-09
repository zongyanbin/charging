package main

import (
	"fmt"
	"strconv"
)

/**
匿名函数 解决通用性解耦的问题
go语言 函数是一等公民
1.是函数
2.他没有名字
 */
type Discount func() float64
type CheckSum func(name string, price float64) float64

func PayOrder(discount Discount) CheckSum  {
	var total float64
	return func(name string, price float64) float64 {
		fmt.Println("菜品名称："+name+"单价:"+strconv.FormatFloat(price,'f',-1,64))
		total = total +price
		if discount ==nil {
			return total
		}
		return  total * discount()
	}
}
func main()  {
	f:=PayOrder(func() float64 {
		return 1
	})
	result := f("红烧肉", 88)
	fmt.Println(result)
	result = f("清蒸鱼", 98)
	fmt.Println(result)
	result = f("溜大虾", 128)
	fmt.Println(result)
	result = f("蒸螃蟹", 198)
	fmt.Println(result)
	result = f("蒜蓉粉丝扇贝", 68)
	fmt.Println(result)

	f2:=PayOrder(func() float64 {
		return 1
	})
	result2 :=f2("红烧肉",88)
	fmt.Println("result2:",result2)
}
