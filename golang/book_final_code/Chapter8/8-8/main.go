package main

import (
	"fmt"
	"time"
)

/**
值方法与指针方法的却别
直接接收者
内部的影响会影响外面的数据

 */

type Honor3 struct {
	Titile string		// 获奖名称
	GetTime time.Time	// 获奖时间
}

type Cherf3 struct {
	Name string
	Age int
	Honor3
	Trainee *Cherf3
}

func (c Cherf3) Cook(name string) string {
	return c.Name +":做好了"+name+ "\n"
}

func (c Cherf3) FavCook(name string) string  {
	return c.Name+":这是我拿来菜" +name+",做好了。\n"
}
func main() {
	zhao :=&Cherf3{
		Name:    "张师傅",
		Age:     55,
		Honor3:  Honor3{},
		Trainee: nil,
	}

	fmt.Printf("%s",zhao.Cook("蛋炒饭"))
	fmt.Printf("%s", zhao.FavCook("小炒肉"))
}
