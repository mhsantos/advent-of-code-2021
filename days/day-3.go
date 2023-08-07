package days

import (
	"fmt"

	"github.com/mhsantos/advent-2021/internal/utils"
)

func Day3(part string) {
	if part == "example" || part == "all" {
		if lines, err := utils.ReadInputAndConvert("day-3-example.txt", byteArrToBitArr); err == nil {
			mostPresentBits := mostPresentBits(lines)
			leastPresentBits := utils.InvertBits(mostPresentBits)
			powerConsumption := utils.BitsToBase10(mostPresentBits) * utils.BitsToBase10(leastPresentBits)
			fmt.Printf("Day 3 - example: Power consumption: %d\n", powerConsumption)
		}
	}
	if part != "example" {
		if lines, err := utils.ReadInputAndConvert("day-3-1.txt", byteArrToBitArr); err == nil {
			if part == "1" || part == "all" {
				mostPresentBits := mostPresentBits(lines)
				leastPresentBits := utils.InvertBits(mostPresentBits)
				powerConsumption := utils.BitsToBase10(mostPresentBits) * utils.BitsToBase10(leastPresentBits)
				fmt.Printf("Day 3 - Part 1: Power consumption: %d\n", powerConsumption)
			}
			if part == "2" || part == "all" {
				oxygenRating := getElementFromMostCommonBit(lines)
				co2Rating := getElementFromLeastCommonBit(lines)
				lifeSupportRating := utils.BitsToBase10(oxygenRating) * utils.BitsToBase10(co2Rating)
				fmt.Printf("Day 3 - Part 2: Life support rating: %d\n", lifeSupportRating)
			}
		}
	}
}

func byteArrToBitArr(bytes []byte) ([]int, error) {
	var asBinary []int
	for _, bit := range bytes {
		if rune(bit) == '0' {
			asBinary = append(asBinary, 0)
		} else {
			asBinary = append(asBinary, 1)
		}
	}
	return asBinary, nil
}

// takes a matrix of 1's and 0's and returns an integer array where each position contains
// the most common element per column.
// It will therefore return an array with the same number of elements as the columns in the matrix
func mostPresentBits(digits [][]int) []int {
	var result []int
	rows := len(digits)
	cols := len(digits[0])
	for col := 0; col < cols; col++ {
		zeroes, ones := 0, 0
		for row := 0; row < rows; row++ {
			if digits[row][col] == 0 {
				zeroes++
			} else {
				ones++
			}
		}
		if zeroes > ones {
			result = append(result, 0)
		} else {
			result = append(result, 1)
		}
	}
	return result
}

// mostCommonBits iterates over matrix for the column defined in column and rows defined in
// selection for instance if we have a matrix with 10 rows, having selection as [0,1,7] will
// only consider rows 0, 1 and 7 for the position defined in column.
// returns the most column bit (0 or 1) and the row positions that contain it. In case of same
// number of 0's and 1's uses the value defined by draw as most comon
func mostCommonBits(matrix [][]int, selection []int, column int) []int {
	ones := 0
	for _, row := range selection {
		if matrix[row][column] == 1 {
			ones++
		}
	}
	mostCommonIsOne := false
	result := make([]int, 0)
	if ones > len(selection)-ones || ones == len(selection)-ones {
		mostCommonIsOne = true
	}
	for _, row := range selection {
		if mostCommonIsOne && matrix[row][column] == 1 {
			result = append(result, row)
		} else if !mostCommonIsOne && matrix[row][column] == 0 {
			result = append(result, row)
		}
	}
	return result
}

// leastCommonBits is the same as mostCommonBits, but searching for least common bits and when
// in draw considering columns with 0
func leastCommonBits(matrix [][]int, selection []int, column int) []int {
	ones := 0
	for _, row := range selection {
		if matrix[row][column] == 1 {
			ones++
		}
	}
	leastCommonIsZero := false
	result := make([]int, 0)
	if ones > len(selection)-ones || ones == len(selection)-ones {
		leastCommonIsZero = true
	}
	for _, row := range selection {
		if leastCommonIsZero && matrix[row][column] == 0 {
			result = append(result, row)
		} else if !leastCommonIsZero && matrix[row][column] == 1 {
			result = append(result, row)
		}
	}
	return result
}

// core logic for the oxygen generator rating, where it will iterate column over column
// to find the most common bit and filter in the rows where having that bit in that colunm
// returns the final row matching the most common bit
func getElementFromMostCommonBit(matrix [][]int) []int {
	columns := len(matrix[0])
	var rows []int
	for i := 0; i < len(matrix); i++ {
		rows = append(rows, i)
	}
	for i := 0; i < columns; i++ {
		rows = mostCommonBits(matrix, rows, i)
		if len(rows) == 1 {
			return matrix[rows[0]]
		}
	}
	return matrix[rows[0]]
}

func getElementFromLeastCommonBit(matrix [][]int) []int {
	columns := len(matrix[0])
	var rows []int
	for i := 0; i < len(matrix); i++ {
		rows = append(rows, i)
	}
	for i := 0; i < columns; i++ {
		rows = leastCommonBits(matrix, rows, i)
		if len(rows) == 1 {
			return matrix[rows[0]]
		}
	}
	return matrix[rows[0]]
}
