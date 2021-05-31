package main

// deque solution
func maxSlidingWindow2(nums []int, k int) []int {
	slice := make([]int, 0)
	result := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		if i >= k && slice[0] == i-k {
			slice = slice[1:len(slice)]
		}
		for j := len(slice); j > 0; j-- {
			if nums[slice[j-1]] <= nums[i] {
				// fmt.Printf("s%v j%v last%v ", slice, j-1, nums[slice[j-1]])
				slice = slice[:len(slice)-1]
			} else {
				break
			}
		}
		slice = append(slice, i)
		// fmt.Printf("after%v\n", slice)
		if i >= k-1 {
			result = append(result, nums[slice[0]])
		}
	}

	return result
}
