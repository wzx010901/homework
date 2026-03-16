package homework01

import (
	"fmt"
	"sort"
)

// 1. 只出现一次的数字
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func SingleNumber(nums []int) int {
	// TODO: implement
	result := 0
	// 遍历数组，依次异或每个元素
	for _, num := range nums {
		result ^= num
	}
	// 最终结果就是只出现一次的数字
	return result
}

// 2. 回文数
// 判断一个整数是否是回文数
func IsPalindrome(x int) bool {
	// TODO: implement
	// 根据回文特征：负数（含负号）、末尾为0且非0的数，直接返回false
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	//把x转成字符串判断长度 根据回文特征 字符串只能是奇数
	str := fmt.Sprintf("%d", x)
	strLen := len(str)
	for i := 0; i < strLen/2; i++ {
		if str[i] != str[strLen-i-1] {
			return false
		}
	}
	return true
}

// 3. 有效的括号
// 给定一个只包括 '(', ')', '{', '}', '[', ']' 的字符串，判断字符串是否有效
func IsValid(s string) bool {
	// TODO: implement
	// 特殊情况：字符串长度为奇数，直接返回false（括号无法成对）
	if len(s)%2 != 0 {
		return false
	}
	// 定义右括号到左括号的映射，方便快速匹配
	m := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	// 用模拟栈，存储左括号
	stack := []rune{}
	for _, char := range s {
		//  如果是右括号（存在于map的key中）
		if leftChar, ok := m[char]; ok {
			// 栈为空 或 栈顶左括号不匹配 无效
			if len(stack) == 0 || stack[len(stack)-1] != leftChar {
				return false
			}
			// 匹配成功，栈顶出栈
			stack = stack[:len(stack)-1]
		} else {
			// 如果是左括号，入栈
			stack = append(stack, char)
		}
	}

	// 4. 遍历结束后，栈必须为空（所有左括号都匹配完成）
	return len(stack) == 0
}

// 4. 最长公共前缀
// 查找字符串数组中的最长公共前缀
func LongestCommonPrefix(strs []string) string {
	// TODO: implement
	// 处理特殊情况：空数组
	if len(strs) == 0 {
		return ""
	}
	//以第一个字符串为前缀
	prefix := strs[0]
	//遍历前缀的每个字符
	for i := 0; i < len(prefix); i++ {
		// 取当前要对比的字符
		currentChar := prefix[i]
		// 检查其他所有字符串的第i个字符是否匹配
		for j := 0; j < len(strs); j++ {
			// 两种不匹配情况：
			// - 第j个字符串长度不足i（没有第i个字符）
			// - 第j个字符串的第i个字符与字符不同
			if i >= len(strs[j]) || strs[j][i] != currentChar {
				// 截取前i个字符作为最长公共前缀
				return prefix[:i]
			}
		}
	}

	// 所有字符都匹配，返回前缀
	return prefix
}

// 5. 加一
// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func PlusOne(digits []int) []int {
	// TODO: implement
	// 从最后一位（个位）向前遍历
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++ // 当前位加 1
		// 取余判断是否进位：余数不为 0 说明无进位，直接返回
		digits[i] %= 10
		if digits[i] != 0 {
			return digits
		}
	}

	// 走到这里说明所有位都是9（如 [9,9]），需要在头部补1
	// 新建，长度为 len(digits)+1，首元素为1，其余默认0
	result := make([]int, len(digits)+1)
	result[0] = 1
	return result
}

// 6. 删除有序数组中的重复项
// 给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
func RemoveDuplicates(nums []int) int {
	// TODO: implement
	// 边界条件：空数组直接返回0
	if len(nums) == 0 {
		return 0
	}
	//初始指向第一个元素
	slow := 0
	//从第二个元素开始遍历
	for fast := 1; fast < len(nums); fast++ {
		// 发现不同元素，慢指针后移并覆盖值
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		// 相同元素则继续前进，初始指向不动
	}
	// 新长度为慢指针索引+1
	return slow + 1
}

// 7. 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
func Merge(intervals [][]int) [][]int {
	// 边界条件：空区间直接返回空
	if len(intervals) == 0 {
		return nil
	}
	// 按区间起始值升序排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 初始化结果，先加入第一个区间
	result := [][]int{intervals[0]}
	// 遍历剩余区间，逐个合并
	for i := 1; i < len(intervals); i++ {
		// 取结果中最后一个区间（待合并的区间）
		last := result[len(result)-1]
		// 当前遍历的区间
		current := intervals[i]

		// 判断是否重叠：当前区间的起始值 <= 区间的结束值
		if current[0] <= last[1] {
			// 重叠则合并：更新区间的结束值为两者的最大值
			if last[1] < current[1] {
				last[1] = current[1]
			}
			// 因为是引用类型，直接修改last会同步到result中
			result[len(result)-1] = last
		} else {
			// 不重叠则直接加入结果
			result = append(result, current)
		}
	}

	return result
}

// 8. 两数之和
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func TwoSum(nums []int, target int) []int {
	// TODO: implement
	if len(nums) == 0 {
		return nil
	}
	// 定义map：key = 数组值，value = 对应索引
	numMap := make(map[int]int)
	// 遍历数组，同时记录已遍历元素的索引
	for idx, num := range nums {
		// 计算需要匹配的补数：target - 当前值
		complement := target - num
		complementIndex, exists := numMap[complement]
		// 检查补数是否已在哈希表中（存在则找到解）
		if exists {
			// 返回补数的索引和当前索引
			return []int{complementIndex, idx}
		}

		// 补数不存在，将当前值和索引存入map
		numMap[num] = idx
	}

	// 题目保证有解，此处仅为语法兜底
	return nil
}
