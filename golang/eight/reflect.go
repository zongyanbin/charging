
/**
ValueOf 用来获取输入参数接口中的数据的值，如果接口为空则返回0

TypeOf 用来动态获取输入参数接口中值的类型，如果接口为空则返回nil
 基本反射
 */
package main

import (
	"fmt"
	"reflect"
)

// 定义方法把值导出
func reflectNum(arg interface{})  {
	// arg 反射处理
	fmt.Println("type:", reflect.TypeOf(arg))
	fmt.Println("value", reflect.ValueOf(arg))
}

func main()  {
	var num float64 = 1.23456
	reflectNum(num)
}