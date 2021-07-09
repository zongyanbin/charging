package main
/*
使用制表符和换行符号
 */
import "fmt"

func main() {
	message:="hello go world"
	fmt.Println(message)

	//\t 制表符空白  \n换行
	fmt.Println("\t"+message)
	fmt.Println("\n"+message)
	fmt.Println(message+"\n\t"+message)
}
