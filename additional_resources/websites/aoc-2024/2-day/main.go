package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Option struct {
	Safe   bool
	Length int
	Root   int
	Leaf   int
	Asc    bool
	Keep   bool
}

func check_rules(root int, leaf int, is_asc bool) bool {
	diff := root - leaf
	if (is_asc == (leaf > root)) && diff >= -3 && diff <= 3 && diff != 0 {
		return true
	}
	return false
}

func check_branch(opt Option, new_node int) []Option {
	safe := opt.Safe
	length := opt.Length
	root := opt.Root
	leaf := opt.Leaf
	is_asc := opt.Asc
	res := []Option{}

	if length == 1 {
		is_asc := new_node > root
		if check_rules(root, new_node, is_asc) {
			res = append(res, Option{safe, 2, root, new_node, is_asc, true})
		} else if safe {
			res = append(res, Option{false, 1, root, 0, true, true})
			res = append(res, Option{false, 1, new_node, 0, true, true})
		}
	}

	if length > 1 {
		if check_rules(leaf, new_node, is_asc) {
			return []Option{{safe, length + 1, leaf, new_node, is_asc, true}}
		} else if safe {
			res = append(res, Option{false, length, root, leaf, is_asc, true})
			if check_rules(root, new_node, is_asc) {
				res = append(res, Option{false, length, root, new_node, is_asc, true})
			}
			if length == 2 {
				if check_rules(root, new_node, new_node > root) {
					res = append(res, Option{false, length, root, new_node, new_node > root, true})
				}
				if check_rules(leaf, new_node, new_node > leaf) {
					res = append(res, Option{false, length, leaf, new_node, new_node > leaf, true})
				}
			}
		}
	}

	if len(res) == 0 {
		res = append(res, Option{false, 0, 0, 0, false, false})
	}
	return res
}

func try_parts(i int, num int, opts []Option) []Option {
	new_opts := make([]Option, 0)
	if i == 0 {
		// new_opts = append(opts, [][]int{{0, 1, num, 0}}...)
		new_opts = append(opts, Option{true, 1, num, 0, true, true})
	} else {
		for _, opt := range opts {
			new_opts = append(new_opts, check_branch(opt, num)...)
		}
	}

	opts = make([]Option, 0)
	for _, new_opt := range new_opts {
		if new_opt.Keep {
			opts = append(opts, new_opt)
		}
	}

	return opts
}

func main() {
	file, err := os.Open("input.csv")
	if err != nil {
		return
	}
	defer file.Close()

	buf := bufio.NewReader(file)
	safe_reports := 0
	super_safe_reports := 0
	// dampened_safe_reports := 0
	test_safe_reports := make([]int, 0)

	for {
		line, err := buf.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}

		parts := strings.Split(line[0:len(line)-1], " ")
		opts := make([]Option, 0)

		for i, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				continue
			}
			opts = try_parts(i, num, opts)
		}

		if len(opts) > 0 {
			test_safe_reports = append(test_safe_reports, 1)
			safe_reports += 1
			for _, opt := range opts {
				if opt.Safe {
					super_safe_reports += 1
					break
				}
			}
		} else {
			test_safe_reports = append(test_safe_reports, 0)
		}
	}

	fmt.Println("Safe reports: ", safe_reports)
	fmt.Println("Super safe reports: ", super_safe_reports)
}
