package main

import (
	"log"
	"net/http"
)

const (
	MaxWorker = 100 // 随便设置值
	MaxQueue  = 200 // 随便设置值
)

// 一个可以发送工作请求的缓冲 channel
var JobQueue chan Job

func init() {
	JobQueue = make(chan Job, MaxQueue)
}

type Payload struct{}

type Job struct {
	PayLoad Payload
}

// 接收请求，把任务筛入 JobQueue。
func payloadHandler(w http.ResponseWriter, r *http.Request) {
	work := Job{PayLoad: Payload{}}
	JobQueue <- work
	_, _ = w.Write([]byte("操作成功"))
}

func main() {
	// 通过调度器创建worker，监听来自 JobQueue的任务
	d := NewDispatcher()
	d.Run()
	http.HandleFunc("/payload", payloadHandler)
	log.Fatal(http.ListenAndServe(":8099", nil))
}
