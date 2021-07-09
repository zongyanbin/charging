package main

import "fmt"

func main()  {
	s := []string{"红烧肉", "清蒸鱼", "溜大虾", "鲍鱼粥"}
	s = append(s,"干炒牛河")
	fmt.Println(s)

	// 扩展
	// foods:=make([]string,5,5)
}
