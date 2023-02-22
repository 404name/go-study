package main

import (
	"context"
	"fmt"
	"time"
)

func main1() {
	ctx, cancle := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			break
		default:
		}
		fmt.Println("有default可以直接出来")
	}(ctx)
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			break
		}
		fmt.Println("没default只能等结束")
	}(ctx)

	time.Sleep(time.Second * 3)
	cancle()
	time.Sleep(time.Second * 3)
}
