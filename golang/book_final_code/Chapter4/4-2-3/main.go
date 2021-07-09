package main

import "fmt"
/**
切片删除
 */
func main() {
	s := []string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "鲍鱼粥","三文鱼"}
	fmt.Println(s[2:4])
	fmt.Println(s[2:])
	fmt.Println(s[:3])

	result :=append(s[:2],s[3:]...)
	fmt.Println(result)
}
