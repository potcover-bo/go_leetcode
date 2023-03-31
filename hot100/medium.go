package main

/*
*
两数相加
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	l1.Val += l2.Val

	if l1.Val >= 10 {
		l1.Next = addTwoNumbers(&ListNode{Val: 1}, l1.Next)
		l1.Val %= 10
	}

	l1.Next = addTwoNumbers(l1.Next, l2.Next)

	return l1
}

/*
*
46、全排列
*/
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	used := make([]bool, len(nums))
	path := make([]int, 0)

	var backtrack func(nums, path []int, used []bool)

	backtrack = func(nums, path []int, used []bool) {
		if len(path) == len(nums) {
			val := make([]int, len(nums))
			copy(val, path)
			res = append(res, val)
			return
		}

		for i := 0; i < len(nums); i++ {
			if !used[i] {
				path = append(path, nums[i])
				used[i] = true
				backtrack(nums, path, used)

				path = path[:len(path)-1]
				used[i] = false
			}
		}
	}

	backtrack(nums, path, used)

	return res
}

/*
*
78、子集
*/
func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)

	var backtrack func(start int)

	backtrack = func(start int) {

		res = append(res, append([]int(nil), path...))

		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])

			backtrack(i + 1)

			path = path[:len(path)-1]
		}
	}

	backtrack(0)

	return res
}

/*
*
22 括号生成
*/
func generateParenthesis(n int) []string {

	res := make([]string, 0)
	path := ""
	var backtrack func(int, int)
	backtrack = func(left int, right int) {
		if left < 0 || right < 0 {
			return
		}

		if right < left {
			return
		}

		if left == 0 && right == 0 {
			res = append(res, path)
			return
		}

		path += "("
		backtrack(left-1, right)
		path = path[:len(path)-1]

		path += ")"
		backtrack(left, right-1)
		path = path[:len(path)-1]
	}

	backtrack(n, n)

	return res
}

/*
*
538
给出二叉 搜索 树的根节点，该树的节点值各不相同，
请你将其转换为累加树（Greater Sum Tree），
使每个节点 node 的新值等于原树中大于或等于 node.val 的值之和。
*/
func convertBST(root *TreeNode) *TreeNode {
	preSum := 0

	var preOrder func(*TreeNode)

	preOrder = func(node *TreeNode) {
		if node != nil {
			preOrder(node.Right)
			node.Val += preSum
			preSum = node.Val
			preOrder(node.Left)

		}
	}

	preOrder(root)

	return root
}

/*
*
238 除自身以外数组的乘积
*/
func productExceptSelf(nums []int) []int {
	left := make([]int, len(nums))
	left[0] = 1
	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	right := 1

	for i := len(left) - 1; i >= 0; i-- {
		left[i] = right * left[i]
		right *= nums[i]
	}

	return left
}

/*
*75. 颜色分类
给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

必须在不使用库内置的 sort 函数的情况下解决这个问题
*/
func sortColors(nums []int) {
	left := -1
	right := len(nums)
	cur := 0

	var swap func(l, r int)
	swap = func(l, r int) {
		tmp := nums[l]
		nums[l] = nums[r]
		nums[r] = tmp
	}

	for cur < right {
		if nums[cur] == 0 {
			swap(left+1, cur)
			left += 1
			cur += 1
		} else if nums[cur] == 1 {
			cur += 1
		} else {
			swap(right-1, cur)
			right -= 1
		}
	}

}

/*
*
刷油漆算法
*/
func sortColors2(nums []int) {
	n0 := 0
	n1 := 0
	var tmp int
	for i := 0; i < len(nums); i++ {
		tmp = nums[i]
		nums[i] = 2

		if tmp < 2 {
			nums[n1] = 1
			n1 += 1
		}

		if tmp < 1 {
			nums[n0] = 0
			n0 += 1
		}
	}
}

/*
* 19. 删除链表的倒数第 N 个结点
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	dummy := &ListNode{Val: -1, Next: head}
	cur := dummy

	cnt := 0

	for head != nil {
		cnt += 1
		head = head.Next
	}

	cnt -= n
	for cnt > 0 {
		cur = cur.Next
		cnt -= 1
	}

	cur.Next = cur.Next.Next

	return dummy.Next
}

/*
* 739. 每日温度
 */
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))

	return res

}
