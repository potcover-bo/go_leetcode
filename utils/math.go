package utils

// Max 返回x和y的最大值
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Swap(x, y int, nums []int) {
	temp := nums[x]
	nums[x] = nums[y]
	nums[y] = temp
}
