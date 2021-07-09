package main
import "fmt"
//panic 终止程序不执行
func main()  {
	foods := []string{"红烧肉", "清蒸鱼", "熘大虾", "蒸螃蟹", "鲍鱼粥"}

	fmt.Printf("%s", foods[5])
}
