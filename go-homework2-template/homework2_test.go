package homework01

import (
	"sync"
	"testing"
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
	defer recordResult(t, "SingleNumber")
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
				t.Errorf("SingleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleElements(t *testing.T) {
	defer recordResult(t, "SingleNumber")
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"Example 1", []int{1, 2, 3, 4, 5}, []int{2, 3, 4, 5, 10}},
		{"Example 2", []int{1, 3, 5, 7, 9}, []int{2, 6, 10, 14, 18}},
		{"Example 3", []int{2, 4, 6, 8, 10}, []int{4, 8, 12, 16, 20}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DoubleElements(&tt.nums); IntSliceEqual(got, tt.want) {
				t.Errorf("SingleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func IntSliceEqual(a, b []int) bool {
	// 第一步：判断长度是否相等（长度不同直接不相等）
	if len(a) != len(b) {
		return false
	}

	// 第二步：长度为0时（包括nil和空切片），视为相等
	if len(a) == 0 {
		return true
	}

	// 第三步：逐元素对比（长度相同且非空）
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	// 所有条件满足，返回相等
	return true
}
