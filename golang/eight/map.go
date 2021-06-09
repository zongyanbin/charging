/**
slice 数组形式
map  哈希 key value 形式
两者的区别
 */
package main

import "fmt"

func printMap(cityMap map[string]string) {
	// cityMap 是一个引用传递
	for key,value := range cityMap{
		fmt.Println("key=",key,"value=",value)
	}
}

func ChangeValue(citMap map[string]string)()  {
	citMap["England"] ="London"
}
func main(){
	//1.  声明Map1是一种map类型 key是string, value是string
	var myMap1 map[string]string
	if myMap1 == nil{
		fmt.Println("myMap1是一个空map")
	}

	//1. 在使用map前，需要先用make给map分配数据空间 map哈希无序
	myMap1 = make(map[string]string,10)
	myMap1["one"] = "java"
	myMap1["two"] = "c++"
	myMap1["three"] = "php"
	fmt.Println("myMap1=",myMap1)
	fmt.Printf("len=%d,map=%v\n",len(myMap1),myMap1)

	//2. 第二种声明方式
	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "c++"
	myMap2[3] = "php"
	fmt.Println(" 第二种声明方式map\n",myMap2)

	//3. 第三种声明方式
	myMap3 := map[string]string{
		"one" :"java",
		"two" : "c++",
		"three": "php",
	}
	fmt.Println(" 第三种声明方式map\n",myMap3)

	fmt.Println("===========map使用方式=========")
	cityMap := make(map[string]string)
	// 增加
	cityMap["China"] = "北京"
	cityMap["JaPAN"] = "Tokoy"
	cityMap["USA"] = "NEWYORK"
	fmt.Println("==========函数map begin==========")
	printMap(cityMap)
	fmt.Println("==========函数map end==========")
	//遍历
	for key, value := range cityMap{
		fmt.Println("key=",key,"value=",value)
	}
	fmt.Println("==========删除==========")
	delete(cityMap,"China")
	fmt.Println("==========修改==========")
	cityMap["USA"]="SB"
	ChangeValue(cityMap)
	fmt.Println("======================")
	//遍历
	for key,value := range cityMap{
		fmt.Println("key=",key,"value=",value)
	}

	// copy 在定义一个map 重新依次赋值
}
