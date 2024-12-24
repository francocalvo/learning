// 242: Valid Anagram
// Neetcode: "Arrays & Hashing"
package main

import (
	"fmt"
	"time"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hm := make(map[rune]int)

	for _, ch := range s {
		hm[ch] += 1
	}

	for _, ch := range t {
		val, ok := hm[ch]
		if !ok || val == 0 {
			return false
		}
		hm[ch] -= 1
	}

	return true
}

func e_242() {
	// Convert the test inputs to []interface{} format
	input := [][]string{
		{"anagram", "nagaram"},
		{"rat", "car"},
	}
	expected := []interface{}{true, false}

	start := time.Now()
	for j := 0; j < len(input); j++ {
		result := isAnagram(input[j][0], input[j][1])
		if result != expected[j] {
			panic("Test failed!")
		}
	}

	elapsed := time.Since(start).Seconds()
	fmt.Println("242. Valid Anagram", elapsed)
}
