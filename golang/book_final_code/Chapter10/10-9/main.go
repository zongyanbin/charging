package main

import (
	"fmt"
	"time"
)

/**
select 检查 channel
*/
func GetFood1(c chan string) {
	time.Sleep(3 * time.Second)
	close(c)
}

func GetFood2(c chan string) {
	time.Sleep(3 * time.Second)
	c <- "清蒸鱼好了"
}

func GetFood3(c chan string) {
	time.Sleep(2 * time.Second)
	c <- "烤生蚝好了。"

}

var ch1 chan string
var ch2 chan string
var ch3 chan string
var foods = []string{"红烧肉", "清蒸鱼", "烤生蚝"}
var chs = []chan string{ch1, ch2, ch3}

func getFood(i int) string {
	fmt.Printf("Foods[%d]\n", i)
	return foods[i]
}
func getChan(i int) chan string {
	fmt.Printf("chs[%d]\n", i)
	return chs[i]
}
func main() {
	// 创建通道
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	go GetFood1(c1)
	go GetFood2(c2)
	go GetFood3(c3)

	// select 先到先得
	select {
	case r:=<-c1:
		fmt.Println(r)
	case r:=<-c2:
		fmt.Println(r)
	case r:=<-c3:
		fmt.Println(r)
	default:
		fmt.Println("菜还没有好。")
	}

	select {
	case r:=<-c1:
		fmt.Println(r)
	case r:=<-c2:
		fmt.Println(r)
	case r:=<-c3:
		fmt.Println(r)
		break
	}

	// 求值 从上往下，从左往右
	select {
	case getChan(0)<-getFood(2):
		fmt.Println("1th case is selected.")
	case getChan(1)<-getFood(1):
		fmt.Println("2th case is selected.")
	default:
		fmt.Println("default!.")

	}
}
