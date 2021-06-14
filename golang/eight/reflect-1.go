/**
复杂反射
 */
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id int
	Name string
	Age int
}

// Call首字母一定要大写
func (this User) Call() {
	fmt.Println("user is called...")
	fmt.Printf("%v\n",this)
}

func main() {
	user := User{1, "Aceld", 18}
	DoFiledAndMethod(user)
}

 // 定义方法传递任意数据类型
func DoFiledAndMethod(input interface{})  {
	// 1.	获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is:", inputType.Name()) // inputType.Name() 打印当前Name名称

	// 2.	获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is:", inputValue)

	// 3.	通过 type 获取里面的字段

	//1. 获取 interface的reflect.type，通过Type得到NumField，有多少个字段,进行遍历
	//2. 得到每个field，数据类型
	//3. 通过filed有一个Interface()方法等到 对应的value
	for i := 0; i < inputType.NumField(); i++{
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}
	//
	fmt.Printf("%v",inputType.NumMethod())
	// 4.	通过 input 获取里面的方法，调用 遍历
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}

	fmt.Println(inputType.NumMethod())
}