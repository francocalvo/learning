package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
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
	rocks_num := 0

	for _, a := range data {
		rocks_num += process_rocks(a, 25)
	}

	fmt.Println("Part 1: ", rocks_num)
	fmt.Println("Part 1 time: ", time.Since(start_time))
	part2_start_time := time.Now()

	for _, a := range data {
		rocks_num += process_rocks(a, 75)
	}

	fmt.Println("Part 2: ", rocks_num)
	fmt.Println("Part 2 time: ", time.Since(part2_start_time))

}

// Key struct to uniquely identify function parameters
type Key struct {
	rock, blinks int
}

// Global memoization map
var memo = make(map[Key]int)

func process_rocks(rock int, blinks int) int {
	key := Key{rock, blinks}

	// Check if result is already memoized
	if val, exists := memo[key]; exists {
		return val
	}

	// Base case
	if blinks == 0 {
		memo[key] = 1
		return 1
	}

	var res int
	if rock == 0 {
		res = process_rocks(1, blinks-1)
	} else {
		digits := get_digits(rock)
		if len(digits)%2 == 0 {
			a, b := get_parts(digits)
			res = process_rocks(a, blinks-1) + process_rocks(b, blinks-1)
		} else {
			res = process_rocks(rock*2024, blinks-1)
		}
	}

	// Store the computed result in memo map
	memo[key] = res
	return res
}

func get_digits(n int) []int {
	digits := make([]int, 0)
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	slices.Reverse(digits)
	return digits
}

func get_parts(ns []int) (int, int) {
	mid := len(ns) / 2
	a := ns[:mid]
	b := ns[mid:]
	r_a := 0
	r_b := 0
	for _, num := range a {
		r_a = r_a*10 + num
	}
	for _, num := range b {
		r_b = r_b*10 + num
	}
	return r_a, r_b
}

func load_data() []int {
	data := make([]int, 0)
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
		for _, c := range strings.Split(line, " ") {
			n, _ := strconv.Atoi(string(c))
			data = append(data, n)
		}
	}

	return data
}
