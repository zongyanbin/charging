/**
结构体：
	面向对象封装
 type struct 属性
尽量少修改别人的代码
软件工程
尽量少改别人的代码  少出BUG
尽量少些代码		 少出BUG

结构体嵌入到其它结构体就是组合
相当于搭积木自由组合
 */
package main

import "time"

// 厨师
type Chef struct {
	Name 	string // 姓名
	Age 	string // 年龄
	Honor			// 荣誉
	Traninee	*Chef	// 徒弟，这里为了演示，可以认为徒弟有多个用 []*Chef 用切片表示
}

//	荣誉
type Honor struct {
	Title 		string 	// 获奖名称
	GetTime 	time.Time 	// 获奖时间
	Publisher	string	// 荣誉的颁发者
}

// 结构体初始化使用
//type struct{} struct{}{}
func main()  {

}
