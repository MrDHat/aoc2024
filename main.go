package main

import (
	"aoc2024/day1"
	"aoc2024/day2"
	"aoc2024/day3"
	"os"
)

func main() {
	runner := os.Args[1]

	switch runner {
	case "d1p1":
		day1.RunPart1("./inputs/locationLists.txt")
	case "d1p2":
		day1.RunPart2("./inputs/locationLists.txt")

	case "d2p1":
		day2.RunPart1("./inputs/reactorReports.txt")
	case "d2p2":
		day2.RunPart2("./inputs/reactorReports.txt")

	case "d3p1":
		day3.RunPart1("./inputs/mullItOver.txt")
	case "d3p2":
		day3.RunPart2("./inputs/mullItOver.txt")
	}
}
