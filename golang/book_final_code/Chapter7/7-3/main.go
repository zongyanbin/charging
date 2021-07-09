package main

import (
	"fmt"
	"net/http"
	"strconv"
)
/**
多返回值
有错误就显示错误
没错误就正常
 */
func main()  {
	m:=map[string]int{"红烧肉":88}
	v,ok:=m["麻辣小龙虾"]  // v 找到字典里的值  ok存在, !ok不存在
	if !ok{
		fmt.Println("没有找到你要的菜品")
	}else{
		fmt.Println("找到了"+strconv.Itoa(v))
	}
	//fmt.Println("找到了"+strconv.Itoa(v))
	res, err:=http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(res)
	}
}
