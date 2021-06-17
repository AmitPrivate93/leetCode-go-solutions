package main

import (
	"fmt"
	"math"
	"sort"
)

func neares2Sum(nums []int, target int) []int {
	m := make(map[int]int)
	nearest := []int{0, 0}
	dif := math.MaxInt32
	for index, value := range nums {
		m[value] = index
	}
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	fmt.Println(nums)
	for left < right {
		tempDif := int(target - nums[left] - nums[right])
		if dif > tempDif && tempDif > 0 {
			dif = tempDif
			nearest[0] = nums[left]
			nearest[1] = nums[right]

		}
		if nums[left]+nums[right] > target {
			right--
		}
		if nums[left]+nums[right] < target {
			left++
		}
		if nums[left]+nums[right] == target {
			return []int{m[nums[left]], m[nums[right]]}
		}
	}
	return []int{m[nearest[0]], m[nearest[1]]}
}
func main() {
	//return array of index with nearest/equal sum to target
	fmt.Println(neares2Sum([]int{1, 3, 2, 4, 10, 12, 15}, 9))
}
