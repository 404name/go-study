package main

import "time"

func main() {
	c := make(chan string)
	//无缓冲通道，必须两边同时接受，不然都会卡主
	// 设置成make(chan string,1) 就是有缓冲通道就长度为1后再阻塞
	go func() {
		println("接受阻塞住了")
		println(<-c)
		println("接受通了")
	}()

	go func() { //注释了就会卡主发送那边
		println("发送塞住了")
		c <- "哈哈哈哈哈"
		println("发送通了")
	}()
	time.Sleep(2 * time.Second)
	println("程序结束")
}
