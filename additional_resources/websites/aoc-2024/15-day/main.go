package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	load_time := time.Now()
	fmt.Println("AoC 2024 - Day 10")
	walls, boxes, robot, moves := load_data()

	fmt.Println("Load time:", time.Since(load_time))
	start_time := time.Now()

	for _, move := range moves {
		robot = apply_move(robot, move, &walls, &boxes)
	}

	gps_sums := 0
	for box, status := range boxes {
		if status {
			gps_sums += box[0]*100 + box[1]
		}
	}

	fmt.Println("Robot:", robot)
	fmt.Println("GPS sum:", gps_sums)
	fmt.Println("Part 1 time:", time.Since(start_time))

	walls, boxes, robot, moves = load_data()
	walls, boxes, robot = double_wide(walls, boxes, robot)
	print_double_wide(walls, boxes, robot)

	for _, move := range moves {
		robot = apply_move_double(robot, move, &walls, &boxes)
		print_move(move)
		print_double_wide(walls, boxes, robot)
	}
}

func apply_move_double(robot [2]int, move [2]int, walls *map[[2]int]bool, boxes *map[[2]int]bool) [2]int {
	nx_move := can_move_double(robot, move, walls, boxes)

	if nx_move {
		robot = [2]int{robot[0] + move[0], robot[1] + move[1]}
		move_boxes_double(robot, move, boxes)
	}

	return robot
}

func move_boxes_double(pos [2]int, move [2]int, boxes *map[[2]int]bool) {
	fmt.Println("Moving boxes, from pos:", pos, "with move:", move)
	res_1, ok_1 := (*boxes)[[2]int{pos[0], pos[1]}]
	res_2, ok_2 := (*boxes)[[2]int{pos[0], pos[1] - 1}]

	if move[0] != 0 {
		fmt.Println("Moving vertically")
		if ok_1 && res_1 {
			fmt.Println("Moving box on axis 1")
			move_boxes_double([2]int{pos[0] + move[0], pos[1]}, move, boxes)
			move_boxes_double([2]int{pos[0] + move[0], pos[1] + 1}, move, boxes)
			(*boxes)[[2]int{pos[0] + move[0], pos[1]}] = true
			(*boxes)[[2]int{pos[0], pos[1]}] = false

		} else if ok_2 && res_2 {
			fmt.Println("Moving box on axis 2")
			move_boxes_double([2]int{pos[0] + move[0], pos[1] - 1}, move, boxes)
			move_boxes_double([2]int{pos[0] + move[0], pos[1]}, move, boxes)
			(*boxes)[[2]int{pos[0] + move[0], pos[1] - 1}] = true
			(*boxes)[[2]int{pos[0], pos[1] - 1}] = false
		}
	} else {
		fmt.Println("Moving horizontally")
		if ok_1 && res_1 {
			fmt.Println("REVISITAR REVISITAR")
		} else if move[1] == -1 && ok_2 && res_2 {
			fmt.Println("Moving box on horizontal axis")
			move_boxes_double([2]int{pos[0], pos[1] + move[1]*2}, move, boxes)
			fmt.Println("Came back from recursion")
			fmt.Println("Moving box on horizontal axis from", (*boxes)[[2]int{pos[0], pos[1] - 1}], "to", (*boxes)[[2]int{pos[0], pos[1] + move[1]*2}])
			(*boxes)[[2]int{pos[0], pos[1] + move[1]*2}] = true
			(*boxes)[[2]int{pos[0], pos[1] - 1}] = false
		}
	}
}

func can_move_double(pos [2]int, move [2]int, walls *map[[2]int]bool, boxes *map[[2]int]bool) bool {
	n_pos := [2]int{pos[0] + move[0], pos[1] + move[1]}
	// If stepping onto a wall, return false
	if res, ok := (*walls)[n_pos]; ok && res {
		return false
	}

	res_1, ok_1 := (*boxes)[[2]int{n_pos[0], n_pos[1]}]
	res_2, ok_2 := (*boxes)[[2]int{n_pos[0], n_pos[1] - 1}]
	if (ok_1 && res_1) || (ok_2 && res_2) {
		return can_move_double(n_pos, move, walls, boxes) && can_move_double([2]int{n_pos[0], n_pos[1] - 1}, move, walls, boxes)
	}

	return true
}

func double_wide(walls map[[2]int]bool, boxes map[[2]int]bool, robot [2]int) (map[[2]int]bool, map[[2]int]bool, [2]int) {
	n_boxes := make(map[[2]int]bool)
	n_walls := make(map[[2]int]bool)

	for box := range boxes {
		n_boxes[[2]int{box[0], box[1] * 2}] = true
	}
	for wall := range walls {
		n_walls[[2]int{wall[0], wall[1] * 2}] = true
		n_walls[[2]int{wall[0], wall[1]*2 + 1}] = true
	}

	robot = [2]int{robot[0], robot[1] * 2}

	return n_walls, n_boxes, robot
}

func print_double_wide(walls map[[2]int]bool, boxes map[[2]int]bool, robot [2]int) {
	grid := make([]string, 0)
	h, w := 10, 20
	for i := 0; i < h; i++ {
		row := make([]string, 0)
		for j := 0; j < w; j++ {
			coord := [2]int{i, j}
			if _, ok := walls[coord]; ok {
				row = append(row, "#")
			} else if res, ok := boxes[coord]; ok && res {
				row = append(row, "[]")
				j++
			} else if robot == coord {
				row = append(row, "@.")
				j++
			} else {
				row = append(row, "..")
				j++
			}
		}
		grid = append(grid, strings.Join(row, ""))
	}
	fmt.Println(strings.Join(grid, "\n"))
}

func apply_move(robot [2]int, move [2]int, walls *map[[2]int]bool, boxes *map[[2]int]bool) [2]int {
	nx_move, new_pos := can_move(robot, move, walls, boxes)

	if nx_move {
		robot = [2]int{robot[0] + move[0], robot[1] + move[1]}
		if robot[0] != new_pos[0] || robot[1] != new_pos[1] {
			(*boxes)[robot] = false
			(*boxes)[new_pos] = true
		}
	}

	return robot
}

func print_move(move [2]int) {
	s := ""
	if move[0] == 1 {
		s = "v"
	} else if move[0] == -1 {
		s = "^"
	} else if move[1] == 1 {
		s = ">"
	} else if move[1] == -1 {
		s = "<"
	}
	fmt.Printf("Move: %s\n", s)
}

func print_grid(walls *map[[2]int]bool, boxes *map[[2]int]bool, robot [2]int) {
	grid := make([]string, 0)
	h, w := 8, 8
	for i := 0; i < h; i++ {
		row := make([]string, 0)
		for j := 0; j < w; j++ {
			coord := [2]int{i, j}
			if _, ok := (*walls)[coord]; ok {
				row = append(row, "#")
			} else if res, ok := (*boxes)[coord]; ok && res {
				row = append(row, "O")
			} else if robot == coord {
				row = append(row, "@")
			} else {
				row = append(row, ".")
			}
		}
		grid = append(grid, strings.Join(row, ""))
	}
	fmt.Println(strings.Join(grid, "\n"))
}

func can_move(pos [2]int, move [2]int, walls *map[[2]int]bool, boxes *map[[2]int]bool) (bool, [2]int) {
	n_pos := [2]int{pos[0] + move[0], pos[1] + move[1]}
	// If stepping onto a wall, return false
	if res, ok := (*walls)[n_pos]; ok && res {
		return false, pos
	}

	if res, ok := (*boxes)[n_pos]; ok && res {
		return can_move(n_pos, move, walls, boxes)
	}

	return true, n_pos
}

func load_data() (map[[2]int]bool, map[[2]int]bool, [2]int, [][2]int) {
	walls := make(map[[2]int]bool)
	boxes := make(map[[2]int]bool)
	var robot [2]int
	moves := make([][2]int, 0)

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil, [2]int{}, nil
	}
	defer file.Close()

	buf := bufio.NewReader(file)
	lineNumber := 0
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading file:", err)
			return nil, nil, [2]int{}, nil
		}

		line = strings.TrimRight(line, "\n")
		lineNumber++

		// Check if the line contains move characters to identify the moves section
		if strings.ContainsAny(line, "<>^v") {
			// Parse move characters
			for pos, ch := range line {
				var move [2]int
				switch ch {
				case '<':
					move = [2]int{0, -1}
				case '>':
					move = [2]int{0, 1}
				case '^':
					move = [2]int{-1, 0}
				case 'v':
					move = [2]int{1, 0}
				default:
					fmt.Printf("Unknown move character '%c' at line %d, position %d\n", ch, lineNumber, pos)
					continue
				}
				moves = append(moves, move)
			}
			continue
		}

		// Parse grid characters
		for x, ch := range line {
			y := lineNumber - 1   // Assuming y starts at 0
			coord := [2]int{y, x} // row, col
			switch ch {
			case '#':
				walls[coord] = true
			case 'O':
				boxes[coord] = true
			case '@':
				robot = coord
			case '.':
				// Empty space, do nothing
			default:
				fmt.Printf("Unknown grid character '%c' at line %d, position %d\n", ch, lineNumber, x)
			}
		}
	}

	return walls, boxes, robot, moves
}
