package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadReportsFromFile(path string) [][]int {
	reports := [][]int{}

	// read the input file
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		report := []int{}
		values := strings.Split(scanner.Text(), " ")
		for _, v := range values {
			intVal, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			report = append(report, intVal)
		}

		reports = append(reports, report)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return reports
}

func RunPart1(inputPath string) {
	reports := ReadReportsFromFile(inputPath)
	fmt.Println(NoOfSafeReports(reports, false))
}

func RunPart2(inputPath string) {
	reports := ReadReportsFromFile(inputPath)
	fmt.Println(NoOfSafeReports(reports, true))
}
