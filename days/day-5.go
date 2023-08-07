package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mhsantos/advent-2021/internal/utils"
)

func Day5(part string) {
	// for Part1, pass allowsDiagonal = false to populateMap
	if part == "example" || part == "all" {
		lines := utils.ReadInput("day-5-example.txt")
		if coordinates, dimensions, err := readCoordinates(lines); err == nil {
			linesMap := populateMap(coordinates, dimensions, false)
			fmt.Printf("Day 5 - example: Points overlapping: %d\n",
				countPointsWithMultipleLines(linesMap))
		}
	}
	if part != "example" {
		lines := utils.ReadInput("day-5-1.txt")
		if coordinates, dimensions, err := readCoordinates(lines); err == nil {
			if part == "1" || part == "all" {
				linesMap := populateMap(coordinates, dimensions, false)
				fmt.Printf("Day 5 - Part 1: Points overlapping: %d\n",
					countPointsWithMultipleLines(linesMap))
			}
			if part == "2" || part == "all" {
				linesMap := populateMap(coordinates, dimensions, true)
				fmt.Printf("Day 5 - Part 2: Points overlapping: %d\n",
					countPointsWithMultipleLines(linesMap))
			}
		}
	}
}

// countPointsWithMultipleLines iterates over the map points and returns the number of points
// in the map where more than one line passes over
func countPointsWithMultipleLines(linesMap [][]int) int {
	multipleLines := 0
	for i := range linesMap {
		for j := range linesMap[i] {
			if linesMap[i][j] > 1 {
				multipleLines++
			}
		}
	}
	return multipleLines
}

// printMap prints the map showing the lines from the input coordinates to show an output
// similar to the one shown in the problem statement
func printMap(linesMap [][]int) {
	cols := len(linesMap)
	rows := len(linesMap[0])

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if linesMap[col][row] == 0 {
				fmt.Print(" .")
			} else {
				fmt.Print(" ", linesMap[col][row])
			}
		}
		fmt.Println("")
	}
}

// populate map creates a map of cols x rows dimensions[0] x dimensions[1]
// where each position in the map indicates how many lines defined by coordinates
// pass by. ignores diagonal lines if allowsDiagonal is false
func populateMap(coordinates [][2][2]int, dimensions [2]int, allowsDiagonal bool) [][]int {
	linesMap := make([][]int, dimensions[0])
	for idx := range linesMap {
		linesMap[idx] = make([]int, dimensions[1])
	}
	for _, line := range coordinates {
		from := line[0]
		to := line[1]
		if from[0] == to[0] {
			// both from and to x positions are the same, so it is a vertical line
			if from[1] < to[1] {
				for i := from[1]; i <= to[1]; i++ {
					linesMap[from[0]][i]++
				}
			} else {
				for i := to[1]; i <= from[1]; i++ {
					linesMap[from[0]][i]++
				}
			}
		} else if from[1] == to[1] {
			// both from and to y positions are the same, so it is a horizontal line
			if from[0] < to[0] {
				for i := from[0]; i <= to[0]; i++ {
					linesMap[i][from[1]]++
				}
			} else {
				for i := to[0]; i <= from[0]; i++ {
					linesMap[i][from[1]]++
				}
			}
		} else if allowsDiagonal {
			// neither from x to to x or from y to to y are the same, so it is a diagonal line.
			// as from the problem definition, diagonal lines can only be in an 40 degrees angle
			// so every point in the line will have x and y incremented by 1

			// to dine how we will iterate over the points in the line, we need to determine
			// the direction for the from -> to line. It can be NE, NW, SE or SW:
			// NE: for each point x++ y--
			// NW: for each point x-- y--
			// SE: for each point x++ y++
			// SW: for each point x-- y++
			if from[0] < to[0] { // ?E
				if from[1] < to[1] { // SE
					for x, y := from[0], from[1]; x <= to[0]; x++ {
						linesMap[x][y]++
						y++
					}
				} else { // NE
					for x, y := from[0], from[1]; x <= to[0]; x++ {
						linesMap[x][y]++
						y--
					}
				}
			} else { // ?W
				if from[1] < to[1] { // SW
					for x, y := from[0], from[1]; x >= to[0]; x-- {
						linesMap[x][y]++
						y++
					}
				} else { // NW
					for x, y := from[0], from[1]; x >= to[0]; x-- {
						linesMap[x][y]++
						y--
					}
				}
			}
		}
	}
	return linesMap
}

// readsCoordinates parses the lines from the input file and return a slice of bi-dimensional
// int arrays, with each slice position containins a 2x2 matrix, were the first position is
// the from integer pair and the second is the to integer pair. It also returns an 2 positions
// array with the dimension (cols and rows) for an eventual map. The input line 0,1 -> 5,3,
// for instance, would return [][2][2]int{[2]int{0,1}, [2]int{5,3}} and dimensions 6,4.
func readCoordinates(input []string) ([][2][2]int, [2]int, error) {
	var lines [][2][2]int
	dimensions := [2]int{}
	for _, line := range input {
		params := strings.Split(line, " -> ")
		fromStr := strings.Split(params[0], ",")
		toStr := strings.Split(params[1], ",")
		from := [2]int{}
		to := [2]int{}
		var err error
		if from[0], err = strconv.Atoi(fromStr[0]); err != nil {
			return lines, dimensions, err
		}
		if from[1], err = strconv.Atoi(fromStr[1]); err != nil {
			return lines, dimensions, err
		}
		if to[0], err = strconv.Atoi(toStr[0]); err != nil {
			return lines, dimensions, err
		}
		if to[1], err = strconv.Atoi(toStr[1]); err != nil {
			return lines, dimensions, err
		}
		coordinates := [2][2]int{from, to}
		lines = append(lines, coordinates)
		if from[0] > dimensions[0] {
			dimensions[0] = from[0]
		}
		if to[0] > dimensions[0] {
			dimensions[0] = to[0]
		}
		if from[1] > dimensions[1] {
			dimensions[1] = from[1]
		}
		if to[1] > dimensions[1] {
			dimensions[1] = to[1]
		}
	}
	dimensions[0]++
	dimensions[1]++
	return lines, dimensions, nil
}
