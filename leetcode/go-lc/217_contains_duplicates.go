// 217: Contains Duplicate
package main

import (
	"fmt"
	"time"
)

func containsDuplicate(nums []int) bool {
	hm := make(map[int]bool)

	for _, num := range nums {
		if _, ok := hm[num]; ok {
			return true
		}
		hm[num] = true
	}

	return false
}

func main() {
	// Convert the test inputs to []interface{} format
	input := make([][]interface{}, 3)
	input[0] = []interface{}{[]int{1, 2, 3, 1}}
	input[1] = []interface{}{[]int{1, 2, 3, 4}}
	input[2] = []interface{}{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}}

	expected := []interface{}{true, false, true}

	start := time.Now()

	for i := 0; i < 1000000; i++ {
		for j := 0; j < len(input); j++ {
			result := containsDuplicate(input[j][0].([]int))
			if result != expected[j] {
				panic("Test failed!")
			}
		}
	}
	elapsed := time.Since(start).Seconds()
	fmt.Println("217. Contains Duplicate", elapsed)
}
