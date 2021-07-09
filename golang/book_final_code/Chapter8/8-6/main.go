package main

import (
	"fmt"
	"time"
)

/**
指针接收者方法 省内存

方法 面向对象的编程 看起来和函数很像
方法	：先要有接收者才能调用，然后创建一个对象
函数	：写好了直接用就可以
*/

type Chef struct {
	Name string
	Age  int
	Honor
	Trainee *Chef
}

type Honor struct {
	Title  string    // 获奖名称
	Getime time.Time // 获奖时间
}

//	接收者
func (c *Chef) Cook(name string) string {
	return c.Name + "：做好了" + name
}

func (c *Chef) FavCook(name string) string {
	return c.Name+"：这是我的拿手菜"+name+",做好了。"
}

func main()  {

	// 第一种 常见
	li := &Chef{
		Name:    "李师傅",
		Age:     20,
		Honor:   Honor{},
		Trainee: nil,
	}
	// 接收者 一定要有对象 c.name 对象的名字
	result :=li.Cook("红烧肉")
	fmt.Printf("%s", result)
	result = li.FavCook("葱烧海参")
	fmt.Printf("%s", result)

	// 第二种 常见
	li2 :=Chef{
		Name:    "王师傅",
		Age:     20,
		Honor:   Honor{},
		Trainee: nil,
	}
	liPoint := &li2
	result2 :=liPoint.Cook("猪蹄子")
	fmt.Printf("%s", result2)
	result2 = liPoint.FavCook("葱烧海参")
	fmt.Printf("%s", result2)

	// 第三种  很少用
	li3 :=Chef{
		Name:    "马师傅",
		Age:     30,
		Honor:   Honor{},
		Trainee: nil,
	}
	reslut3 :=(&li3).Cook("马肉")
	fmt.Printf("%s", reslut3)
	result2 = liPoint.FavCook("葱烧海参")
	fmt.Printf("%s", reslut3)

}