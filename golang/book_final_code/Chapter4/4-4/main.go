package main

import "fmt"

func main() {
	s := []string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "鲍鱼粥"}
	s2:=make([]*string,len(s))

	for i,_ := range s{
		// s2[i] = &v 错误地址都一样
		s2[i] = &s[i]
		fmt.Println(s[i])
	}
	fmt.Println(s2)
}
