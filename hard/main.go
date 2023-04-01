package main

import "fmt"

func main() {
	queue := []int{1, 2, 3}

	queue = queue[1:]

	fmt.Println(queue)
}
