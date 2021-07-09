package main

import "fmt"

/**
动态类型 动态值
静态类型
*/

type ChefInterface interface {
	GetHonor() string
	Hidden(b bool) string
}

type Chef struct {
	Name  string
	Age   int
	Honor string
}

func (c Chef) Hidden(b bool) string {
	if b {
		return c.Name + "有隐藏绝活"
	}
	return ""
}

func (c Chef) GetHonor() string {
	return c.Honor
}

func (c *Chef) SetHonnor(title string) {
	c.Honor = title
}
func main() {
	// 创建对象
	zhang := Chef{
		Name:  "张师傅",
		Age:   20,
		Honor: "米其林1星",
	}
	fmt.Println(zhang.GetHonor())

	var ci ChefInterface = zhang  // 把zhang对象给接口
	zhang.SetHonnor("米其林3星") // 赋值

	fmt.Println(zhang.GetHonor())
	fmt.Println(ci.GetHonor())

	var zhang2 Chef
	var ci2 ChefInterface
	if ci2 == nil{
		fmt.Println("ci2 is nil")
	}

	ci2 = zhang2
	if ci2 == nil{
		fmt.Println("ci2 isnill,to")
	}

	zhang2.SetHonnor("米其林3星")
	fmt.Println(zhang2.GetHonor())
	fmt.Println(ci2.GetHonor())
}
