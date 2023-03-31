package main

import "fmt"

func main() {
	permute := permute([]int{1, 2, 3})

	subsets := subsets([]int{1, 2, 3})

	fmt.Println(permute)
	fmt.Println(subsets)

}
