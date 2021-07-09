package main

import "fmt"
/**
数组循环 for range
debugger 断点调试 从上到下执行
console 最后看控制台
*/
func main() {
	menu := [5]string{
		"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "蒜蓉粉丝扇贝",
	}

	for index, item := range menu {
		// 打印下标和名字 赋值给desc变量
		/**
			%d	代表数字
			%s	代表字符串
		*/
		desc := fmt.Sprintf("%d-%s", index, item)
		fmt.Println(desc)
	}
}
