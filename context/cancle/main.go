package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

/*
*

1. 不取消则自动超时 cancel()
2. 用管道通知 done <- struct{}{}
3. 都不开定时器结束

*
*/
func mockErr(ctx context.Context) error {
	time.Sleep(1000 * time.Millisecond)
	return errors.New("failed")
}

func main() {
	// 新建一个上下文
	ctx, _ := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	done := make(chan struct{}, 1)
	// 在不同的goroutine中运行operation2
	err := mockErr(ctx)
	// 如果这个操作返回错误，取消所有使用相同上下文的操作
	if err != nil {
		//cancel() // 1.主动取消则自动超时
		//done <- struct{}{} // 2. 用管道通知
		// 3.都不开定时器结束
	}
	mockHttp(ctx, done) //会阻塞在这里

	println("服务结束")
}

func mockHttp(ctx context.Context, done chan struct{}) {
	// 我们使用在前面HTTP服务器例子里使用过的类型模式
	fmt.Println("http请求")
	go func() {
		for {
			fmt.Println("处理http请求")
			time.Sleep(500 * time.Millisecond)
		}
	}()
	select { // 会阻塞在这里，三个case都是阻塞的，谁先有消息执行一次谁就退出
	case <-time.After(2000 * time.Millisecond):
		fmt.Println("计时器中断")
		return
	case <-done:
		fmt.Println("被管道中断")
		return
	case <-ctx.Done():
		fmt.Println("被ctx超时/主动中断")
		return
	}
}
