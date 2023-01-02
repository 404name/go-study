package main

import "sort"

func main() {
	nums := []int{1, 3, 3, 6}
	flag := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= 1
	})
	print(flag)
}
