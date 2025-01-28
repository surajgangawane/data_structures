package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ch1 := make(chan int)
	ch2 := make(chan int)

	newCtx, cancel := context.WithCancel(ctx)

	go testing1(newCtx, ch1)
	go testing2(newCtx, ch2)

	select {
	case t1 := <-ch1:
		fmt.Println(t1)
	case t2 := <-ch2:
		fmt.Println(t2)
		fmt.Println("Cancelling context")
		cancel()
	}
}

func testing1(ctx context.Context, ch chan int) {
	select {
	case <-time.After(20 * time.Second):
		ch <- 1
	case <-ctx.Done():
		fmt.Println("Goroutine was cancelled")
	}
}

func testing2(ctx context.Context, ch chan int) {
	time.Sleep(1 * time.Second)
	ch <- -1
}
