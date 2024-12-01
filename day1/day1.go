package day1

import (
	"sort"
)

func CalculateDistance(locationList1 []int, locationList2 []int) int {

	// sort the lists
	sort.Slice(locationList1, func(i, j int) bool {
		return locationList1[i] < locationList1[j]
	})
	sort.Slice(locationList2, func(i, j int) bool {
		return locationList2[i] < locationList2[j]
	})

	distance := 0
	for i := range locationList1 {
		dist := locationList1[i] - locationList2[i]
		if dist < 0 {
			distance += -dist
		} else {
			distance += dist
		}
	}

	return distance

}

func CalculateSimilarityScore(locationList1 []int, locationList2 []int) int {
	secondListFrequency := map[int]int{}

	for _, v := range locationList2 {
		_, ok := secondListFrequency[v]
		if !ok {
			secondListFrequency[v] = 1
		} else {
			secondListFrequency[v]++
		}
	}

	similarityScore := 0
	for _, v := range locationList1 {
		_, ok := secondListFrequency[v]
		if ok {
			similarityScore += v * secondListFrequency[v]
		}
	}

	return similarityScore
}
