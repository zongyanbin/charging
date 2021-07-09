package main

import "fmt"

/**
使用指针接收者实现接口：
*/

// 厨师接口
type ChefInterface interface {
	Cook() bool                   // 做菜
	FavCook(foodName string) bool // 拿手菜
}

// 厨师结构体
type Chef struct {
	Name string
	Age  int
}

func (c *Chef) Cook() bool {
	fmt.Println(c.Name + "饭菜做好了")
	fmt.Printf("%p\n", c)
	return true
}

func (c *Chef) FavCook(foodName string) bool {
	fmt.Println(c.Name + "的拿手菜" + foodName + "做好了")
	fmt.Printf("%p\n", c)
	return true
}

func WorkForDinner(c *Chef, foodsName string) {
	// c *Chef	指针赋值给接口变量
	var ci ChefInterface = c   // 调用接口 ChefInterface

	// 然后用接口方法 去调用接口内的方法
	ci.Cook() // 做饭
	ci.FavCook(foodsName) // 拿手菜
}
func main() {

	// 实现多态性 代码数量减少
	li := Chef{
		Name: "李师",
		Age:  29,
	}
	fmt.Printf("李师傅对象的地址%p\n", &li)
	WorkForDinner(&li, "红烧肉")

	zhang := Chef{
		Name: "张张",
		Age:  22,
	}
	fmt.Printf("张张对象的地址%p\n", &zhang)
	WorkForDinner(&zhang, "猪蹄子")

}
