package days

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/mhsantos/advent-2021/internal/utils"
)

func Day8(part string) {
	if part == "example" || part == "all" {
		lines := utils.ReadInput("day-8-example.txt")
		fmt.Printf("Day 8 - example: Unique digits: %d\n", countUniqueDigits(lines))
	}
	if part == "1" || part == "all" {
		lines := utils.ReadInput("day-8-1.txt")
		fmt.Printf("Day 8 - Part 1: Unique digits: %d\n", countUniqueDigits(lines))
	}
	if part == "2" || part == "all" {
		lines := utils.ReadInput("day-8-1.txt")
		fmt.Printf("Day 8 - Part 2: Sum: %d\n", calculateSum(lines))
	}
}

func lineToDigits(line string) []string {
	return strings.Split(line, " ")
}

// counts the digits 1,4,7 and 8, which have unique numbers of segments
func countUniqueDigits(lines []string) int {
	count := 0
	for _, line := range lines {
		digits := lineToDigits(strings.Split(line, " | ")[1])
		for _, digit := range digits {
			if len(digit) == 2 || len(digit) == 3 || len(digit) == 4 || len(digit) == 7 {
				count++
			}
		}

	}
	return count
}

func calculateSum(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum += getNumber(line)
	}
	return sum
}

func getNumber(line string) int {
	parts := strings.Split(line, " | ")
	patterns := strings.Split(parts[0], " ")
	digits := identifyDigits(patterns)
	return identifyNumber(digits, strings.Split(parts[1], " "))
}

func identifyNumber(patterns [10]string, digits []string) int {
	result := 0
	N := len(digits)
	for i := 0; i < N; i++ {
		switch len(digits[i]) {
		case 2:
			result += int(math.Pow10(N - 1 - i))
		case 3:
			result += 7 * int(math.Pow10(N-1-i))
		case 4:
			result += 4 * int(math.Pow10(N-1-i))
		case 7:
			result += 8 * int(math.Pow10(N-1-i))
		case 5, 6:
			sortedDigit := sortString(digits[i])
			if len(digits[i]) == 5 { // 2, 3 or 5
				if patterns[2] == sortedDigit {
					result += 2 * int(math.Pow10(N-1-i))
				} else if patterns[5] == sortedDigit {
					result += 5 * int(math.Pow10(N-1-i))
				} else {
					result += 3 * int(math.Pow10(N-1-i))
				}
			} else { // 0, 6 or 9
				if patterns[6] == sortedDigit {
					result += 6 * int(math.Pow10(N-1-i))
				} else if patterns[9] == sortedDigit {
					result += 9 * int(math.Pow10(N-1-i))
				}
				// 0 doesn't add anything so we won't consider it
			}
		}
	}
	return result
}

func identifyDigits(patterns []string) [10]string {
	var digits [10]string = [10]string{}
	for i := 0; i < 10; i++ {
		patterns[i] = sortString(patterns[i])
	}
	for _, pattern := range patterns {
		switch len(pattern) {
		case 2:
			digits[1] = pattern
		case 3:
			digits[7] = pattern
		case 4:
			digits[4] = pattern
		case 7:
			digits[8] = pattern
		default:
			continue
		}
	}
	for _, pattern := range patterns {
		oneSegments := []rune(digits[1])
		var fmos []rune // fourMinusOneSegments: the 2 segments in 4 that are not in 1: |_
		for i := 0; i < 4; i++ {
			if !strings.ContainsRune(digits[1], rune(digits[4][i])) {
				fmos = append(fmos, rune(digits[4][i]))
			}
		}
		switch len(pattern) {
		case 2, 3, 4, 7: // digits 1, 4, 7 or 8
			continue
		case 5: // digits 2, 3 or 5
			if strings.ContainsRune(pattern, oneSegments[0]) && strings.ContainsRune(pattern, oneSegments[1]) {
				digits[3] = pattern
			} else if strings.ContainsRune(pattern, fmos[0]) && strings.ContainsRune(pattern, fmos[1]) {
				digits[5] = pattern
			} else {
				digits[2] = pattern
			}

		case 6: // digits 0, 6 or 9
			if !strings.ContainsRune(pattern, fmos[0]) || !strings.ContainsRune(pattern, fmos[1]) {
				// the digit has 6 segments the middle horizontal segment is off so it can only be 0
				digits[0] = pattern
			} else if strings.ContainsRune(pattern, oneSegments[0]) && strings.ContainsRune(pattern, oneSegments[1]) {
				// both segments in 1 are on, so it can only be 9
				digits[9] = pattern
			} else {
				digits[6] = pattern
			}
		}
	}
	return digits
}

func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
