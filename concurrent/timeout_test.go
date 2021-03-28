package main

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func doBadthing(done chan bool) {
	time.Sleep(time.Second)
	done <- true
}

// 超时处理：timeout(doBadthing)
func timeout(f func(chan bool)) error {
	done := make(chan bool)
	go f(done)
	select {
	case <-done:
		fmt.Println("done")
		return nil
	case <-time.After(time.Millisecond):
		return fmt.Errorf("timeout")
	}
}

func test01(t *testing.T, f func(chan bool)) {
	t.Helper()
	for i := 0; i < 1000; i++ {
		timeout(f)
	}
	time.Sleep(time.Second * 2)
	t.Log(runtime.NumGoroutine())

}

func benchmarkBadTimeout(b *testing.B, f func(chan bool)) {
	for n := 0; n < b.N; n++ {
		timeout(f)
	}
	time.Sleep(time.Second * 2)
	b.Log(runtime.NumGoroutine())
}

// go test -run ^TestBadTimeout$ . -v
func TestBadTimeout(t *testing.T) { test01(t, doBadthing) }

// go test -bench='BenchmarkBadTimeou' -benchmem .
func BenchmarkBadTimeout(b *testing.B) { benchmarkBadTimeout(b, doBadthing)