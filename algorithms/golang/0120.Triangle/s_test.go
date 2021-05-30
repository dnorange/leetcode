package main

import "testing"

func TestTriangle(t *testing.T) {
	tests := []struct {
		name string
		arr  [][]int
	}{
		{
			name: "c1",
			arr: [][]int{
				[]int{1},
				[]int{1, 2},
				[]int{1, 2, 3},
				[]int{1, 2, 3, 4},
			},
		},
		{
			name: "c2",
			arr: [][]int{
				[]int{2},
				[]int{3, 4},
				[]int{6, 5, 7},
				[]int{4, 1, 8, 3},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := minimumTotal2(test.arr)
			t.Logf("%d \n", m)
		})
	}
}
