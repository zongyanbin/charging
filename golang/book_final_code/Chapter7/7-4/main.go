package main

/**
函数变量

函数的调用者只关心你和我的类型签名一样就可以
执行到main函数 不关心你函数里面了
总的来说：把不同函数调用功能，让调用者去定义，不用执行者去定义
*/

import "fmt"

// 定义一个函数类型
type Hi func(num string) string

func Hello(num string) string {
	return num + "位客人，欢迎光临"
}

func Hello4DongBei(num string) string {
	return num + "位兄弟，欢迎光临"
}

func Hello4Hannan(num string) string {
	return "雷猴" + num + "位兄弟，欢迎光临"
}

/**
 SayHello  代码稳定
 Sayhello尽量不要修改以前代码，容易产生bug
 尽量用
 */
func SayHello(num string, hi Hi) {
	result := hi(num)
	fmt.Println(result)
}

func main() {
	var hello Hi
	//var xx string
	hello = Hello
	words := hello("3")
	fmt.Printf("%s\n", words)

	hello = Hello4DongBei
	words = hello("5")
	fmt.Printf("%s\n", words)

	target := "海南"
	if target == "东北食客" {
		SayHello("3", Hello4DongBei)
	} else if target == "海南" {
		SayHello("10", Hello4Hannan)
	} else {
		SayHello("6", Hello)
	}

}
