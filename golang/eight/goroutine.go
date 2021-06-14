/**
多线程 多进程
cpu 轮询调用
A B C 同时执行->时间片超过就切换进程

多进程/多线程解决了阻塞问题

又面临新的问题
线程1<->线程2<->线程1
cpu浪费时间成本=切换成本

进程/线程的数量越多
切换成本就越大
也就越浪费
cpu 100% 有可能  60% 执行程序
				40% 在切换中
提高cpu利用率 软件 系统架构需要做的事情

多线程随着同步竞争（如 锁、竞争资源冲突等）
开发设计 变得越来越复杂

			协程（co-routine）用户空间			M
					绑定
cup 视野--->		thread内核空间

				   一个线程					N

M:N 优化写成调度器

协程（co-routine）-> (内存：几kb) --->可以大量
					Goroutine
					灵活调度------>可以长切换

GMP
	G	goroutine 协程
	p	processor 处理器
	m	thread 线程
*/
package main

import (
	"fmt"
	"time"
)

// 子goroutine
func newTask()  {
	i :=0
	for  {
		i++
		fmt.Printf("new Groutine :i = %d\n",i)
		time.Sleep(1 * time.Second)

	}
}

// 主goroutine
func main()  {
	go newTask()
	//i :=0
	//for  {
	//	i++
	//	fmt.Printf("main goroutine: i = %d\n",i)
	//	time.Sleep(1 * time.Second)
	//}
}