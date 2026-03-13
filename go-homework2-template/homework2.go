package homework01

import "fmt"

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
