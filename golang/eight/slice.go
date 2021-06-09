/**
	动态数组
 */
package main

import "fmt"

func  printArray(myArray []int)  {
	// 切片引用传递
	// _表示匿名的变量
	for _,value := range myArray{
		fmt.Println("value=",value)
	}

	myArray[0] =110
}
func main(){
	myArray := []int{1,2,3,4} //动态数组，切片 slice
	fmt.Println("myArray type is %T\n",myArray)

	printArray(myArray)

	fmt.Println("===================")
	for _,value := range myArray{
		fmt.Println("value=",value)
	}

	// 1. 声明slice切片，并且初始化，默认值为1，2，3 长度为len是3
	//slice1 := []int{1,2,3}

	// 2. 声明slice1是一个切片，但是并没有给slice分配空间
	//var slice1 []int
	//slice1 = make([]int,3) //开辟3个空间，默认值是0

	// 3. 声明一个slice是一个切片，同时给slice分配空间，3个空间，初始化为0
	//通过：=推导出slice是一个切片
	//var slice1 []int = make([]int,3)

	// 4.
	//slice1 := make([]int,4)
	//slice1[0]=111

	var slice1 []int
	fmt.Printf("len=%d,slice=%v\n",len(slice1),slice1)

	// 判断一个 slice 是否为空是否合法不是值为空，而是元素和空间

	if slice1 == nil {
		fmt.Println("slice1是一个空切片")
	}else{
		fmt.Println("slice 是有空间的")
	}

	// 切片容量的追加，切片的截取
	fmt.Println("=========切片容量的追加，切片的截取========")
	var numbers = make([]int,3,5)
	fmt.Printf("len=%d,cap=%d,slice=%v\n",len(numbers),cap(numbers),numbers)

	// 向numbers切片里追加一个元素1，numbers len=4,[0,0,0,1], cap=5
	numbers = append(numbers,1)
	fmt.Printf("len=%d,cap=%d,slice=%v\n",len(numbers),cap(numbers),numbers)

	numbers = append(numbers,2)
	fmt.Printf("len=%d,cap=%d,slice=%v\n",len(numbers),cap(numbers),numbers)

	// 向一个容量cap已满的slice 追加元素
	numbers = append(numbers,3)
	fmt.Printf("len=%d,cap=%d,slice=%v\n",len(numbers),cap(numbers),numbers)

	fmt.Println("===============切片len 等于左右指针===========================")
	fmt.Println("===============切片cap 动态追加cap2倍=========================")
	var numbers2 =make([]int,3)
	fmt.Printf("len=%d,cap=%d,slice=%v\n",len(numbers2),cap(numbers2),numbers2)
	numbers2 = append(numbers2,1)
	fmt.Printf("len=%d,cap=%d,slice=%v\n",len(numbers2),cap(numbers2),numbers2)

	fmt.Println("===============切片的截取=========================")
	s := []int{1,2,3,4,5}

	s1 := s[0:2]
	fmt.Println(s1)

	// copy 可以将底层数组的slice一起进行拷贝
	s2 := make([]int,5) // s2=[0,0,0]
	copy(s2,s) //将s中的值依次拷贝到s2中
	fmt.Println(s2)
}
