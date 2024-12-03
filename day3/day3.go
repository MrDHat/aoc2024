package day3

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func FindUncorruptedData(input string) []string {
	regex := regexp.MustCompile(`mul\(\d+(\.\d+)?,\d+(\.\d+)?\)`)
	return regex.FindAllString(input, -1)
}

func Multiply(str string) int {
	regex := regexp.MustCompile(`\(\d+(\.\d+)?,\d+(\.\d+)?\)`)
	data := regex.FindAllString(str, -1)

	numbers := strings.Split(data[0], ",")

	firstNumber, err := strconv.Atoi(strings.ReplaceAll(numbers[0], "(", ""))
	if err != nil {
		log.Fatal(err)
	}
	secondNumber, err := strconv.Atoi(strings.ReplaceAll(numbers[1], ")", ""))
	if err != nil {
		log.Fatal(err)
	}
	return firstNumber * secondNumber
}

func Calculate(input string) {
	cleanData := FindUncorruptedData(input)

	sum := 0

	for _, d := range cleanData {
		sum += Multiply(d)
	}

	fmt.Println(sum)
}
