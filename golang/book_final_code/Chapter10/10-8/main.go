package main
/**
缓存channel
 */
func main() {
	// c:=make(chan int,7)
	//chan发送操作   队列尾部插入元素
	//chan接收操作	队列头部移除一个元素
	//chan 满了 		发送方会阻塞所在goroutine  ----->直到另一个goroutine对此chan接收留出相应空间
	// 	如果 chan 是空则接收的goroutine被阻塞-------->直到另一个goroutine对此chan发送数据

	// 无缓冲通道  是同步的
	// 有缓存通道	 是异步的
	}
