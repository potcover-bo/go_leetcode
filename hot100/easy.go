package main

/*
461
*两个整数之间的 汉明距离 指的是这两个数字对应二进制位不同的位置的数目。

给你两个整数 x 和 y，计算并返回它们之间的汉明距离。
*/
func hammingDistance(x int, y int) int {
	res := x ^ y

	result := 0

	for res > 0 {
		if res%2 == 1 {
			result++
		}
		res /= 2
	}

	return result
}

/*
* 226 给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	temp := root.Left
	root.Left = root.Right
	root.Right = temp

	if root.Left != nil {
		root.Left = invertTree(root.Left)
	}

	if root.Right != nil {
		root.Right = invertTree(root.Right)
	}

	return root
}

/*
*
617
想象一下，当你将其中一棵覆盖到另一棵之上时，两棵树上的一些节点将会重叠（而另一些不会）。你需要将这两棵树合并成一棵新二叉树。合并的规则是：如果两个节点重叠，那么将这两个节点的值相加作为合并后节点的新值；否则，不为 null 的节点将直接作为新二叉树的节点。

返回合并后的二叉树。

注意: 合并过程必须从两个树的根节点开始。
*/
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil || root2 == nil {
		if root1 == nil {
			return root2
		}
		return root1
	}

	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)

	return root1
}

/*
* 104
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
*/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

/*
94
*给定一个二叉树的根节点 root ，返回 它的 中序 遍历
*/
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	// 声明一个函数类型的变量
	var inorderTraversalRecur func(node *TreeNode)
	inorderTraversalRecur = func(node *TreeNode) {

		if node == nil {
			return
		}
		inorderTraversalRecur(node.Left)
		res = append(res, node.Val)
		inorderTraversalRecur(node.Right)
	}
	inorderTraversalRecur(root)
	return res
}

/*
*

1、两数之和
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target
的那 两个 整数，并返回它们的数组下标。
*/
func twoSum(nums []int, target int) []int {
	helpMap := make(map[int]int)

	for index, val := range nums {
		if helpMap[target-val] == 0 {
			helpMap[val] = index + 1
		} else {
			return []int{helpMap[target-val] - 1, index}
		}
	}

	return []int{}
}

/*
*448. 找到所有数组中消失的数字
 */
func findDisappearedNumbers(nums []int) []int {
	/*res := make([]int, 0)
	temp := make([]int, len(nums)+1)

	for _, val := range nums {
		temp[val] = val
	}

	for i := 1; i < len(temp); i++ {
		if temp[i] != i {
			res = append(res, i)
		}
	}

	return res*/
	length := len(nums)
	for _, val := range nums {
		x := (val - 1) % length
		nums[x] += length
	}

	res := make([]int, 0)

	for index, val := range nums {
		if val <= length {
			res = append(res, index+1)
		}
	}

	return res
}

/*
*
543
给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。
*/
func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0

	var dfs func(*TreeNode) int

	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		lDepth := dfs(node.Left)
		rDepth := dfs(node.Right)

		ans = max(ans, lDepth+rDepth+1)

		return max(lDepth, rDepth) + 1
	}

	dfs(root)

	return ans - 1
}

/*
*

283
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

请注意 ，必须在不复制数组的情况下原地对数组进行操作。
*/
func moveZeroes(nums []int) {
	l, r := 0, 0

	for r < len(nums) {
		if nums[r] != 0 {
			nums[l] = nums[r]
			l += 1
		}
		r += 1
	}

	for i := l; i < len(nums); i++ {
		nums[i] = 0
	}
}

/*
*
20、给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。
*/
func isValid(s string) bool {
	stack := make([]byte, 0)
	data := []byte(s)
	for _, c := range data {
		if c == '[' || c == '(' || c == '{' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}

			val := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if c == ')' && val != '(' {
				return false
			}

			if c == '}' && val != '{' {
				return false
			}

			if c == ']' && val != '[' {
				return false
			}

		}
	}

	if len(stack) != 0 {
		return false
	}

	return true
}

/*
*21
将两个升序链表合并为一个新的 升序 链表并返回。
新链表是通过拼接给定的两个链表的所有节点组成的。
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
			cur = cur.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
			cur = cur.Next
		}
	}

	if list1 != nil {
		cur.Next = list1
	}

	if list2 != nil {
		cur.Next = list2
	}

	return head.Next
}

/*
*101. 对称二叉树
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}

	return dfsSymmetric(root.Left, root.Right)
}

func dfsSymmetric(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}

	if left.Val == right.Val && dfsSymmetric(left.Left, right.Right) && dfsSymmetric(left.Right, right.Left) {
		return true
	} else {
		return false
	}
}

/*
*136. 只出现一次的数字
 */
func singleNumber(nums []int) int {
	res := 0

	for _, val := range nums {
		res ^= val
	}

	return res
}

/*
*
70. 爬楼梯
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
*/
func climbStairs(n int) int {
	prePre, pre, cur := 0, 0, 1
	for i := 1; i <= n; i++ {
		prePre = pre
		pre = cur
		cur = pre + prePre

	}

	return cur
}

/*
*
169. 多数元素

给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。
*/
func majorityElement(nums []int) int {
	candidate := nums[0] // 投票对象
	cnt := 1             // 票数

	for i := 1; i < len(nums); i++ {
		if nums[i] == candidate {
			cnt += 1
		} else {
			cnt -= 1
			if cnt < 0 {
				candidate = nums[i]
				cnt = 1
			}
		}
	}

	return candidate
}

/*
*
反转链表
*/
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode

	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}

	return pre
}

/*
*141. 环形链表
 */
func hasCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}

	}

	return false
}
