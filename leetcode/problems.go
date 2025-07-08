package main

import (
	"fmt"
	"math"
)

func findKDistantIndices(nums []int, key int, k int) []int {
	var result []int

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[j] == key && int(math.Abs(float64(i-j))) <= k {
				result = append(result, i)
				break
			}
		}
	}

	return result
}

func main() {
	key := 9
	k := 1
	fmt.Println(findKDistantIndices([]int{3, 4, 9, 1, 3, 9, 5}, key, k))
}
