package main

import "fmt"


 /**
 切片扩容
 原切片一起复制到新切片中
 小于1024 新切片不超过原切片扩容就是2倍

 如果大于1024 就是就切片的1.25倍
  */
func main() {
	menu := [5]string{
		"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "蒜蓉粉丝扇贝",
	}

	// 第一种
	var s1 []string = menu[0:3]
	fmt.Println(s1)

	// 第二种
	//Slice
	var s2 = []string{
		"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "蒜蓉粉丝扇贝",
	}
	fmt.Println(s2)

	// 第三种 长度为6
	s3 := make([]string, 6)
	r := fmt.Sprintf("len:%d,cap:%d", len(s3), cap(s3))
	fmt.Println(r)
	// 长度为6容量为8
	s4 := make([]string, 6, 8)
	r2 := fmt.Sprintf("len%d,cap:%d", len(s4), cap(s4))
	fmt.Println(r2)

	s5:=[]string{"红烧肉"}
	r5:=fmt.Sprintf("len%d,cap:%d", len(s5), cap(s5))
	fmt.Println(r5)

	s5=append(s5,"清蒸鱼")
	r5= fmt.Sprintf("len%d,cap:%d", len(s5), cap(s5))
	fmt.Println(r5)

	s5=append(s5,"溜大虾")
	r5 = fmt.Sprintf("len%d,cap:%d", len(s5), cap(s5))
	fmt.Println(r5)

	s5=append(s5,"蒸螃蟹")
	r5 = fmt.Sprintf("len%d,cap:%d", len(s5), cap(s5))
	fmt.Println(r5)

	s5=append(s5,"鲍鱼粥")
	r5 = fmt.Sprintf("len%d,cap:%d", len(s5), cap(s5))
	fmt.Println(r5)
}
