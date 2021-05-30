package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 4, 5, 6, 7}, 8))
}

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]*int)
	for i, x := range nums {
		// fmt.Printf("index is %d, value is %d\n", i, x)
		val := target - x
		if hashMap[val] != nil {
			return []int{i, *hashMap[val]}
		}
		newVal := i
		hashMap[x] = &newVal
	}
	return []int{}
}
