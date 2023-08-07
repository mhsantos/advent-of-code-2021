// day-1-1.go solves the first part of day 1 of Advent of Code 2021
// https://adventofcode.com/2021/day/1
package days

import (
	"fmt"
	"strconv"

	"github.com/mhsantos/advent-2021/internal/utils"
)

func Day1(part string) {
	if part == "example" || part == "all" {
		lines := utils.ReadInput("day-1-example.txt")
		fmt.Printf("Day 1 - example: Increases: %d\n", calculateIncreases(lines))
	}
	if part == "1" || part == "all" {
		lines := utils.ReadInput("day-1-1.txt")
		fmt.Printf("Day 1 - Part 1: Increases: %d\n", calculateIncreases(lines))
	}
	if part == "2" || part == "all" {
		intLines, err := utils.ReadInputAndConvert("day-1-1.txt", byteArrToInt)
		if err == nil {
			fmt.Printf("Day 1 - Part 2: Rolling increases: %d\n", calculateRollingWindowIncreases(intLines))
		}
	}
}

// Part 1
func calculateIncreases(lines []string) int {
	var increases, prev, depth int
	var err error
	for idx, value := range lines {
		if depth, err = strconv.Atoi(value); err != nil {
			break
		}
		if idx > 0 {
			if depth > prev {
				increases++
			}
		}
		prev = depth
	}
	return increases
}

func byteArrToInt(bytes []byte) (int, error) {
	return strconv.Atoi(string(bytes[:]))
}

// Part 2
func calculateRollingWindowIncreases(lines []int) int {
	N := len(lines)
	increases := 0
	if N < 3 {
		return 0
	}
	curr := 0
	for _, value := range lines[0:3] {
		curr += value
	}
	for idx, value := range lines[3:] {
		prev := curr
		curr += value
		curr -= lines[idx]
		if curr > prev {
			increases++
		}
	}
	return increases

}
