package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

// 89013607104091 too high
// 89159109749791

func main() {
	load_time := time.Now()
	fmt.Println("AoC 2024 - Day 10")
	data := load_data()
	fmt.Println("Load time: ", time.Since(load_time))
	start_time := time.Now()
	tokens := int64(0)

	for _, m := range data {
		a, b := solve_machine(m)
		if (a*m.a_button_x+b*m.b_button_x) == m.prize_x && (a*m.a_button_y+b*m.b_button_y) == m.prize_y {
			tokens += 3*a + b
		}
	}

	fmt.Println("Tokens: ", tokens)
	fmt.Println("Part 1 time: ", time.Since(start_time))
	part2_time := time.Now()

  tokens = 0
	offset := int64(10000000000000)
	for _, m := range data {
		m.prize_x += offset
		m.prize_y += offset

		// fmt.Println("Button A: X+", m.a_button_x, "Y+", m.a_button_y)
		// fmt.Println("Button B: X+", m.b_button_x, "Y+", m.b_button_y)
		// fmt.Println("Prize: X=", m.prize_x, "Y=", m.prize_y)
		// fmt.Println("")
		a, b := solve_machine(m)
		// fmt.Println("A=", a, "B=", b)

		if (a*m.a_button_x+b*m.b_button_x) == m.prize_x && (a*m.a_button_y+b*m.b_button_y) == m.prize_y {
			tokens += 3*a + b
		}
	}
	fmt.Println("Tokens: ", tokens)
	fmt.Println("Part 2 time: ", time.Since(part2_time))
	fmt.Println("Total time: ", time.Since(load_time))
}

type Machine struct {
	a_button_x int64
	a_button_y int64
	b_button_x int64
	b_button_y int64
	prize_x    int64
	prize_y    int64
}

func solve_machine(m Machine) (int64, int64) {
	// Cramers rule
	//   A   B
	// (a_x b_x) = (prize_x)
	// (a_y b_y) = (prize_y)
	det := m.a_button_x*m.b_button_y - m.a_button_y*m.b_button_x
	a := (m.prize_x*m.b_button_y - m.prize_y*m.b_button_x) / det
	b := (m.a_button_x*m.prize_y - m.a_button_y*m.prize_x) / det
	return a, b
}

func load_data() []Machine {
	data := make([]Machine, 0)
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	buf := bufio.NewReader(file)
	m := Machine{}
	c := 0
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
		switch c {
		case 0:
			a_x, _ := strconv.Atoi(line[12:14])
			a_y, _ := strconv.Atoi(line[18:20])
			m.a_button_x = int64(a_x)
			m.a_button_y = int64(a_y)
			c++
		case 1:
			b_x, _ := strconv.Atoi(line[12:14])
			b_y, _ := strconv.Atoi(line[18:20])
			m.b_button_x = int64(b_x)
			m.b_button_y = int64(b_y)
			c++
		case 2:
			x_a, x_b, y_a := 0, 0, 0
			for i, c := range line {
				if c == '=' {
					if x_a == 0 {
						x_a = i + 1
					} else {
						y_a = i + 1
						break
					}
				}
				if c == ',' {
					x_b = i
				}
			}

			x, _ := strconv.Atoi(line[x_a:x_b])
			y, _ := strconv.Atoi(line[y_a:len(line)])
			m.prize_x = int64(x)
			m.prize_y = int64(y)
			c++
		case 3:
			data = append(data, m)
			m = Machine{}
			c = 0
		}
	}

	return data
}
