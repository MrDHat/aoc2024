package day3

import (
	"fmt"
	"os"
)

func ReadDatFromFile(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}

	return string(file)
}

func RunPart1(path string) {
	input := ReadDatFromFile(path)
	Calculate(input)
}
