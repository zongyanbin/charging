package main

func main() {
	// chan接收和发送都是互斥的---有我没你，有你没我
	// 发送和接收都是原子性 ----足够小，小到不能再分割
	// 发送成功之前会被阻塞，接收也是如此----阻塞不能继续执行了
	// 缓冲通道 		channl nil
	// 无缓冲通道
}
