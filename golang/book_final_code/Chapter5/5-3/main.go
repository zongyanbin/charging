package main

import "fmt"

func main()  {
	i:=0
	switch i {
	case 0:
		fmt.Println("红烧肉")
		fallthrough // 穿透向下执行 fallthrough
	case 1:
		fmt.Println("鲍鱼粥")


	}
}
