package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)
 
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	load_time := time.Now()
	fmt.Println("AoC 2024 - Day 10")

	data := load_data()

	fmt.Println("Load time: ", time.Since(load_time))
	start_time := time.Now()
	processed_tile := make(map[[2]int]bool)

	cost := 0
	for r, row := range data {
		for c := range row {
			if _, ok := processed_tile[[2]int{r, c}]; ok {
				continue
			} else {
				area, perimeter := process_region(data, [2]int{r, c}, &processed_tile)
				// fmt.Println("For region ", data[r][c], "at position", r, c, "area is", area, "and perimeter is", perimeter)
				cost += area * perimeter
			}
		}
	}

	fmt.Println("Part 1 cost: ", cost)
	fmt.Println("Part 1 time: ", time.Since(start_time))

	part_2 := time.Now()
	processed_tile = make(map[[2]int]bool)
	cost = 0
	for r, row := range data {
		for c := range row {
			if _, ok := processed_tile[[2]int{r, c}]; ok {
				continue
			} else {
				area, sides := process_region_sides(data, [2]int{r, c}, &processed_tile)
				cost += area * sides
			}
		}
	}

	fmt.Println("Part 2 cost: ", cost)
	fmt.Println("Part 2 time: ", time.Since(part_2))
	fmt.Println("Total time: ", time.Since(load_time))

}

func process_region_sides(data [][]rune, tile [2]int, processed_tile *map[[2]int]bool) (int, int) {
	queue := make([][2]int, 0)
	queue = append(queue, tile)
	area := 0
	sides := 0

	area_tiles := make(map[[2]int]bool)

	for {
		// If queue is empty, we are done
		if len(queue) == 0 {
			break
		}
		// If the tile has already been processed, skip it
		if _, ok := (*processed_tile)[queue[0]]; ok {
			queue = queue[1:]
			continue
		}
		area += process_tile_sides(data, queue[0], &queue, &area_tiles)
		(*processed_tile)[queue[0]] = true
	}

	vxs := make(map[[2]int]int)
	for tile := range area_tiles {
		check_vertexes(&area_tiles, tile, len(data), &vxs)
	}
  for _, v := range vxs {
    sides += v
  }
	return area, sides
}

func check_vertexes(area_tiles *map[[2]int]bool, tile [2]int, max int, vxs *map[[2]int]int) {
	left := false
	if _, ok := (*area_tiles)[[2]int{tile[0], tile[1] - 1}]; ok {
		left = true
	}
	right := false
	if _, ok := (*area_tiles)[[2]int{tile[0], tile[1] + 1}]; ok {
		right = true
	}
	top := false

	if _, ok := (*area_tiles)[[2]int{tile[0] - 1, tile[1]}]; ok {
		top = true
	}

	bottom := false
	if _, ok := (*area_tiles)[[2]int{tile[0] + 1, tile[1]}]; ok {
		bottom = true
	}

	left_top := false
	if _, ok := (*area_tiles)[[2]int{tile[0] - 1, tile[1] - 1}]; ok {
		left_top = true
	}
	left_bottom := false
	if _, ok := (*area_tiles)[[2]int{tile[0] + 1, tile[1] - 1}]; ok {
		left_bottom = true
	}
	right_top := false
	if _, ok := (*area_tiles)[[2]int{tile[0] - 1, tile[1] + 1}]; ok {
		right_top = true
	}
	right_bottom := false
	if _, ok := (*area_tiles)[[2]int{tile[0] + 1, tile[1] + 1}]; ok {
		right_bottom = true
	}

	lt_v := tile
	lb_v := [2]int{tile[0] + 1, tile[1]}
	rt_v := [2]int{tile[0], tile[1] + 1}
	rb_v := [2]int{tile[0] + 1, tile[1] + 1}

	if tile[0] == 0 && tile[1] == 0 {
		(*vxs)[lt_v] = 1
	}
	if tile[0] == 0 && tile[1] == max-1 {
		(*vxs)[rt_v] = 1
	}
	if tile[0] == max-1 && tile[1] == 0 {
		(*vxs)[lb_v] = 1
	}
	if tile[0] == max-1 && tile[1] == max-1 {
		(*vxs)[rb_v] = 1
	}

	// Left top
	if (!left && !top && left_top) || (left && top && !left_top) || (!left && !top && !left_top) {
		(*vxs)[lt_v] = 1
		if !left && !top && left_top {
			(*vxs)[lt_v] = 2
		}
	}

	if (!left && !bottom && left_bottom) || (left && bottom && !left_bottom) || (!left && !bottom && !left_bottom) {
		(*vxs)[lb_v] = 1
		if !left && !bottom && left_bottom {
			(*vxs)[lb_v] = 2
		}
	}

	if (!right && !top && right_top) || (right && top && !right_top) || (!right && !top && !right_top) {
		(*vxs)[rt_v] = 1
		if !right && !top && right_top {
			(*vxs)[rt_v] = 2
		}
	}

	if (!right && !bottom && right_bottom) || (right && bottom && !right_bottom) || (!right && !bottom && !right_bottom) {
		(*vxs)[rb_v] = 1
		if !right && !bottom && right_bottom {
			(*vxs)[rb_v] = 2
		}
	}
	return
}

func process_tile_sides(data [][]rune, tile [2]int, queue *[][2]int, area_tiles *map[[2]int]bool) int {
	area := 1

	// Check top
	if tile[0] > 0 && data[tile[0]-1][tile[1]] == data[tile[0]][tile[1]] {
		(*queue) = append((*queue), [2]int{tile[0] - 1, tile[1]})
	}

	// Check bottom
	if tile[0] < len(data)-1 && data[tile[0]+1][tile[1]] == data[tile[0]][tile[1]] {
		(*queue) = append((*queue), [2]int{tile[0] + 1, tile[1]})
	}

	// Check left
	if tile[1] > 0 && data[tile[0]][tile[1]-1] == data[tile[0]][tile[1]] {
		(*queue) = append((*queue), [2]int{tile[0], tile[1] - 1})
	}

	// Check right
	if tile[1] < len(data[0])-1 && data[tile[0]][tile[1]+1] == data[tile[0]][tile[1]] {
		(*queue) = append((*queue), [2]int{tile[0], tile[1] + 1})
	}

	(*area_tiles)[tile] = true

	return area
}

func process_region(data [][]rune, tile [2]int, processed_tile *map[[2]int]bool) (int, int) {
	queue := make([][2]int, 0)
	queue = append(queue, tile)
	area, perimeter := 0, 0

	for {
		// If queue is empty, we are done
		if len(queue) == 0 {
			break
		}
		// If the tile has already been processed, skip it
		if _, ok := (*processed_tile)[queue[0]]; ok {
			queue = queue[1:]
			continue
		}
		n_area, n_perimeter := process_tile(data, queue[0], &queue)
		area += n_area
		perimeter += n_perimeter
		(*processed_tile)[queue[0]] = true
	}

	return area, perimeter
}

func process_tile(data [][]rune, tile [2]int, queue *[][2]int) (int, int) {
	area, perimeter := 1, 0

	// Check top
	if tile[0] > 0 {
		if data[tile[0]-1][tile[1]] == data[tile[0]][tile[1]] {
			(*queue) = append((*queue), [2]int{tile[0] - 1, tile[1]})
		} else {
			perimeter++
		}
	} else {
		perimeter++
	}

	// Check bottom
	if tile[0] < len(data)-1 {
		if data[tile[0]+1][tile[1]] == data[tile[0]][tile[1]] {
			(*queue) = append((*queue), [2]int{tile[0] + 1, tile[1]})
		} else {
			perimeter++
		}
	} else {
		perimeter++
	}

	// Check left
	if tile[1] > 0 {
		if data[tile[0]][tile[1]-1] == data[tile[0]][tile[1]] {
			(*queue) = append((*queue), [2]int{tile[0], tile[1] - 1})
		} else {
			perimeter++
		}
	} else {
		perimeter++
	}

	// Check right
	if tile[1] < len(data[0])-1 {
		if data[tile[0]][tile[1]+1] == data[tile[0]][tile[1]] {
			(*queue) = append((*queue), [2]int{tile[0], tile[1] + 1})
		} else {
			perimeter++
		}
	} else {
		perimeter++
	}

	return area, perimeter
}

func load_data() [][]rune {
	data := make([][]rune, 0)
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			return nil
		}
		line = line[:len(line)-1]
		row := make([]rune, 0)
		for _, c := range line {
			row = append(row, c)
		}
		data = append(data, row)
	}

	return data
}
