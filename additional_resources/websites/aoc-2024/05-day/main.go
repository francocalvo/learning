package main

// 6205 too high

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func check_rules(updates []string, rules map[string][]string) bool {
	excluded := make([]string, 0)
	for i := range updates {
		if slices.Contains(excluded, updates[i]) {
			return false
		}
		excluded = append(excluded, rules[updates[i]]...)
	}
	return true
}

func fix_update(updates []string, rules map[string][]string) []string {
	new_pos := make([]map[string]int, 0)
	excluded := make([]string, 0)
	excluded_incr := make([]int, 0)

	for i := 0; i < len(updates); i++ {
		if !slices.Contains(excluded, updates[i]) {
			excluded = append(excluded, rules[updates[i]]...)
			excluded_incr = append(excluded_incr, len(excluded))
		} else {
			// new_i := new_pos[updates[i]]
			var new_i int
			for j := range new_pos {
				var ok bool
				new_i, ok = new_pos[j][updates[i]]
				if ok {
					break
				}
			}
			a := updates[new_i]
			updates[new_i] = updates[i]
			updates[i] = a
			i = new_i
			if i == 0 {
				excluded = make([]string, 0)
				excluded_incr = make([]int, 0)
				new_pos = make([]map[string]int, 0)

			} else {
				excluded = excluded[:excluded_incr[i]]
				excluded_incr = excluded_incr[:i]
				new_pos = new_pos[:i]
			}
			excluded = append(excluded, rules[updates[i]]...)
			excluded_incr = append(excluded_incr, len(excluded))
		}

		new_pos = append(new_pos, make(map[string]int))
		for _, val := range rules[updates[i]] {
			new_pos[i][val] = i
		}
	}
	return updates
}

func main() {
	fmt.Println("Day 5 AoC 2024")
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	buf := bufio.NewReader(file)
	rules := make(map[string][]string)
	updates := make([][]string, 0)
	// err_updates := make([][]string, 0) // part 2

	rules_part := true

	// Parsing data
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}

		if line == "\n" {
			rules_part = false
			continue
		}

		if rules_part {
			parts := strings.Split(line[:len(line)-1], "|")
			_, ok := rules[parts[1]]
			if !ok {
				rules[parts[1]] = make([]string, 0)
			}
			rules[parts[1]] = append(rules[parts[1]], parts[0])
		} else {
			update_list := strings.Split(line[:len(line)-1], ",")
			updates = append(updates, update_list)
		}
	}

	// Solving problem
	start_time := time.Now()
	mid_sum := 0
	fix_mid_sum := 0

	for i := range updates {
		is_valid := check_rules(updates[i], rules)
		// if !is_valid {
		// updates[i] = fix_update(updates[i], rules)
		// }
		if is_valid {
			middle_cell := int(math.Floor(float64((len(updates[i]))) / 2))
			val, err := strconv.Atoi(updates[i][middle_cell])
			if err != nil {
				fmt.Println(err)
				return
			}
			mid_sum += val
		} else {
			// fix_mid_sum += val
		}
	}

	time_taken := time.Since(start_time)
	fmt.Println("Time taken: ", time_taken)

	empty := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	fmt.Println("Sum of middle elements of valid updates: ", mid_sum)
	fmt.Println("Sum of middle elements of fixed updates: ", fix_mid_sum)
	fmt.Println("empty list: ", empty[0:1])
}
