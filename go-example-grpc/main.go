package main

import (
    "fmt"
    "time"
)

func main() {
    // 创建一个每隔1秒触发一次的定时器
    ticker := time.NewTicker(1 * time.Second)

    // 使用匿名的 goroutine 来处理定时触发的事件
    go func() {
        for {
            // 通过 <-ticker.C 从通道中接收时间信号，此处会每隔1秒触发一次
            // 可以在这里执行你想要的操作
            fmt.Println("Tick at", <-ticker.C)
			fmt.Println("123")
        }
    }()

    // 主函数继续运行，这里只是演示，让主程序不会立即退出
    // 在实际使用中，你可能需要让程序运行一段时间或者通过某种方式等待goroutine执行完毕
    time.Sleep(5 * time.Second)

    // 停止定时器
    ticker.Stop()
    fmt.Println("Ticker stopped")
}