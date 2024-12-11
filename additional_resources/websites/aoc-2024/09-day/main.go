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
	fmt.Println("AoC 2024 - Day 9")
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	buf := bufio.NewReader(file)

	max_cells := 0

	mem := make([]int, 0)
	for {
		line, _, err := buf.ReadLine()
		if err == io.EOF {
			break
		}

		for i, c := range string(line) {

			n, err := strconv.Atoi(string(c))
			if err != nil {
				continue
			}
			if i%2 == 0 {
				max_cells += n
			}
			mem = append(mem, n)
		}
	}
	fmt.Println("Load time:", time.Since(load_time))
	start := time.Now()

	sum := part_1(mem, max_cells)

	fmt.Println("Sum:", sum)
	fmt.Println("Part 1 time:", time.Since(start))
	start2 := time.Now()

	sum2 := part_2(mem, max_cells)
	fmt.Println("Sum2:", sum2)
	fmt.Println("Part 2 time:", time.Since(start2))
	fmt.Println("Total time:", time.Since(load_time))
}

type Group struct {
	label int
	size  int
}

func part_2(mem []int, max_cells int) int {
	if len(mem)%2 == 0 {
		mem = mem[:len(mem)-1]
	}
	grps := make([]Group, 0)
	sum := 0
	for i := 0; i < len(mem); i++ {
		if i%2 == 0 {
			grps = append(grps, Group{label: i / 2, size: mem[i]})
		} else {
			grps = append(grps, Group{label: -1, size: mem[i]})
		}
	}

	for i := len(grps) - 1; i >= 0; i-- {
		if grps[i].label == -1 {
			continue
		}
		for j := 0; j < i; j++ {
			if grps[j].label == -1 && grps[j].size >= grps[i].size {
        pos := i
				if grps[j].size == grps[i].size {
					// Replace position
					grps[j].label = grps[i].label
				} else {
					// Insert position
					nw_space := grps[j].size - grps[i].size
					grps = insert(grps, j, grps[i])
					grps[j+1].size = nw_space
          pos ++
				}

				size := grps[pos].size
				if j != pos-1 && grps[pos-1].label == -1 {
					grps[pos-1].size += size
          if pos+1 < len(grps) && grps[pos+1].label == -1 {
						grps[pos-1].size += grps[pos+1].size
						grps[pos+1].size = 0
					}
					grps = delete(grps, pos)
				} else if pos+1 < len(grps) && grps[pos+1].label == -1 {
					grps[pos+1].size += size
					grps = delete(grps, pos)
				} else {
					grps[pos].label = -1
				}
				break
			}
		}
	}

	idx := 0
	for i := 0; i < len(grps); i++ {
		for j := 0; j < grps[i].size; j++ {
			if grps[i].label != -1 {
				sum += idx * grps[i].label
			}
			idx++
		}
	}

	return sum
}

func print_groups(grps []Group) {
	str := []string{}
	for _, g := range grps {
		for i := 0; i < g.size; i++ {
			if g.label == -1 {
				str = append(str, ".")
			} else {
				str = append(str, strconv.Itoa(g.label))
			}
		}
	}
	fmt.Println(str)
}

func delete(a []Group, index int) []Group {
	return append(a[:index], a[index+1:]...)
}

func insert(a []Group, index int, value Group) []Group {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

func part_1(mem []int, max_cells int) int {
	if len(mem)%2 == 0 {
		mem = mem[:len(mem)-1]
	}

	idx := 0
	gps := 0
	sum := 0
	queue := make([]int, 0)
	lag := len(mem) - 1
	processed_cells := 0

	for i := 0; i < len(mem); i++ {
		if i%2 == 0 {
			if mem[i] <= 0 {
				continue
			}
			for j := 0; j < mem[i]; j++ {
				if processed_cells >= max_cells {
					break
				}
				sum += idx * gps
				processed_cells++
				idx++
			}
			gps++
		} else if i%2 == 1 {
			left := mem[i]
			for left > 0 {
				if i >= max_cells {
					break
				}

				// Get group from back
				gp := (lag) / 2
				for j := 0; j < mem[lag]; j++ {
					queue = append(queue, gp)
				}
				lag -= 2

				// Fill void from front
				var limit int
				if len(queue) > left {
					limit = left
					left = 0
				} else {
					limit = len(queue)
					left -= len(queue)
				}

				for j := 0; j < limit; j++ {
					if processed_cells >= max_cells {
						break
					}
					sum += queue[0] * idx
					idx++
					processed_cells += 2 // fills one cell and empties one
					queue = queue[1:]
				}
			}
		}
	}

	return sum
}
