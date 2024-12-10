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

	eqs := make(map[uint64][][]uint64)
	for {
		line, _, err := buf.ReadLine()
		if err == io.EOF {
			break
		}
		// fmt.Println("Line:", string(line))
		parts := strings.Split(string(line), ":")
		nums := strings.Split(parts[1], " ")
		k, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if _, ok := eqs[uint64(k)]; !ok {
			eqs[uint64(k)] = make([][]uint64, 0)
		}
    eqs[uint64(k)] = append(eqs[uint64(k)], make([]uint64, 0))
		eq_pos := len(eqs[uint64(k)]) - 1
		for _, n := range nums {
			if n != "" {
				// v, err := strconv.Atoi(n)
				v, err := strconv.ParseInt(n, 10, 64)
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
				eqs[uint64(k)][eq_pos] = append(eqs[uint64(k)][eq_pos], uint64(v))
			}
		}
	}

	// fmt.Println("Equations:", eqs)
	fmt.Println("Load time:", time.Since(load_time))
	start := time.Now()

	calibration := uint64(0)
	for k, v := range eqs {
		for _, eq := range v {
			res := try_calibrate(k, eq)
			if res {
				calibration += k
			}
		}
	}

	fmt.Println("Part 1 time:", time.Since(start))
	fmt.Println("Calibration part 1:", calibration) // 7885693428401
	fmt.Println("")

	start_2 := time.Now()

	calibration = uint64(0)
	for k, v := range eqs {
		for _, eq := range v {
			res := try_calibrate2(k, eq)
			if res {
				calibration += k
			}
		}
	}

	fmt.Println("Part 2 time:", time.Since(start_2))
	fmt.Println("Calibration part 2:", calibration)


}

func concat_nums(num_a uint64, num_b uint64) uint64 {
	b_len := 1
  res := num_b
	for num_b >= 10 {
		num_b /= 10
		b_len++
	}
	for i := 0; i < b_len; i++ {
		num_a *= 10
	}
	return num_a + res
}

func try_calibrate(k uint64, v []uint64) bool {
	if len(v) == 0 || v[0] > k {
		return false
	}
	return try_calibrate_iter(k, v[0], v[1:])
}

func try_calibrate_iter(k uint64, first uint64, v []uint64) bool {
	if k < first {
		return false
	}
	if len(v) == 0 {
		if k == first {
			return true
		}
		return false
	}

	return try_calibrate_iter(k, first*v[0], v[1:]) || try_calibrate_iter(k, first+v[0], v[1:])
}

func try_calibrate2(k uint64, v []uint64) bool {
	if len(v) == 0 || v[0] > k {
		return false
	}
	return try_calibrate_iter2(k, v[0], v[1:])
}

func try_calibrate_iter2(k uint64, first uint64, v []uint64) bool {
	if k < first {
		return false
	}
	if len(v) == 0 {
		if k == first {
			return true
		}
		return false
	}

	return try_calibrate_iter2(k, first*v[0], v[1:]) || try_calibrate_iter2(k, first+v[0], v[1:]) || try_calibrate_iter2(k, concat_nums(first, v[0]), v[1:])
}
