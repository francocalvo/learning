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
		return
	}

	sum := 0

	buf := bufio.NewReader(file)
	// It starts enabled
	enabled := true
	for {
		// Iterate over each line of the input file
		line, err := buf.ReadString('\n')
		if err == io.EOF {
			break
		}

		for cur := 0; cur < len(line)-8; {
			// For each line, iterate over each character.
			if enabled {
				// If we are enabled, we can check for the don't() and mul() functions
				if line[cur:cur+7] == "don't()" {
					// If we find the don't() function, we disable the enabled flag and skip 7 characters
					enabled = false
					cur += 7
					continue
				}
				if line[cur:cur+4] == "mul(" {
					// If we find the mul() function, we iterate over the next 12 characters to find the closing parenthesis
					advanced := false
					for f_curr := cur + 5; f_curr < len(line) && f_curr <= cur+12; f_curr++ {
						if line[f_curr] == ')' {
							sum += operate_mul(line[cur+4 : f_curr])
							cur = f_curr
							advanced = true
							break
						}
					}
					if !advanced {
						// If we didn't find the closing parenthesis skip 4 characters
						cur += 4
					}
				}
			} else {
				// if we are disabled, we only check for the do() function
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
