package main

import (
	"aoc2024/day1"
	"os"
)

func main() {
	runner := os.Args[1]

	switch runner {
	case "d1p1":
		day1.RunPart1("./inputs/locationLists.txt")
	case "d1p2":
		day1.RunPart2("./inputs/locationLists.txt")
	}
}
