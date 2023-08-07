package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mhsantos/advent-2021/internal/utils"
)

func Day6(part string) {
	if part == "example" || part == "all" {
		lines := utils.ReadInput("day-6-example.txt")
		fishesTimer := createTimer(lines[0])
		fmt.Printf("Day 6 - example: Fishes after %d days: %d\n", 80, iterateDays(fishesTimer, 80))
	}
	if part != "example" {
		lines := utils.ReadInput("day-6-1.txt")
		fishesTimer := createTimer(lines[0])
		if part == "1" || part == "all" {
			fmt.Printf("Day 6 - Part 1: Fishes after %d days: %d\n", 80, iterateDays(fishesTimer, 80))
		}
		if part == "2" {
			fmt.Printf("Day 6 - Part 2: Fishes after %d days: %d\n", 256, iterateDays(fishesTimer, 256))
		}
		if part == "all" {
			fishesTimer = createTimer(lines[0])
			fmt.Printf("Day 6 - Part 2: Fishes after %d days: %d\n", 256, iterateDays(fishesTimer, 256))
		}
	}
}

func createTimer(line string) [9]int {
	fishesTimer := [9]int{}
	fishTimerStr := strings.Split(line, ",")
	for _, timerStr := range fishTimerStr {
		if timer, err := strconv.Atoi(timerStr); err == nil {
			fishesTimer[timer]++
		}
	}
	return fishesTimer
}

// iterateDays runs days iterations over fishesTimer, simulating the passage of a day. For each
// iteration it shift all numbers to the left, simulating that a day has passed. It then will
// add  will copy the previous value of position 0 to position 8, to simulate new fishes. It will
// add that number to position 6, to simulate that those fishes will continue producing fishes.
func iterateDays(fishesTimer [9]int, days int) int {
	for day := 0; day < days; day++ {
		prevPos0 := fishesTimer[0]
		for i := 0; i < 8; i++ {
			fishesTimer[i] = fishesTimer[i+1]
		}
		fishesTimer[8] = prevPos0
		fishesTimer[6] += prevPos0
	}
	fishes := 0
	for _, val := range fishesTimer {
		fishes += val
	}
	return fishes
}
