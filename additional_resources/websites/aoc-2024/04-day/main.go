package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

func check_xmas(r_center, c_center int, r_incr int, c_incr int, input []string) bool {
	// Check if we can get it inside bounds
	if r_center+r_incr*3 < 0 || r_center+r_incr*3 >= len(input) {
		return false
	}
	if c_center+c_incr*3 < 0 || c_center+c_incr*3 >= len(input[0]) {
		return false
	}

	// Check if next cell is M
	if input[r_center+r_incr][c_center+c_incr] != 'M' {
		return false
	}
	// Check if next cell is A
	if input[r_center+r_incr*2][c_center+c_incr*2] != 'A' {
		return false
	}
	// Check if next cell is S
	if input[r_center+r_incr*3][c_center+c_incr*3] != 'S' {
		return false
	}

	return true
}

func check_x_mas(r_center, c_center int, input []string) bool {
	// Check bounds
	if r_center+1 >= len(input) || r_center-1 < 0 {
		return false
	}
	if c_center+1 >= len(input[0]) || c_center-1 < 0 {
		return false
	}

	var side_a string = string(input[r_center-1][c_center-1]) + string(input[r_center+1][c_center+1])
	var side_b string = string(input[r_center-1][c_center+1]) + string(input[r_center+1][c_center-1])
	chances := []string{"SM", "MS"}
	if !slices.Contains(chances, side_a) || !slices.Contains(chances, side_b) {
		return false
	}
	return true
}

func main() {
	fmt.Println("Hello, World!")
	file, err := os.Open("input.txt")
	if err != nil {
		return
	}

	input := make([]string, 0)
	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return
		}
		input = append(input, strings.TrimSpace(line))
	}

	// Create matrix of same size as input with all numbers 0
	// matrix := make([][]int, len(input))
	//  for i := range matrix {
	//    matrix[i] = make([]int, len(input[0]))
	//

	// Part 1
	counter_p1 := 0
	counter_p2 := 0
	for r := range input {
		for c := range input[r] {
			// Part 1
			if input[r][c] != 'X' {
				dir_results := []bool{
					check_xmas(r, c, 0, 1, input),   // Right
					check_xmas(r, c, 1, 1, input),   // Down Right
					check_xmas(r, c, 1, 0, input),   // Down
					check_xmas(r, c, 1, -1, input),  // Down Left
					check_xmas(r, c, 0, -1, input),  // Left
					check_xmas(r, c, -1, -1, input), // Up Left
					check_xmas(r, c, -1, 0, input),  // Up
					check_xmas(r, c, -1, 1, input),  // Up Right
				}

				for _, res := range dir_results {
					if res {
						counter_p1++
					}
				}
			}

			if input[r][c] == 'A' {
				if check_x_mas(r, c, input) {
					counter_p2++
				}
			}
		}
	}

	fmt.Println("Part 1: ", counter_p1)
  fmt.Println("Part 2: ", counter_p2)
}
