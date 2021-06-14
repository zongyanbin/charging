package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Write interface {
	WriteBook()
}

// 具体类型
type Book struct {

}

func (this *Book) ReadBook(){
	fmt.Println("Read a Book")
}

func (this *Book) WriteBook(){
	fmt.Println("Write a Book")
}

func main(){
	var a string
	//pair <stype:string>, value:"aceld">
	a = "aceld"

	//pair <type:string, value:"aceld">
	var allType interface{} //万能指针

	// 断言
	allType = a
	str,_:=allType.(string)
	fmt.Println(str)

	fmt.Println("===================================")

	// pair <type:Book,value:book{}地址>
	b := &Book{}

	// r:pair<type:,value>
	var r Reader
	// r: pair<type:Book,value:book{]地址>
	r=b

	r.ReadBook()

	var w Write
    // r:pair<type:Book,value:book{}地址>
	w=r.(Write) // 此处的断言为什么会成功？因为w r具体的type是一一致
	w.WriteBook()
}