/**
通过万能类型 ：interface{} 空接口
interface{}类型可以，引用任意数据类型
 */
package main

import "fmt"

// interface{} 万能数据类型
func myFunc(arg interface{}){
 //fmt.Println("myFunc is called...")
 //fmt.Println(arg)

	// interface{} 如何区分 此时引用的底层数据类型到底是什么？
	// 给interface{} 提供	【类型断言】	机制
     value, ok:=arg.(string)
     if !ok{
		 fmt.Println("arg is no string type")
	 } else{
	 	 fmt.Println("arg is string type ,value=",value)
		 fmt.Printf("value type  is %T\n",value)
	 }
}
type Book struct {
	Auth string
}
func main(){
	 book :=Book{"test"}
	 myFunc(book)
	 myFunc(100)
	 myFunc("ABC")
	 myFunc(3.14)
}