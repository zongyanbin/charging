package main

import "fmt"

// 多态实现与基本要素 需要有 interface 接口
// 接口定义-->子类继承重写 ->达成多态关系

// 本质是一个指针
type  AnimaiIF interface {
	Sleep()
	GetColor() string // 获取动物颜色
	GetType() string // 获取动物的种类
}

// 具体的类
type Cat struct {
	color string // 猫的颜色
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is Sleep")
}

func (this *Cat) GetColor() string  {
	return this.color
}

func (this *Cat) GetType() string  {
	return "Cat"
}

// 具体的类
type  Dog struct {
	color string
}


func main()  {
	
}