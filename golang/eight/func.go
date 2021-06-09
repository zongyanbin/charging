package main

import "fmt"

func fool(a string, b int) int {

	fmt.Println("a=",a)
	fmt.Println("b=",b)

	c := 100
	return c
}

// 返回多个返回值 匿名
func foo2(a string, b int)(int,int){
	fmt.Println("a=",a)
	fmt.Println("b=",b)
	return 666,777
}

// 返回多个返回值 有形参名
func foo3 (a string, b int) (r1 int, r2 int)  {
	fmt.Println("----foo3----")
	fmt.Println("a=",a)
	fmt.Println("b=",b)

	r1=1000
	r2=2000
	return
}

// 形参是相同类型
func foo4(a string,b int)(r1,r2 int)  {
	fmt.Println("----foo4----")
	fmt.Println("a=",a)
	fmt.Println("b=",b)

	r1=1000
	r2=2000
	return
}
func main()  {
	c := fool("abc", 555)
	fmt.Println("c=",c)

	ret1, ret2 := foo2("hh",999)
	fmt.Println("ret1=",ret1,"ret2=", ret2)

	ret1, ret2 = foo3("foo3",300)
	fmt.Println("ret1=",ret1,"ret2=",ret2)

	ret1, ret2 = foo4("foo4",400)
	fmt.Println("ret1=",ret1,"ret2=",ret2)
}

