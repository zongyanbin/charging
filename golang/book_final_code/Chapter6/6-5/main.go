package main

import "fmt"

func main()  {
	var m map[string]int = map[string]int{
		"红烧肉": 88,
		"清蒸鱼": 98,
		"溜大虾": 128,
		"蒸螃蟹": 198,
		"鲍鱼粥": 68,
	}
	prices :=[]*int{}
	for k ,v  :=range m{
		fmt.Println("k:[%p], V:[%p]",&k,&v)
		//prices = append(prices,&v) 错误提示
		vv:=m[k] //取地址用map 键值取，否则取的都是最后一个地址
		prices =append(prices, &vv)
	}

	for _,p :=range prices{
		fmt.Println(*p)
	}
}
