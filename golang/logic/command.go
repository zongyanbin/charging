/**
命令行工具：
os.Args
切片
循环
错误处理
函数
 */
package main
import (
	"fmt"
	"os"
	"strconv"
)
func main()  {
	var s[]int
	//fmt.Println(os.Args)
	//fmt.Println(reflect.TypeOf(os.Args))
	osArgs :=os.Args[1:]
	for i := 0; i<len(osArgs); i++ {
		item, err := strconv.Atoi(osArgs[i])
		if err != nil{
			fmt.Println("strconv err", err)
		}
		s = append(s,item)
	}
	fmt.Println(s)
	fmt.Println(getAvg(s))
}

/**
求平均值函数
 */
func getAvg(s []int) (avg int) {
	sum := 0
	for i :=0; i<len(s); i++ {
		sum +=s[i]
	}
	avg = sum/len(s)
	return
}
