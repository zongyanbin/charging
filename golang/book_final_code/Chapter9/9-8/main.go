package main

import "fmt"

/**
类型分支  断言配合 switch
*/
func main() {
	var x interface{}
	x = 123
	switch x.(type) {
	case string:
		fmt.Println("这是string类型")
	case bool:
		fmt.Println("这个是bool")
	case int:
		fmt.Println("这个是int类型")
	}
}
