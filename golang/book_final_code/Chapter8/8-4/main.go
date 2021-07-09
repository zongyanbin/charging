/**
结构体与JSON
数据库数据给前端都是json
 */
package main

import (
	"encoding/json"
	"fmt"
)
 // 定义厨师结构体对象
type Chef struct {
	Name 	string
	Age 	int
}
func main()  {
	c:=Chef{
		Name: "老王",
		Age:  28,
	}

	// 对象序列化成	json	字符串
	marsha1,err := json.Marshal(&c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(marsha1)) // 转换string 类型

	var cc Chef
	s:=`{"Name":"老王","Age":28}`
	// 反序列化 Unmarshal json转换成对象
	err = json.Unmarshal([]byte(s), &cc)  // &cc 取地址符 可以修改里面字段
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cc.Name)
	fmt.Println(cc.Age)
}
