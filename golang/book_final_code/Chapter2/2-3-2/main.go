package main

import (
	"fmt"
	"strings"
)

/**
字符串合并（拼接）
*/

func main() {
	fondName := "红烧肉"
	fondType := "卤菜"
	fmt.Println(fondName + fondType)

	// strings.Join("切片")
	result1:=strings.Join([]string{"九转大肠","卤菜"}," 拼接上酸菜炖粉条 ")
	fmt.Println(result1)

	// .Sprintf(%s @ %s)
	result2 := fmt.Sprintf("%s@%s","麻婆豆腐","川菜")
	fmt.Println(result2)

	// bytes.Buffer{}
}
