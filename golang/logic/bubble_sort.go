package main

import (
	"fmt"
	"math/rand"
	"time"
)

func InitData(s []int)  {
	// 设置种子
	rand.Seed(time.Now().UnixNano())
	for i :=0; i<len(s); i++ {
		s[i] = rand.Intn(100) // 100以内的随机数
	}
}

// 冒泡排序
func BubbleSort (s []int){
	n := len(s)
	for i := 0; i < n-1; i++ {
		for j := 0; j< n-1-i; j++ {
			if s[j] > s[j+1]{
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

func main()  {
	// 设置种子，只需一次
	rand.Seed(time.Now().UnixNano())
	var a [10]int
	n := len(a)

	for i:=0; i<n; i++ {
		a[i] = rand.Intn(100) //100以内随机数
		fmt.Printf("%d ", a[i])
	}
	fmt.Println("\n")

	// 冒泡排序 按这两个元素比价，升序（大于则交换）

	for i := 0; i<n-1; i++ {
		for j :=0; j<n-1-i; j++ {
			if a[j]>a[j+1]{
				a[j], a[j+1] = a[j+1],a[j]
			}
		}
	}

	fmt.Printf("\n排序后：\n")
	for i := 0; i<n; i++{
		fmt.Printf("%d, ", a[i])
	}
	fmt.Printf("\n")

	array := []int{0,1,2,3,4,5,6,7,8,9}

	s1 := array[:]
	fmt.Printf("len=%d,cap = %d\n",len(s1),cap(s1))

	data := array[1]
	fmt.Println("data=",data)

	s2 := array[1:6:7]

	fmt.Printf("len=%d,cap=%d\n",len(s2),cap(s2))

	fmt.Println("==============切片做函数参数==================")

	arr := 10
	// 创建一个切片, len为n
	s := make([]int, arr)
	InitData(s) // 初始化数组
	fmt.Println("排序前： ", s)

	BubbleSort(s) // 冒泡拍戏
	fmt.Println("排序后： ", s)
}
