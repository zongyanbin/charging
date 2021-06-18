package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	// 数组 ，同一个类型的集合
	var id [50]int
	// 操作数组，通过下标，从0开始，到len()-1， 赋值每个元素
	for i := 0; i < len(id); i++ {
		id[i]=i+1
		fmt.Printf("id[%v]=%v\n",i,id[i])
	}

	// 注意： 定义数组是，指定的数组元素必须是常量
	// 操作数组元素，是从0开始，到len()-1,不对称元素，这个数字，叫下标
	// 这个下标，可以是变量或常量

	// 打印
	// 第一个返回下标，第二个返回元素
	for i, data := range id{
		fmt.Printf("id[%d]=%d\n",i,data)
	}

	// 声明定义赋值，叫初始化
	// 1、全部初始化
	var a [5]int = [5]int{1,2,3,4,5}
	fmt.Println("a=",a)
	b:=[5]int{1,2,3,4,5}
	fmt.Println("b=",b)

	// 2、部分初始化，没初始化的元素，自动赋值为0
	c := [5]int{1,2,3}
	fmt.Println("c=",c)

	// 指定某个元素初始化
	d :=[5]int{2:10,4:20}
	fmt.Println("d=",d)

	// 二位数组
	TwoArray()

	// 生成随机数
	RandArray()
}

// 有多少个 [] 就是多个维
// 有多少个 [] 就用多少个循环
func TwoArray()  {
	var a [3][4]int

	 k := 0
	for i:=0; i<3; i++ {
		for j:=0; j<4; j++ {
			k++
			a[i][j] = k
			fmt.Printf("a[%d][%d] = %d, ", i, j, a[i][j])
		}
	}
	fmt.Printf("\n",a)

	// 二维数组初始化
	b := [3][4]int{ {1,2,3,4}, {5,6,7,8}, {9,10,11,12}}
	fmt.Println("b=",b)

	// 部分初始化
	c := [3][4]int{ {1,2,3,4}, {5,6,7,8}, {9,10}}
	fmt.Println("c=",c)

	// 数组支持比较，只支持 == 或 !=, 比较是不是每一个元素都一样，2个数组比较，类型也要一样

	// 同类型的数组可以赋值

}

// go 随机数
func RandArray()  {
	// 设置种子，只需要一次
	// 如果种子参数一样，每次运行程序产生的随机数都一样
	rand.Seed(time.Now().UnixNano()) // 以当前系统时间作为种子参数
	for i:=0; i<5; i++ {
		// 产生随机数
		// fmt.Println("rand=",rand.Int()) //随机很大的数

		// 限制在100以内的数
		fmt.Println("rand=",rand.Intn(100))
	}

}