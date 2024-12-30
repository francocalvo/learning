package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

func groupAnagrams(strs []string) [][]string {
	hm := make(map[string][]string)

	for _, str := range strs {
		l_str := strings.Split(str, "")
		sort.Strings(l_str)
		s_str := strings.Join(l_str, "")
		hm[s_str] = append(hm[s_str], str)
	}

	res := make([][]string, 0)
	for _, v := range hm {
		res = append(res, v)
	}

	return res
}

func main() {
	// Convert the test inputs to []interface{} format
	input := make([][]interface{}, 3)

	input[0] = []interface{}{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}}
	input[1] = []interface{}{[]string{""}}
	input[2] = []interface{}{[]string{"a"}}

	expected := []interface{}{[][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}, [][]string{{""}}, [][]string{{"a"}}}

	start := time.Now()

	for j := 0; j < len(input); j++ {
		result := groupAnagrams(input[j][0].([]string))

		for i := 0; i < len(result); i++ {
			if len(result[i]) != len(expected[j].([][]string)[i]) {
				fmt.Println("Test case ", j, " failed")
				return
			}
		}
	}

	elapsed := time.Since(start).Seconds()
	fmt.Println("049. groupAnagrams", elapsed)
}
