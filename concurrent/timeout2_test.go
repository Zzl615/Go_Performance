package main

// 超时时间为 1 ms，而 doBadthing 需要 1s 才能结束运行。
// 因此 timeout(doBadthing) 一定会触发超时。
// timeout(doBadthing) 调用了 1000 次，理论上会启动 1000 个子协程。
// 利用 runtime.NumGoroutine() 打印当前程序的协程个数。
// 因为 doBadthing 执行时间为 1s，因此打印协程个数前，等待 2s，确保函数执行完毕。

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

// Problem: 协程没正常退出
// 有缓冲区的 channel
func timeoutWithBuffer(f func(chan bool)) error {
	done := make(chan bool, 1)
	go f(done)
	select {
	case <-done:
		fmt.Println("done")
		return nil
	case <-time.After(time.Millisecond):
		return fmt.Errorf("timeout")
	}
}

// go test -run='TestBufferTimeout' . -v
func TestBufferTimeout(t *testing.T) {
	for i := 0; i < 1000; i++ {
		timeoutWithBuffer(doBadthing)
	}
	time.Sleep(time.Second * 2)
	t.Log(runtime.NumGoroutine())
}

// 处理完，select 尝试发送
func doGoodthing(done chan bool) {
	time.Sleep(time.Second)
	select {
	case done <- true:
	default:
		return
	}
}

// t, beforeClosed := <-taskCh 判断 channel 是否已经关闭，beforeClosed 为 false 表示信道已被关闭。若关闭，则不再阻塞等待，直接返回，对应的协程随之退出。
// sendTasks 函数中，任务发送结束之后，使用 close(taskCh) 将 channel taskCh 关闭
func doCheckClose(taskCh chan int) {
	for {
		select {
		case t, beforeClosed := <-taskCh:
			if !beforeClosed {
				fmt.Println("taskCh has been closed")
				return
			}
			time.Sleep(time.Millisecond)
			fmt.Printf("task %d is done\n", t)
		}
	}
}

func sendTasksCheckClose() {
	taskCh := make(chan int, 10)
	go doCheckClose(taskCh)
	for i := 0; i < 1000; i++ {
		taskCh <- i
	}
	close(taskCh)
}

func TestDoCheckClose(t *testing.T) {
	t.Log(runtime.NumGoroutine())
	sendTasksCheckClose()
	time.Sleep(time.Second)
	runtime.GC()
	t.Log(runtime.NumGoroutine())
}

func TestGoodTimeout(t *testing.T) { test01(t, doGoodthing) }
