package main

import "fmt"
/**
数字遍历
 */
func main() {
	s := []string{"红烧肉", "清蒸鱼", "溜大虾", "鲍鱼粥"}
	for index, item := range s {
		r:=fmt.Sprintf("%d---- %s",index,item)
		fmt.Println("-----------")
		fmt.Println(r)
	}
}
