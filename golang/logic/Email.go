package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func Gemail(url string) []string{
	resp,err:=http.Get(url) //读取网页
	if err !=nil{
		return []string{}
	}
	page,err:=ioutil.ReadAll(resp.Body) //读取网页
	//fmt.Println(string(page))

	//reg:=`[\w-]+(\.[\w-]+)*@[\w-]+(\.[\w-]+)+` // 编译正则表达式\
	reg := `(\d+)@qq.com`
	rgx:=regexp.MustCompile(reg) // 编辑正则表达式
	return rgx.FindAllString(string(page),-1)
}

func main(){
	fmt.Println(Gemail("https://tieba.baidu.com/p/6051076813?red_tag=1573533731"))
}

func  str()  {
	// 包相关  import package
	// 声明相关 var const type    struct interface func		channel  map go
	// 循环声明 for range
	// 条件相关 if else switch case select
	// 中断返回	return break default continue goto fallthrough

	// 延迟执行 defer

	//true false iota nil

	// make len cap new append copy close delete complex panic recover
}