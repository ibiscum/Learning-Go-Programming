package main

import (
	"fmt"
	"math/rand"
)

const (
	size     = 1000
	variance = 100000
)

var (
	nums [size]int
)

func init() {
	for i := 0; i < size; i++ {
		nums[i] = rand.Intn(variance)
	}
}

func max(lo, hi int, nums [size]int) int {
	if lo == hi {
		return nums[lo]
	}

	mid := int((lo + hi) / 2)

	i := max(lo, mid, nums)
	j := max(mid+1, hi, nums)

	if i > j {
		return i
	}
	return j
}

func main() {
	fmt.Println(nums)
	fmt.Println("Max num is: ", max(0, size-1, nums))
}
