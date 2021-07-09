package main

import "fmt"

func main()  {
	price:=1288
	if price<1288{
		fmt.Println("没有优惠")
	}else if price==1288 {
		fmt.Println("赠送优惠卡一张")
	}else {
		fmt.Println("赠送优惠卡两张")
	}
}
