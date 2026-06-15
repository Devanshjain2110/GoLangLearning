package main 

import (
	"fmt"
)

func sums(nums ...int) int {
	fmt.Print(nums, " ")
	total := 0
	for _, n := range nums {
		total += n
	}
	fmt.Println(total)
	return total
}

func main() {
	sums(1, 2)
	sums(1, 2, 3)

	numSlice := []int{1, 2, 3, 4}
	sums(numSlice...)
}
