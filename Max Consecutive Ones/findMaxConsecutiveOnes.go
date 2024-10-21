package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	maxCount := 0

	for _, num := range nums {
		if num == 1 {
			count++
		} else {
			if count > maxCount {
				maxCount = count
			}
			count = 0
		}
	}
	if count > maxCount {
		maxCount = count
	}
	return maxCount
}

func main() {
	nums := []int{1, 1, 1, 1, 1, 0, 2, 1, 1, 1, 1}
	result := findMaxConsecutiveOnes(nums)
	fmt.Println(result)
}

//time : O(N)
//space : O(1)
