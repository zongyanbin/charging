package main

import "fmt"

func main()  {
	s := []string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "鲍鱼粥"}
	fmt.Println(s[0])
	fmt.Println(s[2])
	fmt.Println(s[len(s)-1])
}
