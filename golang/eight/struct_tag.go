package main

import (
	"fmt"
	"reflect"
)

type resum struct {
	Name string		`info:"name" doc:"我的名字"` // 标签说明 判断在包里怎么用
	Sex string		`info:"sex"`
}

// 通过反射解析 结构体标签
func findTag(str interface{})  {
	t := reflect.TypeOf(str).Elem()
	for i:=0;i<t.NumField();i++{
		tagsinfo :=t.Field(i).Tag.Get("info")
		tagdoc :=t.Field(i).Tag.Get("doc")
		fmt.Println("info:",tagsinfo)
		fmt.Println("doc:",tagdoc)
	}
}
func main(){
	var re resum
	findTag(&re)
}