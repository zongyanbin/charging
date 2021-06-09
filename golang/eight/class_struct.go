package main

import "fmt"

// 结构体首字母大学表示public
// 属性首字母大写可以外包调用，首字母小写只能本包使用
// begin class
type Hero struct {
	Name string
	Ad int
	Level int
}
/*
 这种方法读可以写不行需要改进
func (this Hero) Show()  {
	fmt.Println("hero=",this)
	fmt.Println("============分别打印========")
	fmt.Println("Name=",this.Name)
	fmt.Println("Ad=",this.Ad)
	fmt.Println("Level=",this.Level)
}

func (this Hero) GetName()string {
	return this.Name
}

func (this Hero) SetName(newName string)  {
	// this 是调用该方法的对象的一个副本（拷贝）
	this.Name = newName
}
*/
// end class

func (this *Hero) Show()  {
	fmt.Println("hero=",this)
	fmt.Println("============分别打印========")
	fmt.Println("Name=",this.Name)
	fmt.Println("Ad=",this.Ad)
	fmt.Println("Level=",this.Level)
}

func (this *Hero) GetName()string {
	return this.Name
}

func (this *Hero) SetName(newName string)  {
	this.Name = newName
}


func main()  {

	// 创建一个对象
	hero := Hero{
		Name: "dzhangsan",
		Ad: 100,
		Level:1,
	}

	hero.Show()
	hero.SetName("李四")
	fmt.Println("=======setName======")
	hero.Show()
}