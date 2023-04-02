package utils

func getLeftChild(index int) int {
	return 2*index + 1
}

func getRightChild(index int) int {
	return 2*index + 2
}

func getParent(index int) int {
	return (index - 1) / 2
}

// PriorityQueue 优先级队列
type PriorityQueue struct {
	val []int
}

func (pq *PriorityQueue) heapInsert() {
	index := len(pq.val) - 1

	for index > 0 {
		parent := getParent(index)
		if parent >= 0 && pq.val[index] < pq.val[parent] {
			Swap(index, parent, pq.val)
			index = parent
		} else {
			return
		}
	}
}

func (pq *PriorityQueue) heapify() {
	index := 0
	for index < len(pq.val) {
		leftChildIndex := getLeftChild(index)
		rightChildIndex := getRightChild(index)

		minIndex := index
		if leftChildIndex < len(pq.val) && pq.val[leftChildIndex] < pq.val[minIndex] {
			minIndex = leftChildIndex
		}

		if rightChildIndex < len(pq.val) && pq.val[rightChildIndex] < pq.val[minIndex] {
			minIndex = rightChildIndex
		}

		if minIndex == index {
			return
		} else {
			Swap(index, minIndex, pq.val)
			index = minIndex
		}
	}
}

func (pq *PriorityQueue) Add(val int) {
	pq.val = append(pq.val, val)
	pq.heapInsert()
}

func (pq *PriorityQueue) Poll() int {
	if len(pq.val) > 0 {

		// 交换堆顶和最后一个元素
		Swap(0, len(pq.val)-1, pq.val)

		// 移除最后一个元素
		res := pq.val[len(pq.val)-1]
		pq.val = pq.val[:len(pq.val)-1]

		// 进行heapify
		pq.heapify()

		return res
	} else {
		panic("illegle")
	}
}

func (pq *PriorityQueue) Size() int {
	return len(pq.val)
}
