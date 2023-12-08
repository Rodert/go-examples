package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

//场景1: 简单并发任务

func RunScenario1() {
	count := 10
	var wg sync.WaitGroup

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			doSomething(index)
		}(i)
	}

	wg.Wait()
}

// 场景2: 按时间来持续并发

func RunScenario2() {
	timeout := time.Now().Add(time.Second * time.Duration(10))
	n := runtime.NumCPU()

	waitForAll := make(chan struct{})
	done := make(chan struct{})
	concurrentCount := make(chan struct{}, n)

	for i := 0; i < n; i++ {
		concurrentCount <- struct{}{}
	}

	go func() {
		for time.Now().Before(timeout) {
			<-done
			concurrentCount <- struct{}{}
		}

		waitForAll <- struct{}{}
	}()

	go func() {
		for {
			<-concurrentCount
			go func() {
				doSomething(rand.Intn(n))
				done <- struct{}{}
			}()
		}
	}()

	<-waitForAll
}

// 场景3：以 worker pool 方式 并发做事/发送请求

func RunScenario3() {
	numOfConcurrency := runtime.NumCPU()
	taskTool := 10
	jobs := make(chan int, taskTool)
	results := make(chan int, taskTool)
	var wg sync.WaitGroup

	// workExample
	workExampleFunc := func(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
		defer wg.Done()
		for job := range jobs {
			res := job * 2
			fmt.Printf("Worker %d do things, produce result %d \n", id, res)
			time.Sleep(time.Millisecond * time.Duration(100))
			results <- res
		}
	}

	for i := 0; i < numOfConcurrency; i++ {
		wg.Add(1)
		go workExampleFunc(i, jobs, results, &wg)
	}
	totalTasks := 100

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < totalTasks; i++ {
			n := <-results
			fmt.Printf("Got results %d \n", n)
		}
		close(results)
	}()

	for i := 0; i < totalTasks; i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}

// 场景4: 等待异步任务执行结果(goroutine+select+channel)

func RunScenario4() {
	sth := make(chan string)
	result := make(chan string)
	go func() {
		id := rand.Intn(100)
		for {
			sth <- doSomething(id)
		}
	}()
	go func() {
		for {
			result <- takeSomthing(<-sth)
		}
	}()

	select {
	case c := <-result:
		fmt.Printf("Got result %s ", c)
	case <-time.After(time.Duration(30 * time.Second)):
		fmt.Errorf("指定时间内都没有得到结果")
	}
}

var doUploadMock = func() bool {
	time.Sleep(time.Millisecond * time.Duration(100))
	n := rand.Intn(100)
	if n > 50 {
		return true
	} else {
		return false
	}
}

// 场景5: 定时的反馈结果(Ticker) 测试上传接口的性能，要实时给出指标: 吞吐率，成功率等

func RunScenario5() {
	totalSize := int64(0)
	totalCount := int64(0)
	totalErr := int64(0)

	concurrencyCount := runtime.NumCPU()
	stop := make(chan struct{})
	fileSizeExample := int64(10)

	timeout := 10 // seconds to stop

	go func() {
		for i := 0; i < concurrencyCount; i++ {
			go func(index int) {
				for {
					select {
					case <-stop:
						return
					default:
						break
					}

					res := doUploadMock()
					if res {
						atomic.AddInt64(&totalCount, 1)
						atomic.AddInt64(&totalSize, fileSizeExample)
					} else {
						atomic.AddInt64(&totalErr, 1)
					}
				}
			}(i)
		}
	}()

	t := time.NewTicker(time.Second)
	index := 0
	for {
		select {
		case <-t.C:
			index++
			tmpCount := atomic.LoadInt64(&totalCount)
			tmpSize := atomic.LoadInt64(&totalSize)
			tmpErr := atomic.LoadInt64(&totalErr)
			fmt.Printf("吞吐率: %d，成功率: %d \n", tmpSize/int64(index), tmpCount*100/(tmpCount+tmpErr))
			if index > timeout {
				t.Stop()
				close(stop)
				return
			}
		}

	}
}
