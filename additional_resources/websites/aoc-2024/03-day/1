package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func operate_mul(vals string) int {
	parts := strings.Split(vals, ",")

	// If there are not two parts, then it is not a valid mul
	if len(parts) != 2 {
		return 0
	}

	// If there are any spaces in the parts, then it is not a valid mul
	if len(strings.TrimSpace(parts[0])) != len(parts[0]) || len(strings.TrimSpace(parts[1])) != len(parts[1]) {
		return 0
	}

	val_1, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0
	}
	val_2, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0
	}

	return val_1 * val_2
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	sum := 0

	buf := bufio.NewReader(file)
	enabled := true
	for {
		line, err := buf.ReadString('\n')
		if err == io.EOF {
			break
		}

		for cur := 0; cur < len(line)-8; {
			fmt.Println("Line max: ", len(line), "Cur: ", cur, "Line - 8: ", len(line)-8)
			if enabled {
				if line[cur:cur+7] == "don't()" {
					enabled = false
					cur += 7
					continue
				}
				if line[cur:cur+4] == "mul(" {
					fmt.Println("Found mulc")
					for f_curr := cur + 5; f_curr < len(line) && f_curr <= cur+12; f_curr++ {
						if line[f_curr] == ')' {
							fmt.Println("Between ", cur, f_curr, line[cur+4:f_curr])
							fmt.Println("Number: ", line[cur+4:f_curr])
							sum += operate_mul(line[cur+4 : f_curr])
							cur = f_curr
							break
						}
					}
				}
			} else {
				if line[cur:cur+4] == "do()" {
					enabled = true
					cur += 4
				}
			}
			cur++
		}
	}

	fmt.Println("Sum: ", sum)
}


// 81458033 too high
// first line:  13189667
