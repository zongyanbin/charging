package main

import "fmt"

func main()  {
	s := []string{"红烧肉", "清蒸鱼", "溜大虾", "鲍鱼粥"}
	s2:=s[:]
	fmt.Println(s2)

	//s3 := []string{"红烧肉", "清蒸鱼", "溜大虾", "鲍鱼粥"}
	//s4 :=[]string{"米饭","面条"}
	//result:=copy(s3,s4)
	//fmt.Println(s3)
	//fmt.Println(s4)
	//fmt.Println(result)

	s5 := []string{"红烧肉", "清蒸鱼", "溜大虾", "鲍鱼粥"}
	s6 :=[]string{"米饭","面条"}
	result:=copy(s5,s6)
	fmt.Println(s5)
	fmt.Println(s6)
	fmt.Println("result：",result)
}
