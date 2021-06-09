package main

import "fmt"

// 声明一种新的数据类型 myint 是int的一个别名
type myint int

// 定义一个结构体
type Book struct {
	title string
	auth string
}

func changeBook(book Book)  {
	// 传递一个book的副本
	book.auth =  "666"
}

func changeBook2(book *Book)  {
	book.auth =  "指针888"
}
func main()  {
	var book1 Book
	book1.title ="golang"
	book1.auth = "张三"

	fmt.Printf("%v\n",book1)
	changeBook(book1)
	fmt.Printf("%v\n",book1)
	fmt.Println("=========book2 指针=====")
	changeBook2(&book1)
	fmt.Printf("%v\n",book1)

}
