package main

import "fmt"

type Human struct {
	name string
	sex string
}

func (this *Human) Eat() {

	fmt.Println("Human.Eat()...")
}

func (this *Human) Walk() {
	fmt.Println("Huma.Walk()..")
}

//================================
// 继承
type Superman struct {
	Human // SuperMan类 继承了Human类的方法
	level int
}

// 重新定义父类的方法Eat()
func (this *Superman) Eat()  {
	fmt.Println("SuperMan .Eat()...")
}

// 子类的新方法
func (this *Superman) Fly() {
	fmt.Println("SuperMan .Fly()...")
}

func (this *Superman) Print() {
	fmt.Println("name=",this.name)
	fmt.Println("sex=",this.sex)
	fmt.Println("level=",this.level)
}
func main()  {
	h :=Human{
		name: "zhangsan3",
		sex:  "female",
	}
	h.Eat()
	h.Walk()

	// 定义一个子对象
	// s := Superman{Human{"LI44","female"},88}
	var s Superman
	s.name = "lisi"
	s.sex = "male"
	s.level = 88
	s.Walk() // 父类的方法
	s.Eat() // 子类的方法
	s.Fly() // 子类的方法
	s.Print()
}
