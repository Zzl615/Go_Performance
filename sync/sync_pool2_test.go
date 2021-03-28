//  BenchmarkBuffer-4                   498951              2014 ns/op           10240 B/op            1 allocs/op
//  -4表示4个CPU线程执行  op:每次执行  总共执行498951次       耗时4531纳秒         10240字节内存           分配1次对象

// go test -bench="Buffer" -benchmem sync_pool2_test.go
package main

import (
	"bytes"
	"sync"
	"testing"
)

var bufferPool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

var data = make([]byte, 10000)

func BenchmarkBufferWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := bufferPool.Get().(*bytes.Buffer)
		buf.Write(data)
		buf.Reset()
		bufferPool.Put(buf)
	}
}

func BenchmarkBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var buf bytes.Buffer
		buf.Write(data)
	}
}
