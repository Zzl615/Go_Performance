package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

var wg sync.WaitGroup

func doThing(i int, ch chan struct{}) {
	defer wg.Done()
	fmt.Println(i)
	<-ch
	// panic: too many concurrent operations on a single file or socket (max 1048575)
	// 对单个 file/socket 的并发操作个数超过了系统上限
	// fmt.Printf 函数引起的，fmt.Printf 将格式化后的字符串打印到屏幕，即标准输出。在 linux 系统中，标准输出也可以视为文件，内核(kernel)利用文件描述符(file descriptor)来访问文件，标准输出的文件描述符为 1，错误输出文件描述符为 2，标准输入的文件描述符为 0。
	// 简而言之，系统的资源被耗尽
	time.Sleep(time.Second)
}

// math.MaxInt32个协程的并发。乱序输出 1 -> 2^31 个数字。

func main() {
	ch := make(chan struct{}, 3)
	for i := 0; i < math.MaxInt32; i++ {
		ch <- struct{}{}
		wg.Add(1)
		go doThing(i, ch)
	}
	wg.Wait()
}
