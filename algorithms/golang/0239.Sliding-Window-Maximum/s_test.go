package main

import "testing"

func TestS1(t *testing.T) {
	result := maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	t.Log(result)
}

func TestS2(t *testing.T) {
	result := maxSlidingWindow2([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	t.Log(result)
}
