package main

/*
*
239. 滑动窗口最大值
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回 滑动窗口中的最大值 。
*/
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) < 2 {
		return nums
	}
	res := make([]int, len(nums)-k+1)
	queue := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		// 维护队列
		for len(queue) > 0 && nums[queue[len(queue)-1]] < nums[i] {
			queue = queue[:len(queue)-1]
		}

		// 当前元素加入队列
		queue = append(queue, i)

		// 判断队首的元素是否合法
		if queue[0] <= i-k {
			queue = queue[1:]
		}

		// 窗口大小是否合法
		if i+1 >= k {
			res[i-k+1] = nums[queue[0]]
		}

	}

	return res
}
