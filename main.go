package main

import (
	"fmt"
	"potcover.cc/go_leetcode/utils"
)

func main() {
	q := &utils.PriorityQueue{}

	q.Add(3)
	q.Add(1)
	q.Add(4)
	q.Add(2)
	q.Add(52)
	q.Add(25)
	q.Add(71)

	size := q.Size()

	for i := 0; i < size; i++ {
		fmt.Printf("%d\t", q.Poll())
	}

}
