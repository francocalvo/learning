package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func dis_to_incr(direction int) (int, int) {
	switch direction {
	case 0:
		return 0, -1
	case 1:
		return 1, 0
	case 2:
		return 0, 1
	case 3:
		return -1, 0
	default:
		return 0, 0
	}
}

func get_distance_iter(grid [][]int, blocks_distances [][][4]int, x_pos, y_pos, dir int) int {
  fmt.Println("%%%%%%%%%%%%%% Get distance iter")
	fmt.Println("Pos:", x_pos, y_pos, dir)
	if y_pos > len(grid) || x_pos > len(grid[0]) {
		return 0
	}

	x_incr, y_incr := dis_to_incr(dir)
	ex_x_pos := x_pos - x_incr
	ex_y_pos := y_pos - y_incr
  fmt.Println("Ex:", ex_x_pos, ex_y_pos)
	if ex_x_pos > 0 && ex_y_pos > 0 {
		fmt.Println("Skipping because of ex")
		return blocks_distances[ex_x_pos][ex_y_pos][dir] + x_incr + y_incr
	}

	counter := 0
	fmt.Println("Incr:", x_incr, y_incr)
	for x_pos < len(grid[0]) && y_pos < len(grid) && grid[y_pos][x_pos] != '#' {
		fmt.Println("Pos:", x_pos, y_pos, "Counter:", counter)
		x_pos += x_incr
		y_pos += y_incr
    fmt.Println("Grid?")

		if x_pos < 0 || y_pos < 0 {
      fmt.Println("Breaking because of bounds")
			break
		}

    fmt.Println("Grid?")
    fmt.Println("Pos:", x_pos, y_pos, "Counter:", counter)
		counter++
	}
	return counter
}

func get_distances(grid [][]int, blocks_distances [][][4]int) {
	for y := 0; y < len(grid); y++ {
		blocks_distances = append(blocks_distances, make([][4]int, 0))
		for x := 0; x < len(grid[0]); x++ {
      fmt.Println("")
      fmt.Println("&&&&&&&&&&&&&%%%%%%%%%%%%%%&&&&&&&&&&&&&&&&")
      fmt.Println("&&&&&&&&&&&&&%%%%%%%%%%%%%%&&&&&&&&&&&&&&&&")
      fmt.Println("Getting distance for pos:", x, y)
			if grid[y][x] == '#' {
				blocks_distances[x] = append(blocks_distances[x], [4]int{0, 0, 0, 0})
			}
      fmt.Println("Not a wall")
			blocks_distances[y] = append(blocks_distances[y], [4]int{0, 0, 0, 0})
      fmt.Println("Getting distances")
			for dir := 0; dir < 4; dir++ {
        res := get_distance_iter(grid, blocks_distances, x, y, dir)
				blocks_distances[x][y][dir] = res
        fmt.Println(blocks_distances[x][y])
			}
		}
	}
}

func main() {
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
	blocks_distances := make([][][4]int, 0)
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
				pos[0] = len(grid[cursor]) - 1
				pos[1] = cursor
				pos[2] = 0
			}
		}
		cursor++
	}

	fmt.Println("Pos:", pos)
	for _, row := range grid {
		fmt.Println(row)
	}

	get_distances(grid, blocks_distances)

	for _, row := range blocks_distances {
		fmt.Println(row)
	}
}
