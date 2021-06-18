package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	// 设置种子，只需一次
	rand.Seed(time.Now().UnixNano())
	var a [10]int
	n := len(a)
	for i:=0; i<n; i++ {
		a[i] = rand.Intn(100) //100以内随机数
		fmt.Println("%d", a[i])
	}
	fmt.Println("\n")

	// 冒泡排序 按这两个元素比价，升序（大于则交换）

	for i := 0, i<n-1; i++ {
		for j :=0; j<n-1-i; j++ {
			if a[j]>a[j+1]{
				a[j], a[j+1] = a[j+1],a[j]
			}
		}
	}
}
