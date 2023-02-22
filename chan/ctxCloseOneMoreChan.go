package main

import (
	"context"
	"fmt"
	"time"
)

// chan的关闭会都通知到，但收发两边只能同时有一个人

func main2() {
	ctx, cancle := context.WithCancel(context.Background())
	ch := make(chan string)
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			break
		}
		fmt.Println("ctx我运行完了")
	}(ctx)
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			break
		}
		fmt.Println("ctx我运行完了")
	}(ctx)
	go func(ch chan string) {
		select {
		case str := <-ch:
			fmt.Println("????", str)
		}
		fmt.Println("ch我运行完了")
	}(ch)
	go func(ch chan string) {
		select {
		case str := <-ch:
			fmt.Println("????", str)
		}
		fmt.Println("ch我运行完了")
	}(ch)
	time.Sleep(time.Second * 2)
	cancle()
	// ch <- "123"
	close(ch)
	time.Sleep(time.Second * 3)
}
