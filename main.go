/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mhsantos/advent-2021/days"
	"golang.org/x/exp/slices"
)

func main() {
	args := os.Args
	if !validateArgs(args) {
		printMenu()
	} else {
		if len(args) == 2 {
			runDays("all", "all")
		} else {
			runDays(args[2], args[4])
		}
	}
}

func validateArgs(args []string) bool {
	if len(args) != 2 && len(args) != 5 {
		return false
	}
	if len(args) == 2 {
		if args[1] != "--all" {
			return false
		}
		return true
	}
	if args[1] != "-d" && args[1] != "--day" {
		return false
	}
	if args[3] != "-p" && args[3] != "--part" {
		return false
	}
	if _, err := strconv.Atoi(args[2]); err != nil {
		return false
	}
	if !slices.Contains([]string{"example", "1", "2", "all"}, args[4]) {
		return false
	}
	return true
}

func runDays(day string, part string) {
	funcRegistry := map[string]func(string){
		"Day1": days.Day1,
		"Day2": days.Day2,
		"Day3": days.Day3,
		"Day4": days.Day4,
		"Day5": days.Day5,
		"Day6": days.Day6,
		"Day7": days.Day7,
	}
	if day == "all" {
		for d := 1; d < 8; d++ {
			funcRegistry[fmt.Sprintf("Day%d", d)]("all")
		}
	} else {
		funcRegistry[fmt.Sprintf("Day%s", day)](part)
	}
}

func printMenu() {
	fmt.Println("usage: advent --all -d, --day <day> -p, --part <example|1|2|all>")
	fmt.Println("--all				runs the challenges for all days, all parts")
	fmt.Println("-d, --day day		runs all the parts for the specified challenge day")
	fmt.Println("-p, --part part	runs the specified part for the specified challenge day")
	fmt.Println("						valid options are: example, 1, 2")
}
