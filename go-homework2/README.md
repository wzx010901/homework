# Go Homework 01 Template

## 使用说明

1. 编辑 `homework.go`，完成各个函数的实现。
2. 在本地运行 `go test -v` 验证代码。
3. 提交代码到你的分支。GitHub Actions 会自动运行测试。

## 题目列表

1. Add Ten (值增加10)
2. PrintOdd Even (两个协程，一个打印奇数，一个打印偶数)
3. Run (接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。)
4. PrintShapeInfo ( Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口)
5. ChannelDemo (编写一个程序，使用通道实现两个协程之间的通信。)
6. BufferedChannelDemo (实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。)
7. MutexDemo (编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。)
8. AtomicDemo (使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。)

