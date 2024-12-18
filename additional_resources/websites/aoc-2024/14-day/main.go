package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

// 89013607104091 too high
// 89159109749791

func main() {
	load_time := time.Now()
	fmt.Println("AoC 2024 - Day 10")
	data := load_data()
	// h := 7
	// w := 11
	h := int64(103)
	w := int64(101)

	fmt.Println("Load time: ", time.Since(load_time))
	start_time := time.Now()

	final_positions := make([]Point, len(data))
	q_1, q_2, q_3, q_4 := 0, 0, 0, 0
	q_1_x := int64(w/2) - 1
	q_1_y := int64(h/2) - 1
	q_2_x := int64(w/2+2) - 1
	q_2_y := int64(h/2) - 1
	q_3_x := int64(w/2) - 1
	q_3_y := int64(h/2+2) - 1
	q_4_x := int64(w/2+2) - 1
	q_4_y := int64(h/2+2) - 1

	for i := range data {
		final_positions[i] = update_position(data[i], w, h, 100)
		if final_positions[i].a_x <= q_1_x && final_positions[i].a_y <= q_1_y {
			q_1++
		} else if final_positions[i].a_x >= q_2_x && final_positions[i].a_y <= q_2_y {
			q_2++
		} else if final_positions[i].a_x <= q_3_x && final_positions[i].a_y >= q_3_y {
			q_3++
		} else if final_positions[i].a_x >= q_4_x && final_positions[i].a_y >= q_4_y {
			q_4++
		}
	}

	fmt.Println("Safety score: ", q_1*q_2*q_3*q_4)
	fmt.Println("Part 1 time: ", time.Since(start_time))
  part2_time := time.Now()

  // Simple but long heuristic I found in Reddit. 
  // If there is an imagen forming, there are going to be more clusters of 
  // points in the grid and the secuirity score will be lower
  // So, we can try to find the minimum security score by iterating over a 
  // range of steps and checking the security score. The minimum score will be
  // the one that forms the image
	min_sq := int64(q_1 * q_2 * q_3 * q_4)
	var min_sq_steps int64 = 100
	for j := 0; j < 10000; j++ {
		final_positions := make([]Point, len(data))
		var q_1, q_2, q_3, q_4 int64 = 0, 0, 0, 0
		for i := range data {
			final_positions[i] = update_position(data[i], w, h, int64(j))

			if final_positions[i].a_x <= q_1_x && final_positions[i].a_y <= q_1_y {
				q_1++
			} else if final_positions[i].a_x >= q_2_x && final_positions[i].a_y <= q_2_y {
				q_2++
			} else if final_positions[i].a_x <= q_3_x && final_positions[i].a_y >= q_3_y {
				q_3++
			} else if final_positions[i].a_x >= q_4_x && final_positions[i].a_y >= q_4_y {
				q_4++
			}
		}
		if q_1*q_2*q_3*q_4 < min_sq {
			min_sq = q_1 * q_2 * q_3 * q_4
			min_sq_steps = int64(j)
		}
	}

  fmt.Println("Part 2: ", min_sq_steps)
  fmt.Println("Part 2 time: ", time.Since(part2_time))
	fmt.Println("Total time: ", time.Since(load_time))

}

func update_position(p Point, w, h, secs int64) Point {
	p.a_x += p.v_x * secs
	p.a_y += p.v_y * secs

	for p.a_x < 0 {
		p.a_x += w
	}

	for p.a_y < 0 {
		p.a_y += h
	}

	for p.a_x >= w {
		p.a_x -= w
	}

	for p.a_y >= h {
		p.a_y -= h
	}

	return p
}

type Point struct {
	a_x, a_y, v_x, v_y int64
}

func load_data() []Point {
	data := make([]Point, 0)
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
		point := Point{}
		line = line[:len(line)-1]
		parts_a := strings.Split(line, "=")

		p_parts := strings.Split(parts_a[1], ",")
		point.a_x, _ = strconv.ParseInt(p_parts[0], 10, 64)
		point.a_y, _ = strconv.ParseInt(strings.Split(p_parts[1], " ")[0], 10, 64)

		v_parts := strings.Split(parts_a[2], ",")
		point.v_x, _ = strconv.ParseInt(v_parts[0], 10, 64)
		point.v_y, _ = strconv.ParseInt(v_parts[1], 10, 64)

		data = append(data, point)
	}

	return data
}
