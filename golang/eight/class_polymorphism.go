package main

import "fmt"

// 基本要素：d
//	1.有个父类有一个接口
//	2.有子类（实现了父类得全部接口方法）
//	3.父类类型的变量（指针）指向（引用）子类得具体数据变量
// 多态实现与基本要素 需要有 interface 接口
// 接口定义-->子类继承重写 ->达成多态关系

// 本质是一个指针
// 一个接口实现两个具体的类
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

func (this *Dog) Sleep() {
	fmt.Println("Dog is Sleep")
}

func (this *Dog)GetColor() string {
	return this.color
}

func (this *Dog) GetType() string  {
	return "Dog"
}

func showAnima(animal AnimaiIF)  {
	animal.Sleep() // 多态
	fmt.Println("color=",animal.GetColor())
	fmt.Println("kind=",animal.GetType())
}
func main()  {
/* // 用接口来触发做他的条件
	var animal AnimaiIF // 接口的数据类型， 父类指针
	animal = &Cat{"Green"}
	animal.Sleep() //调用的就是Cat的Sleep()方法

	animal = &Dog{"Yellow"}
	animal.Sleep() // 调用Dog的Sleep方法，多态的现象
*/

	cat := Cat{"Green"}
	dog := Dog{"Yellow"}

	showAnima(&cat)
	showAnima(&dog)
}