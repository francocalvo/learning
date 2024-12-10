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
	fmt.Println("AoC 2024 - Day 7")
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	buf := bufio.NewReader(file)

	eqs := make(map[string][][2]int)
	cs := 0
	rs := 0
	for {
		line, _, err := buf.ReadLine()
		if err == io.EOF {
			break
		}
		rs++

		for i, c := range string(line) {
			if rs == 1 {
				cs++
			}
			if c == '.' {
				continue
			} else {
				if _, ok := eqs[string(c)]; !ok {
					eqs[string(c)] = make([][2]int, 0)
				}
				eqs[string(c)] = append(eqs[string(c)], [2]int{rs - 1, i})
			}
		}
	}
	fmt.Println("Load time:", time.Since(load_time))
	start := time.Now()

	grid := make(map[[2]int]bool)
	c := 0

	for _, v := range eqs {
		c += process_combination(v, grid, rs, cs)
	}

	fmt.Println("Antinodes:", c)
	fmt.Println("Part 1 time:", time.Since(start))

	start = time.Now()

	grid = make(map[[2]int]bool)
	c = 0

	for _, v := range eqs {
		c += process_combination2(v, grid, rs, cs)
	}

	fmt.Println("Antinodes:", c)
	fmt.Println("Part 2 time:", time.Since(start))
}

func process_combination(comb [][2]int, grid map[[2]int]bool, rs int, cs int) int {
	return process_combination_iter(comb[0], comb[1:], grid, rs, cs)
}

func process_combination_iter(head [2]int, tail [][2]int, grid map[[2]int]bool, rs int, cs int) int {
	if len(tail) == 0 {
		return 0
	}

	c := 0

	for _, t := range tail {
		dis := [2]int{head[0] - t[0], head[1] - t[1]}

		an_1 := [2]int{head[0] + dis[0], head[1] + dis[1]}
		if an_1[0] >= 0 && an_1[0] < rs && an_1[1] >= 0 && an_1[1] < cs {
			if _, ok := grid[an_1]; !ok {
				grid[an_1] = true
				c++
			}
		}

		an_2 := [2]int{t[0] - dis[0], t[1] - dis[1]}
		if an_2[0] >= 0 && an_2[0] < rs && an_2[1] >= 0 && an_2[1] < cs {
			if _, ok := grid[an_2]; !ok {
				grid[an_2] = true
				c++
			}
		}
	}

	return c + process_combination_iter(tail[0], tail[1:], grid, rs, cs)
}

func process_combination2(comb [][2]int, grid map[[2]int]bool, rs int, cs int) int {
	return process_combination_iter2(comb[0], comb[1:], grid, rs, cs, 0)
}

func process_combination_iter2(head [2]int, tail [][2]int, grid map[[2]int]bool, rs int, cs int, acc int) int {
	if len(tail) == 0 {
		return acc
	}

	if _, ok := grid[head]; !ok {
		grid[head] = true
		acc++
	}

	for _, t := range tail {
		dis := unit_distance(head[0]-t[0], head[1]-t[1])

		dir_a := [2]int{head[0] + dis[0], head[1] + dis[1]}
		for dir_a[0] >= 0 && dir_a[0] < rs && dir_a[1] >= 0 && dir_a[1] < cs {
			if v, ok := grid[dir_a]; !ok || !v {
				grid[dir_a] = true
				acc++
			}
			dir_a[0] += dis[0]
			dir_a[1] += dis[1]
		}

		dir_b := [2]int{head[0] - dis[0], head[1] - dis[1]}
		for dir_b[0] >= 0 && dir_b[0] < rs && dir_b[1] >= 0 && dir_b[1] < cs {
			if v, ok := grid[dir_b]; !ok || !v {
				grid[dir_b] = true
				acc++
			}
			dir_b[0] -= dis[0]
			dir_b[1] -= dis[1]
		}
	}

	return process_combination_iter2(tail[0], tail[1:], grid, rs, cs, acc)
}

func unit_distance(a, b int) [2]int {
	if a == b {
		return [2]int{1, 1}
	}
	if a == 1 || b == 1 {
		return [2]int{a, b}
	}
	if a%b == 0 {
		return [2]int{a / b, 1}
	}
	if b%a == 0 {
		return [2]int{1, b / a}
	}
	if b%a != 0 && a%b != 0 {
		return [2]int{a, b}
	}
	if a%b != 0 && b%a == 0 {
		return unit_distance(a, b/a)
	}
	if b%a != 0 && a%b == 0 {
		return unit_distance(a/b, b)
	}
	return [2]int{1, 1}
}
