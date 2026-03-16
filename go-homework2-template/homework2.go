package homework01

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

// 题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
func AddTen(numPtr *int) int {
	// 检查指针是否为nil（避免空指针panic，增强代码健壮性）
	if numPtr == nil {
		fmt.Println("错误：传入的指针为nil，无法修改值")
		return 0
	}
	// *numPtr 表示解引用指针，访问指针指向的内存地址中的值
	*numPtr += 10 // 等价于 *numPtr = *numPtr + 10
	return *numPtr
}

// 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
func DoubleElements(slicePtr *[]int) []int {
	// 1. 检查指针是否为nil（避免空指针panic）
	if slicePtr == nil {
		fmt.Println("错误：传入的切片指针为nil")
		return nil
	}

	//  解引用指针，获取原始切片（*slicePtr 表示指针指向的切片）
	slice := *slicePtr

	// 3. 检查切片是否为空（空切片无需处理）
	if len(slice) == 0 {
		fmt.Println("提示：切片为空，无需修改")
		return nil
	}

	// 遍历切片，每个元素乘以2
	for i := range slice {
		slice[i] *= 2
	}

	return slice
	// 注意：由于切片是引用类型，修改 slice 会同步到原始切片，无需重新赋值
}

// 编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 考察点 ： go 关键字的使用、协程的并发执行。
func PrintOddEven() {
	// 定义WaitGroup，设置需要等待的协程数量为2
	var wg sync.WaitGroup
	wg.Add(2)

	// 使用go关键字启动协程，传入WaitGroup指针
	go PrintOdd(&wg)
	go PrintEven(&wg)

	// 等待所有协程执行完成
	wg.Wait()
	fmt.Println("所有协程执行完毕")
}

func PrintOdd(wg *sync.WaitGroup) {
	// 函数执行完毕后，通知WaitGroup完成一个任务
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		fmt.Printf("奇数协程：%d\n", i)
	}
}

// printEven 打印2到10的偶数
func PrintEven(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		fmt.Printf("偶数协程：%d\n", i)
	}
}

// 设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。
// Task 定义任务类型：接收任务名，返回错误（模拟任务执行可能出错）
type Task func(taskName string) error

// TaskResult 存储单个任务的执行结果
type TaskResult struct {
	TaskName    string        // 任务名
	ElapsedTime time.Duration // 执行耗时
	Err         error         // 执行错误（无错误则为nil）
}

// Scheduler 任务调度器结构体
type Scheduler struct {
	tasks     map[string]Task // 待执行的任务（key:任务名，value:任务函数）
	results   []TaskResult    // 所有任务的执行结果
	mutex     sync.Mutex      // 保护results的并发写入
	waitGroup sync.WaitGroup  // 等待所有任务完成
}

// NewScheduler 创建调度器实例
func NewScheduler() *Scheduler {
	return &Scheduler{
		tasks:   make(map[string]Task),
		results: make([]TaskResult, 0),
	}
}

// PrintResults 打印所有任务的执行结果
func (s *Scheduler) PrintResults() {
	fmt.Println("===== 任务执行结果 =====")
	for _, res := range s.results {
		if res.Err != nil {
			fmt.Printf("任务 [%s] 执行失败：%v\n", res.TaskName, res.Err)
		} else {
			fmt.Printf("任务 [%s] 执行成功，耗时：%v\n", res.TaskName, res.ElapsedTime)
		}
	}
}

// AddTask 向调度器添加任务
func (s *Scheduler) AddTask(taskName string, task Task) {
	s.tasks[taskName] = task
}

// Run 执行所有任务并统计时间
func (s *Scheduler) Run() {
	// 为每个任务添加等待计数
	s.waitGroup.Add(len(s.tasks))
	// 遍历所有任务，启动协程执行
	for name, task := range s.tasks {
		// 注意：循环变量在协程中会复用，需捕获当前值
		taskName := name
		taskFunc := task
		go func() {
			defer s.waitGroup.Done() // 任务完成后减少计数
			// 记录任务开始时间
			start := time.Now()
			// 执行任务
			err := taskFunc(taskName)
			// 计算执行耗时
			elapsed := time.Since(start)
			// 安全写入结果（加锁避免并发冲突）
			s.mutex.Lock()
			s.results = append(s.results, TaskResult{
				TaskName:    taskName,
				ElapsedTime: elapsed,
				Err:         err,
			})
			s.mutex.Unlock()
		}()
	}
	// 等待所有任务执行完成
	s.waitGroup.Wait()
}

//定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
//在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
//考察点 ：接口的定义与实现、面向对象编程风格。

// 定义接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle 实现Shape接口
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle 实现Shape接口
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func PrintShapeInfo(s Shape, name string) string {
	fmt.Printf("%s: Area=%.2f, Perimeter=%.2f\n",
		name, s.Area(), s.Perimeter())
	return fmt.Sprintf("%s: Area=%.2f, Perimeter=%.2f\n",
		name, s.Area(), s.Perimeter())
}

//使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，
//再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
//为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
//考察点 ：组合的使用、方法接收者。

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (employee *Employee) PrintInfo() string {
	// 直接访问嵌入的Person字段（等价于 e.Person.Name / e.Person.Age）
	fmt.Println("===== 员工信息 =====")
	fmt.Printf("员工ID：%d\n", employee.EmployeeID)
	fmt.Printf("姓名：%s\n", employee.Name)
	fmt.Printf("年龄：%d\n", employee.Age)
	return fmt.Sprintf("员工ID：%d\n姓名：%s\n年龄：%d\n", employee.EmployeeID, employee.Name, employee.Age)
}

// 编写一个程序，使用通道实现两个协程之间的通信。
// 一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
// 考察点 ：通道的基本使用、协程间通信。
func ChannelDemo() {
	// 无缓冲channel
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	//生产者协程：生成1-10的整数并发送到通道
	go func() {
		defer wg.Done() // 协程结束时通知WaitGroup
		for i := 1; i <= 10; i++ {
			fmt.Printf("生产者发送：%d\n", i)
			ch <- i // 将整数发送到通道
		}
		close(ch) // 发送完成后关闭通道，告知接收方无数据可接收
	}()

	//消费者协程：从通道接收整数并打印
	go func() {
		defer wg.Done() // 协程结束时通知WaitGroup
		// 循环接收通道数据，通道关闭后循环自动退出
		for num := range ch {
			fmt.Printf("消费者接收并打印：%d\n", num)
		}
		fmt.Println("通道已关闭，消费者停止接收")
	}()

	// 4. 主协程等待所有子协程完成
	wg.Wait()
	fmt.Println("所有协程执行完毕")
}

// 实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
// 考察点 ：通道的缓冲机制。
func BufferedChannelDemo() {
	ch := make(chan int, 5)
	var wg sync.WaitGroup
	wg.Add(2)
	//生产者协程：生成1-100的整数并发送到缓冲通道
	go func() {
		defer func() {
			close(ch) // 发送完成后关闭通道，避免消费者阻塞
			wg.Done() // 通知等待组：生产者完成
		}()

		for i := 1; i <= 100; i++ {
			ch <- i // 发送数据到缓冲通道
			// 可选：打印发送状态，直观展示缓冲特性
			fmt.Printf("生产者发送：%d | 通道当前元素数：%d\n", i, len(ch))
		}
		fmt.Println("生产者已发送所有100个整数，关闭通道")
	}()

	//费者协程：从缓冲通道接收数据并打印
	go func() {
		defer wg.Done() // 通知等待组：消费者完成

		// 循环接收通道数据，通道关闭后自动退出循环
		count := 0 // 统计接收的整数数量
		for num := range ch {
			fmt.Printf("消费者接收：%d\n", num)
			count++
		}
		fmt.Printf("消费者已接收所有数据，总计：%d个\n", count)
	}()

	// 主协程等待所有子协程完成
	wg.Wait()
	fmt.Println("所有协程执行完毕")
}

// 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全。

// Counter 定义带互斥锁的计数器结构体，封装共享资源和锁
type Counter struct {
	mu    sync.Mutex // 互斥锁，保护count的并发修改
	count int        // 共享的计数器
}

// Incr 计数器递增方法，加锁保证并发安全
func (c *Counter) Incr() {
	c.mu.Lock()         // 加锁：同一时间只有一个协程能进入临界区
	defer c.mu.Unlock() // 解锁：函数退出时执行，避免忘记解锁导致死锁
	c.count++           // 临界区：修改共享资源
}

// GetCount 获取计数器当前值，加锁保证读取也安全
func (c *Counter) GetCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}
func MutexDemo() {
	// 初始化计数器
	counter := &Counter{}
	var wg sync.WaitGroup

	// 配置等待组：等待10个协程
	wg.Add(10)

	// 启动10个协程，每个协程执行1000次递增
	for i := 0; i < 10; i++ {
		go func(coroutineID int) {
			defer wg.Done() // 协程结束时通知等待组
			// 每个协程递增1000次
			for j := 0; j < 1000; j++ {
				counter.Incr()
			}
			fmt.Printf("协程%d：完成1000次递增\n", coroutineID)
			fmt.Printf("协程%d：完成1000次递增,结果是%d\n", coroutineID, counter.GetCount())
		}(i) // 传入协程ID，方便打印日志
	}

	// 等待所有协程执行完毕
	wg.Wait()

	// 输出最终计数器值（预期为10*1000=10000）
	fmt.Printf("\n最终计数器值：%d\n", counter.GetCount())

}

// 使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
func AtomicDemo() {
	var count int32
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(coroutineID int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt32(&count, 1)
			}
			fmt.Printf("协程%d：完成1000次递增\n", coroutineID)
			fmt.Printf("协程%d：完成1000次递增,结果是%d\n", coroutineID, atomic.LoadInt32(&count))
		}(i)
	}
	wg.Wait()
	fmt.Printf("\n最终计数器值：%d\n", atomic.LoadInt32(&count))
}
