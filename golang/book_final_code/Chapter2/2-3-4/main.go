package main

import (
	"fmt"
	"strings"
)
/**
删除空格函数
strings.Trim(foods, " ")
 */
func main()  {
	foods:="    米饭, 馒头,饺子,包子"
	r:=strings.Trim(foods, " ")
	fmt.Println(r)
}
