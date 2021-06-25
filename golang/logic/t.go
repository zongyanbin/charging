package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main()  {

	fmt.Println("=============for====变形体1==============")
	lesson := [8]int{50, 80, 78, 36, 70, 85, 90, 87}
	var sum int
	for i:=0; i<8; i++{
		sum = sum+lesson[i]
	}
	fmt.Println("sum=",sum)

	x, _:= strconv.Atoi("12")
	println(x)

	avg := sum / len(lesson)
	fmt.Println(avg)

	xx := 42
	for xx>5 {
		xx--
		fmt.Println("xx=",xx)
	}
	fmt.Println("=============for i v range======变形体2============")
	for i, v := range lesson{
		fmt.Println(i,v)
	}
	fmt.Println("=============for{ do }=====变形体3============")

	i := 0
	for  {
		i++
		if(i>7){
			break
		}
		fmt.Println(lesson[i])
	}

	fmt.Println("=============切片 slice============")
	arr :=[]int{50,80,78,36,70,85,90,89}
	s1 := arr[0:2]
	fmt.Println(arr,reflect.TypeOf(arr),s1,reflect.TypeOf(s1))
	fmt.Println("=============func============")

	aa := PrintScore(arr)
	fmt.Println(aa)
	fmt.Println("=============func 匿名函数============")
	func (s int){
		println(s)
	}(11111111)
	//2
	add := func(x,y int) int {
		return x+y
	}
	println(add(1,2))
}
func PrintScore(s []int) (avg int ){
	sum :=0
	for i:=0; i<len(s); i++ {
		sum =sum+s[i]
	}
	avg = sum/len(s)
	return
	//return avg
}