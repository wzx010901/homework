package homework01

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
	"time"
)

var (
	failedQuestions []string
	totalQuestions  int
	mu              sync.Mutex
)

func recordResult(t *testing.T, name string) {
	mu.Lock()
	defer mu.Unlock()
	totalQuestions++
	if t.Failed() {
		failedQuestions = append(failedQuestions, name)
	}
}

func TestAddTen(t *testing.T) {
	defer recordResult(t, "AddTen")
	tests := []struct {
		name string
		nums int
		want int
	}{
		{"Example 1", 5, 15},
		{"Example 2", 6, 16},
		{"Example 3", 7, 17},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddTen(&tt.nums); got != tt.want {
				t.Errorf("AddTen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleElements(t *testing.T) {
	defer recordResult(t, "DoubleElements")
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"Example 1", []int{1, 2, 3, 4, 5}, []int{2, 4, 6, 8, 10}},
		{"Example 2", []int{1, 3, 5, 7, 9}, []int{2, 6, 10, 14, 18}},
		{"Example 3", []int{2, 4, 6, 8, 10}, []int{4, 8, 12, 16, 20}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DoubleElements(&tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DoubleElements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintOddEven(t *testing.T) {
	defer recordResult(t, "PrintOddEven")
	PrintOddEven()
}

func TestTaskScheduler(t *testing.T) {
	defer recordResult(t, "TaskScheduler")
	// 1. 创建调度器
	scheduler := NewScheduler()

	// 2. 添加任务
	scheduler.AddTask("测试任务1", mockTask(200*time.Millisecond))
	scheduler.AddTask("测试任务2", mockTask(150*time.Millisecond))
	// 3. 执行所有任务
	scheduler.Run()
	// 4. 打印执行结果
	scheduler.PrintResults()
}

// 示例任务：模拟耗时任务（休眠指定时间）
func mockTask(sleepTime time.Duration) Task {
	return func(taskName string) error {
		fmt.Printf("开始执行任务：%s\n", taskName)
		time.Sleep(sleepTime) // 模拟任务耗时
		fmt.Printf("完成执行任务：%s\n", taskName)
		return nil
	}
}

func TestShape(t *testing.T) {
	defer recordResult(t, "Shape")
	tests := []struct {
		name      string
		shape     Shape
		shapeName string
		want      string
	}{
		{"Example 1", Rectangle{Width: 10, Height: 5}, "矩形", "矩形: Area=50.00, Perimeter=30.00"},
		{"Example 2", Circle{Radius: 5}, "圆形", "Area=78.54, Perimeter=31.42"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrintShapeInfo(tt.shape, tt.shapeName); got == tt.want {
				t.Errorf("Shape() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestEmployee(t *testing.T) {
	employee := Employee{EmployeeID: 1, Person: Person{
		Name: "李四",
		Age:  28,
	}}
	employee.PrintInfo()

}

func TestChannelDemo(t *testing.T) {
	ChannelDemo()

}

func TestBufferedChannelDemo(t *testing.T) {
	BufferedChannelDemo()

}

func TestMutexDemo(t *testing.T) {
	MutexDemo()
}

func TestAtomicDemo(t *testing.T) {
	AtomicDemo()
}
