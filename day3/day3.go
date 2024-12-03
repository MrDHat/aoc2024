package day3

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func FindUncorruptedData(input string) []string {
	return regexp.MustCompile(`mul\(\d+(\.\d+)?,\d+(\.\d+)?\)`).FindAllString(input, -1)
}

func FindUncorruptedDataWithPast(input string, isLastOneAMatch bool) ([]string, bool) {
	regex := regexp.MustCompile(`do\(\)|don't\(\)|mul\(\d+(\.\d+)?,\d+(\.\d+)?\)`)
	resp := []string{}

	matches := regex.FindAllString(input, -1)
	isAMatch := isLastOneAMatch
	for _, m := range matches {
		if m == "do()" {
			isAMatch = true
		} else if m == "don't()" {
			isAMatch = false
		} else {
			if isAMatch {
				resp = append(resp, m)
			}
		}
	}

	return resp, isAMatch
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

func Calculate(input string, supportPast bool) {
	cleanData := []string{}
	lines := strings.Split(input, "\n")
	isLastOneAMatch := true
	for _, line := range lines {
		if supportPast {
			var data []string
			data, isLastOneAMatch = FindUncorruptedDataWithPast(line, isLastOneAMatch)
			cleanData = append(cleanData, data...)
		} else {
			cleanData = append(cleanData, FindUncorruptedData(line)...)
		}
	}

	sum := 0

	for _, d := range cleanData {
		sum += Multiply(d)
	}

	fmt.Println(sum)
}
