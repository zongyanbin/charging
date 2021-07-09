package main

import "fmt"

func start()  {
	fmt.Println("程序开始执行...")
}

//写的切片数组越界
func testFood()  {
	foods := []string{"红烧肉", "清蒸鱼", "熘大虾", "蒸螃蟹", "鲍鱼粥"}
	defer func() {
		if err :=recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("defer finished")
	}()
	fmt.Println(foods[5])
}

func end()  {
	fmt.Println("程序执行结束")
}
/**
recover 确保程序主流程的完整性，
权衡是不是要把错误吃掉，记录日志

有错误抛给上层 上层 上层 重要的事情说三遍
 */
func main()  {
	start()
	testFood()
	end()
}
