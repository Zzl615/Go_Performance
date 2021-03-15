# Go_Performance
Performance exploration of Golang

## 使用 -benchmem 参数看到内存分配
go test -bench='Generate' -benchmem .

## 使用 -count 参数可以用来设置 benchmark 的轮数

## 使用 -benchtime 可以设置的执行时间，可以设置具体的执行次数
- 执行30次，可以用 -benchtime=30x
- 执行5s，可以用-benchtime 指定为 5s
