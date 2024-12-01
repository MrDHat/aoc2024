package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadListsFromFile(path string) ([]int, []int) {
	list1 := []int{}
	list2 := []int{}

	// read the input file
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), "   ")
		val1, err := strconv.Atoi(values[0])
		if err != nil {
			log.Fatal(err)
		}
		list1 = append(list1, val1)

		val2, err := strconv.Atoi(values[1])
		if err != nil {
			log.Fatal(err)
		}
		list2 = append(list2, val2)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return list1, list2
}

func RunPart1(inputPath string) {
	list1, list2 := ReadListsFromFile(inputPath)
	fmt.Println(CalculateDistance(list1, list2))
}

func RunPart2(inputPath string) {
	list1, list2 := ReadListsFromFile(inputPath)
	fmt.Println(CalculateSimilarityScore(list1, list2))
}
