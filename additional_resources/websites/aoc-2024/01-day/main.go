package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	file, err := os.Open("input.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	buf := bufio.NewReader(file)
	list_b := make([]int, 0)
	list_a := make([]int, 0)
	for {
		line, err := buf.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}

		parts := strings.Split(line[0:len(line)-1], "   ")
		if len(parts) != 2 {
			continue
		}

		a, err := strconv.Atoi(parts[0])
		b, err := strconv.Atoi(parts[1])
		list_a = append(list_a, a)
		list_b = append(list_b, b)
	}

	// First part of the exercise
	sort.Ints(list_a)
	sort.Ints(list_b)

	distance := 0
	for i := 0; i < len(list_a); i++ {
	  distance += Abs(list_a[i] - list_b[i])
	}

	fmt.Println("Distance score: ", distance)

	// Second part of the exercise
	b_map := make(map[int]int)
	for i := 0; i < len(list_b); i++ {
		curr_value, ok := b_map[list_b[i]]
		if !ok {
			b_map[list_b[i]] = 1
		} else {
			b_map[list_b[i]] = curr_value + 1
		}
	}

	sim_score := 0
	for i := 0; i < len(list_a); i++ {
		val, ok := b_map[list_a[i]]
		if !ok {
			val = 0
		}
		sim_score += val * list_a[i]
	}

	fmt.Println("Similarity score: ", sim_score)
}
