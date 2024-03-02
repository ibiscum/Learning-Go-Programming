package main

import (
	"fmt"
	"math/rand"
)

const size = 1024 * 1024

type numbers [size]int

func initialize(nums *numbers) {
	for i := 0; i < size; i++ {
		nums[i] = rand.Intn(10000)
	}
}

func max(nums *numbers) int {
	temp := nums[0]
	for _, val := range nums {
		if val > temp {
			temp = val
		}
	}
	return temp
}

func main() {
	var nums *numbers = new(numbers)
	initialize(nums)
	fmt.Println(nums)
	fmt.Println("Max num is: ", max(nums))
}
