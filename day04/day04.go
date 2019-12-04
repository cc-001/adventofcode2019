package day04

import (
	"fmt"
	"strconv"
)

// note you don't need to atoi these can compare chars instead
func Day04_solve(lower int, upper int, part2 bool) int {
	count := 0
	for i := lower; i < upper; i++ {
		str := fmt.Sprintf("%d", i)
		d := [6]int{}
		d[0], _ = strconv.Atoi(str[0:1])
		d[1], _ = strconv.Atoi(str[1:2])
		if d[1] < d[0] {
			continue
		}
		d[2], _ = strconv.Atoi(str[2:3])
		if d[2] < d[1] {
			continue
		}
		d[3], _ = strconv.Atoi(str[3:4])
		if d[3] < d[2] {
			continue
		}
		d[4], _ = strconv.Atoi(str[4:5])
		if d[4] < d[3] {
			continue
		}
		d[5], _ = strconv.Atoi(str[5:6])
		if d[5] < d[4] {
			continue
		}

		longest_run := 0
		has_double := false
		len_d := len(d)
		for j := 0; j < len_d-1; {
			current_run := 1
			for k := j+1; k < len_d; k++ {
				if d[j] == d[k] {
					current_run++
					j++
				}
			}
			if current_run == 2 {
				has_double = true
			}
			if current_run > longest_run {
				longest_run = current_run
			}
			if current_run == 1 {
				j++
			}
		}

		if part2 {
			if has_double == false {
				continue
			}
		} else if longest_run < 2 {
			continue
		}

		// fmt.Printf("%v - has_double: %v longest_run: %d\n", d, has_double, longest_run)
		count++
	}
	return count
}