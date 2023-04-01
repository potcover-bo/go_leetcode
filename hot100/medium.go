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
		// 生成新的节点
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
		// 终止条件
		if len(path) == len(nums) {
			//val := make([]int, len(nums))
			//copy(val, path)
			//res = append(res, val)
			res = append(res, append([]int(nil), path...))
			return
		}

		for i := 0; i < len(nums); i++ {
			// 当前位置的元素使用过
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

		// 从当前哪个位置开始
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
		// 终止条件 没有可以添加的括号
		if left < 0 || right < 0 {
			return
		}

		//不合法的括号
		if right < left {
			return
		}

		//合法的一个结果
		if left == 0 && right == 0 {
			res = append(res, path)
			return
		}

		//添加左括号
		path += "("
		backtrack(left-1, right)
		path = path[:len(path)-1]

		//添加右括号
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

	// 先序遍历的逆序
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
	//left[i] 表示[0, i-1]的乘积
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
	// 三路快拍的partition 荷兰国旗问题
	// nums[0, left] = 0
	// nums[left+1, cur] = 1
	// nums[right, len(nums) - 1] = 2
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

单调栈解法
*/
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := make([]int, 0)
	for i := 0; i < len(temperatures); i++ {
		if len(stack) == 0 {
			stack = append(stack, i)
		} else {
			for len(stack) != 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
				curIndex := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				res[curIndex] = i - curIndex

			}
			stack = append(stack, i)
		}
	}

	return res

}

/*
* 114. 二叉树展开为链表
 */
func flatten(root *TreeNode) {
	for root != nil {
		// 左子树为空的时候 直接指向右子树
		if root.Left == nil {
			root = root.Right
		} else {
			// 找到左子树最右边的节点
			pre := root.Left
			for pre.Right != nil {
				pre = pre.Right
			}

			// 左子树最右边的节点指向root的right
			pre.Right = root.Right

			// root的右节点指向root的左节点
			root.Right = root.Left

			// root的左节点置空
			root.Left = nil

			//root指向他的右节点
			root = root.Right
		}
	}
}

/*
*
236、二叉树的最近公共祖先

	两个节点 p,q 分为两种情况：
		p 和 q 在相同子树中
		p 和 q 在不同子树中


	从根节点遍历，递归向左右子树查询节点信息
	递归终止条件：如果当前节点为空或等于 p 或 q，则返回当前节点

	递归遍历左右子树，如果左右子树查到节点都不为空，
	则表明 p 和 q 分别在左右子树中，因此，当前节点即为最近公共祖先；
	如果左右子树其中一个不为空，则返回非空节点。
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// root == nil  已经到了叶子节点没有找到 直接返回
	// p == root、q == root p或者q的祖先节点就是root
	if root == nil || p == root || q == root {
		return root
	}

	// 以root为根的左子树中 p和q的公共祖先
	left := lowestCommonAncestor(root.Left, p, q)
	// 以root为根的右子树中 p和q的公共祖先
	right := lowestCommonAncestor(root.Right, p, q)

	// left和right都不为空的时候 说明root就是最近的公共祖先
	if left != nil && right != nil {
		return root
	}

	//左子树中没找到 返回右子树
	if left == nil {
		return right
	} else {
		// 说明右子树为空 左子树中去找
		return left
	}
}

/*
*
96. 不同的二叉搜索树
*/
func numTrees(n int) int {
	cntMap := make([]int, n+1)

	var memo func(int) int

	memo = func(cur int) int {
		if cur == 0 || cur == 1 {
			return 1
		}

		if cntMap[cur] != 0 {
			return cntMap[cur]
		}

		cnt := 0
		for i := 1; i <= cur; i++ {
			left := memo(i - 1)

			right := memo(n - i)

			cnt += left * right
		}

		cntMap[cur] = cnt

		return cnt
	}

	return memo(n)
}

/*
*	给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。

candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
*/
func combinationSum(candidates []int, target int) [][]int {
	return nil
}
