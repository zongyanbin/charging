<<<<<<< HEAD
=======
/**
明确目标（确定在哪个网站搜索）
爬（爬下内容）
取（筛选想要的）
处理数据（按照你的想法去处理）
 */
>>>>>>> 519a0d541ab7aa6d7e274a1273e4da6eb32925e2
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

<<<<<<< HEAD
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
=======
// 这个是一个简单版本获取qq邮箱并且进行封装操作，另外爬出来的数据也没有进行去重操作

var (
	reQQEmail = `(\d+)@qq.com`
)

// 爬邮箱
func GetEmail()  {
	// 1.去网站拿数据
	resp, err := http.Get("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
	HandleError(err, "http.Get url")
	defer resp.Body.Close()
	// 2.读取页面内容
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll")
	// 字节转字符串
	pageStr := string(pageBytes)
	//fmt.Println(pageStr)
	// 3.过滤数据，过滤qq邮箱
	re := regexp.MustCompile(reQQEmail)
	// -1代表取全部
	results := re.FindAllStringSubmatch(pageStr, -1)
	//fmt.Println(results)

	// 遍历结果
	for _, result := range results {
		fmt.Println("email:", result[0])
		fmt.Println("qq:", result[1])
	}
}

// 处理异常
func HandleError(err error, why string){
	if err != nil{
		fmt.Println(why, err)
	}
}

func main()  {
	GetEmail()
>>>>>>> 519a0d541ab7aa6d7e274a1273e4da6eb32925e2
}