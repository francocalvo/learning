package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
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

	data, trailheads := load_data()

	for k, v := range trailheads {
		fmt.Println(k, v)
	}
	for i, ar := range data {
		fmt.Println(i, ar)
	}
	fmt.Println("Load time: ", time.Since(load_time))
	start_time := time.Now()

	memo := make(map[[2]int][][2]int)
	acc := 0
	for _, k := range trailheads {
		val := calc_paths(data, k, make(map[[2]int]bool), 0, memo)
		acc += val
	}

	fmt.Println("Number of paths: ", acc)
	fmt.Println("Part 1: ", time.Since(start_time))

	start_time = time.Now()
	memo = make(map[[2]int][][2]int)
	acc = 0
	for _, k := range trailheads {
		val := calc_ratings(data, k, make(map[[2]int]bool), 0, memo)
		acc += val
	}

	fmt.Println("Number of ratings: ", acc)
	fmt.Println("Part 2: ", time.Since(start_time))
}

func calc_paths(data [][]int, trailhead [2]int, visited map[[2]int]bool, acc int, memo map[[2]int][][2]int) int {
	if data[trailhead[0]][trailhead[1]] == 9 {
		if _, ok := visited[trailhead]; ok {
			return acc
		} else {
			visited[trailhead] = true
			return acc + 1
		}
	}
	row := trailhead[0]
	col := trailhead[1]

	if row > 0 && data[row][col] == data[row-1][col]-1 {
		acc = calc_paths(data, [2]int{row - 1, col}, visited, acc, memo)
	}

	if row < len(data)-1 && data[row][col] == data[row+1][col]-1 {
		acc = calc_paths(data, [2]int{row + 1, col}, visited, acc, memo)
	}

	if col > 0 && data[row][col] == data[row][col-1]-1 {
		acc = calc_paths(data, [2]int{row, col - 1}, visited, acc, memo)
	}

	if col < len(data[0])-1 && data[row][col] == data[row][col+1]-1 {
		acc = calc_paths(data, [2]int{row, col + 1}, visited, acc, memo)
	}

	return acc
}

func calc_ratings(data [][]int, trailhead [2]int, visited map[[2]int]bool, acc int, memo map[[2]int][][2]int) int {
	if data[trailhead[0]][trailhead[1]] == 9 {
		return acc + 1
	}
	row := trailhead[0]
	col := trailhead[1]

	if row > 0 && data[row][col] == data[row-1][col]-1 {
		acc = calc_ratings(data, [2]int{row - 1, col}, visited, acc, memo)
	}

	if row < len(data)-1 && data[row][col] == data[row+1][col]-1 {
		acc = calc_ratings(data, [2]int{row + 1, col}, visited, acc, memo)
	}

	if col > 0 && data[row][col] == data[row][col-1]-1 {
		acc = calc_ratings(data, [2]int{row, col - 1}, visited, acc, memo)
	}

	if col < len(data[0])-1 && data[row][col] == data[row][col+1]-1 {
		acc = calc_ratings(data, [2]int{row, col + 1}, visited, acc, memo)
	}

	return acc
}

func load_data() ([][]int, [][2]int) {
	data := make([][]int, 0)
	trailheads := make([][2]int, 0)
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	buf := bufio.NewReader(file)
	i := 0
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			return nil, nil
		}
		line = line[:len(line)-1]
		ar := make([]int, 0)
		for j, c := range line {
			if string(c) == "0" {
				trailheads = append(trailheads, [2]int{i, j})
			}
			n, _ := strconv.Atoi(string(c))
			fmt.Println(i, j, n)
			ar = append(ar, n)
		}
		data = append(data, ar)
		i++
	}

	return data, trailheads
}
