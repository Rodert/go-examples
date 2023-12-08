package main

import (
	"context"
	"fmt"
	"time"
)

func handleRequest(ctx context.Context) {
	select {
	case <-time.After(time.Second * 2):
		fmt.Println("request succeeded")
	case <-ctx.Done():
		fmt.Println("request canceled or timed out")
	}
}

func main() {
	// runtime.GOMAXPROCS(0)

	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(1 * time.Second)

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()

	// go handleRequest(ctx)

	// time.Sleep(time.Second * 3)
	// fmt.Println("main goroutine exit")
}
