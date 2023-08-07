package days

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/mhsantos/advent-2021/internal/utils"
)

func Day7(part string) {
	if part == "example" || part == "all" {
		fmt.Printf("Day 7 - example: Fuel cost: %d\n", day7Part1("example"))
	}
	if part == "1" || part == "all" {
		fmt.Printf("Day 7 - Part 1: Fuel cost: %d\n", day7Part1("1"))
	}
	if part == "2" || part == "all" {
		lines := utils.ReadInput("day-7-1.txt")
		subsPositions := populatePositions(lines[0])
		fuelCosts := calculateDistancesExponentially(subsPositions)
		fmt.Printf("Day 7 - Part 2 Fuel cost: %d\n", findLowestFuelCost(fuelCosts))
	}

}

// populatePositions parses the input file and returns a list of arrays with 2 positions
// positon 0 corresponds to the position in the map, position 1 contains the number of
// submarines in that position. As an example, if we had the input 0,3,2,0,5,3,7,2,7,9, the
// process input function would return the following output:
// [[0][2],[2][2],[3][2],[5][1],[7][2],[9][1]]
func populatePositions(line string) [][2]int {
	crabSubmarineStrs := strings.Split(line, ",")
	subsMap := make(map[int]int)
	for _, crabSubmarineStr := range crabSubmarineStrs {
		if crabSubmarine, err := strconv.Atoi(crabSubmarineStr); err == nil {
			if _, exists := subsMap[crabSubmarine]; !exists {
				subsMap[crabSubmarine] = 0
			}
			subsMap[crabSubmarine]++
		}
	}
	positions := [][2]int{}
	for key, val := range subsMap {
		positions = append(positions, [2]int{key, val})
	}
	sort.Slice(positions, func(i, j int) bool {
		return positions[i][0] < positions[j][0]
	})
	return positions
}

// calculateDistancesExponentially takes a list of 2 positions as argument. For each element
// in that list, position 0 is the current position in the map and position 1 has the number of
// submarines in that position. The function returns another list where each position has two
// elements. Position 0 is the position in the map and position 2 means the cost for the neighbors
// to get to that position
func calculateDistancesExponentially(positions [][2]int) [][2]int {
	if len(positions) < 2 {
		return nil
	}

	firstPosition := positions[0][0]
	currentPosition := firstPosition
	lastPosition := positions[len(positions)-1][0]
	result := make([][2]int, 0)
	currentSubs, currentFuelCost, currentFuelFactor := 0, 0, 0

	// first calculates fuel cost from left to right
	for _, position := range positions {
		for currentPosition < position[0] {
			result = append(result, [2]int{currentPosition, currentFuelCost})
			currentFuelFactor += currentSubs
			currentFuelCost += currentFuelFactor
			currentPosition++
		}
		result = append(result, [2]int{currentPosition, currentFuelCost})
		currentSubs += position[1]
		currentFuelFactor += currentSubs
		currentFuelCost += currentFuelFactor
		currentPosition++
	}

	// then calculates from right to left
	currentSubs, currentFuelCost, currentFuelFactor = 0, 0, 0
	currentPosition = lastPosition
	currentResultPosition := len(result) - 1
	for i := len(positions) - 1; i >= 0; i-- {
		position := positions[i]
		for currentPosition > position[0] {
			result[currentResultPosition][1] += currentFuelCost
			currentFuelFactor += currentSubs
			currentFuelCost += currentFuelFactor
			currentPosition--
			currentResultPosition--
		}
		result[currentResultPosition][1] += currentFuelCost
		currentSubs += position[1]
		currentFuelFactor += currentSubs
		currentFuelCost += currentFuelFactor
		currentPosition--
		currentResultPosition--
	}

	return result
}

func findLowestFuelCost(positions [][2]int) int {
	lowestCost := 0
	for _, val := range positions {
		if lowestCost == 0 || val[1] < lowestCost {
			lowestCost = val[1]
		}
	}
	return lowestCost
}
