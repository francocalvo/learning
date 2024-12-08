// Before reading the puzzle correcly, I wanted to know the _distance_ the
// patrol went through. I did a matrix with the distances to the next block,
// being the block a wall or the end of the grid. Then, I started to iterate
// through the grid, updating the position and the direction of the patrol.
// This is not question. I needed the distinct positions the patrol went through
//

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"sync"
	"sync/atomic"
	"time"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func print_map(grid [][]int, pos [3]int) {
	for i, row := range grid {
		for j, block := range row {
			if i == pos[0] && j == pos[1] {
				switch pos[2] {
				case 0:
					fmt.Print("^")
				case 1:
					fmt.Print(">")
				case 2:
					fmt.Print("v")
				case 3:
					fmt.Print("<")
				}
			} else if block == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func dis_to_incr(direction int) (int, int) {
	switch direction {
	case 0:
		return -1, 0
	case 1:
		return 0, 1
	case 2:
		return 1, 0
	case 3:
		return 0, -1
	default:
		fmt.Println("Invalid direction")
		return 0, 0
	}
}

func dir_to_txt(direction int) string {
	switch direction {
	case 0:
		return "Up"
	case 1:
		return "Right"
	case 2:
		return "Down"
	case 3:
		return "Left"
	default:
		return "Invalid"
	}
}
func main() {
  load_time := time.Now()
	fmt.Println("AoC 2024 - Day 6")
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	buf := bufio.NewReader(file)
	grid := make([][]int, 0)
	// x, y, direction
	// 0 = up, 1 = right, 2 = down, 3 = left
	var pos [3]int
	cursor := 0
	for {
		line, _, err := buf.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println("Line:", string(line))
		grid = append(grid, make([]int, 0))
		for _, char := range line {
			if char == '.' {
				grid[cursor] = append(grid[cursor], 0)
			} else if char == '#' {
				grid[cursor] = append(grid[cursor], 1)
			} else if char == '^' {
				grid[cursor] = append(grid[cursor], 0)
				pos[0] = cursor
				pos[1] = len(grid[cursor]) - 1
				pos[2] = 0
			}
		}
		cursor++
	}

  fmt.Println("Load time:", time.Since(load_time))

	start := time.Now()

	visited := make(map[[2]int][]int)
	visited[[2]int{pos[0], pos[1]}] = []int{pos[2]}
	count := 1

	var initial_pos [3]int
	initial_pos[0] = pos[0]
	initial_pos[1] = pos[1]
	initial_pos[2] = pos[2]

	// Part 1
	for {
		incr_c, incr_r := dis_to_incr(pos[2])
		next_pos := [3]int{pos[0] + incr_c, pos[1] + incr_r, pos[2]}
		if next_pos[0] < 0 || next_pos[0] >= len(grid) || next_pos[1] < 0 || next_pos[1] >= len(grid[0]) {
			if _, ok := visited[[2]int{pos[0], pos[1]}]; !ok {
				if _, ok := visited[[2]int{pos[0], pos[1]}]; !ok {
					visited[[2]int{pos[0], pos[1]}] = []int{pos[2]}
				} else {
					visited[[2]int{pos[0], pos[1]}] = append(visited[[2]int{pos[0], pos[1]}], pos[2])
				}
			}
			break
		}
		if grid[next_pos[0]][next_pos[1]] == 1 {
			pos[2] = get_next_dir(pos[2])
			continue
		}
		pos = next_pos
		if _, ok := visited[[2]int{pos[0], pos[1]}]; !ok {
			visited[[2]int{pos[0], pos[1]}] = []int{pos[2]}
			count++
		} else {
			visited[[2]int{pos[0], pos[1]}] = append(visited[[2]int{pos[0], pos[1]}], pos[2])
		}
	}

	fmt.Println("Part 1 took:", time.Since(start))

	// Part 2
	var loops atomic.Int64 // Thread-safe counter
	txt := make([]string, 0)
	var txtMutex sync.Mutex // Mutex for the txt slice since we'll be appending to it
	var wg sync.WaitGroup   // WaitGroup to wait for all goroutines to finish

	for k, v := range visited {
		if [2]int{initial_pos[0], initial_pos[1]} == k {
			continue
		}

		wg.Add(1)
		go func(k [2]int, v []int) {
			defer wg.Done()

			// Clone the grid
			new_grid := make([][]int, 0)
			for i := range grid {
				new_row := make([]int, 0)
				for j := range grid[i] {
					new_row = append(new_row, grid[i][j])
				}
				new_grid = append(new_grid, new_row)
			}
			new_grid[k[0]][k[1]] = 1

			// Clone the visited map
			new_visited := make(map[[2]int][]int)

			// Check if the new grid has a loop
			if is_loop(new_grid, new_visited, initial_pos) {
				loops.Add(1) // Thread-safe increment
				txtMutex.Lock()
				txt = append(txt, fmt.Sprintf("Loop on %v with direction %v", k, v))
				txtMutex.Unlock()
			}
		}(k, v)
	}
	wg.Wait() // Wait for all goroutines to finish

	fmt.Println("Part 2 took:", time.Since(start))
  fmt.Println("Total time:", time.Since(load_time))
	fmt.Println("Visited:", count)
	fmt.Println("Loops:", loops.Load())
}

func is_loop(grid [][]int, visited map[[2]int][]int, pos [3]int) bool {
	for {
		incr_c, incr_r := dis_to_incr(pos[2])
		next_pos := [3]int{pos[0] + incr_c, pos[1] + incr_r, pos[2]}
		if next_pos[0] < 0 || next_pos[0] >= len(grid) || next_pos[1] < 0 || next_pos[1] >= len(grid[0]) {
			break
		}
		if grid[next_pos[0]][next_pos[1]] == 1 {
			pos[2] = get_next_dir(pos[2])
			continue
		}
		pos = next_pos
		if _, ok := visited[[2]int{pos[0], pos[1]}]; !ok {
			visited[[2]int{pos[0], pos[1]}] = []int{pos[2]}
		} else {
			if slices.Contains(visited[[2]int{pos[0], pos[1]}], pos[2]) {
				return true
			}
			visited[[2]int{pos[0], pos[1]}] = append(visited[[2]int{pos[0], pos[1]}], pos[2])
		}
	}
	return false
}

func get_next_dir(dir int) int {
	if dir == 3 {
		return 0
	}
	return dir + 1
}
