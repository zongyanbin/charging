package main

import "fmt"

type Hi func(num string) string

func SayHell0(num string, hi Hi)  {
	result:=hi(num)
	fmt.Println(result)
}

func main()  {
	target :="东北食客"
	if target =="东北适合" {
		SayHell0("5", func(num string) string {
			return num +"位兄弟，欢迎光临"
		})
	}else {
		SayHell0("3", func(num string) string {
			return num + "位客人，欢迎光临"
		})
	}
}

