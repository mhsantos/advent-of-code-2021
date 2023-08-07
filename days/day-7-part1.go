package days

import (
	"strconv"
	"strings"

	"github.com/emirpasic/gods/maps/treemap"
	"github.com/mhsantos/advent-2021/internal/utils"
)

type hPosition struct {
	position        int // this position
	crabSubs        int // how many crab subs are in this position
	crabSubsRight   int // total number of crab subs on the right of this position
	crabSubsLeft    int // total number of crab subs on the left of this position
	fuelSubsRight   int // how many fuel units are necessary for all the subs on the right to get here
	fuelSubsLeft    int // how many fuel units are necessary for all the subs on the left to get here
	fuelFactorLeft  int // for Part 2. current fuel calculation factor. = previous fuel factor + current total number of subs
	fuelFactorRight int // current fuel calculation factor. = previous fuel factor + current total number of subs
}

func day7Part1(part string) int {
	var lines []string
	if part == "example" {
		lines = utils.ReadInput("day-7-example.txt")
	} else {
		lines = utils.ReadInput("day-7-1.txt")
	}
	subsPositions := populateHPositions(lines[0])
	calculateDistances(subsPositions)
	return lowestFuelCost(subsPositions)
}

func populateHPositions(line string) *treemap.Map {
	crabSubmarineStrs := strings.Split(line, ",")
	subsMap := treemap.NewWithIntComparator()
	for _, crabSubmarineStr := range crabSubmarineStrs {
		if crabSubmarine, err := strconv.Atoi(crabSubmarineStr); err == nil {
			if _, exists := subsMap.Get(crabSubmarine); !exists {
				subsMap.Put(crabSubmarine, hPosition{position: crabSubmarine})
			}
			positionInterface, _ := subsMap.Get(crabSubmarine)
			position := positionInterface.(hPosition)
			position.crabSubs++
			subsMap.Put(crabSubmarine, position)
		}
	}
	return subsMap
}

func calculateDistances(positions *treemap.Map) {
	keys := positions.Keys()
	if len(keys) < 2 {
		return
	}
	// first iterate from left to right
	for i := 1; i < len(keys); i++ {
		prev, _ := positions.Get(keys[i-1])
		curr, _ := positions.Get(keys[i])
		prevPosition := prev.(hPosition)
		currPosition := curr.(hPosition)
		distanceFromPrev := currPosition.position - prevPosition.position
		crabSubsLeft := prevPosition.crabSubs + prevPosition.crabSubsLeft
		currPosition.crabSubsLeft = crabSubsLeft
		currPosition.fuelSubsLeft = prevPosition.fuelSubsLeft + crabSubsLeft*distanceFromPrev
		positions.Put(keys[i], currPosition)
	}
	// iterate from right to left
	for i := len(keys) - 2; i >= 0; i-- {
		prev, _ := positions.Get(keys[i+1])
		curr, _ := positions.Get(keys[i])
		prevPosition := prev.(hPosition)
		currPosition := curr.(hPosition)
		distanceFromPrev := prevPosition.position - currPosition.position
		crabSubsRight := prevPosition.crabSubs + prevPosition.crabSubsRight
		currPosition.crabSubsRight = crabSubsRight
		currPosition.fuelSubsRight = prevPosition.fuelSubsRight + crabSubsRight*distanceFromPrev
		positions.Put(keys[i], currPosition)
	}
}

func lowestFuelCost(subsPositions *treemap.Map) int {
	lowestCost := 0
	it := subsPositions.Iterator()
	for it.Next() {
		position := it.Value().(hPosition)
		fuelCost := position.fuelSubsLeft + position.fuelSubsRight
		if lowestCost == 0 || fuelCost < lowestCost {
			lowestCost = fuelCost
		}
	}
	return lowestCost
}
