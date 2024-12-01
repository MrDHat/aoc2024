package p1distance

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Run(path string) {
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

	fmt.Println(Distance(list1, list2))
}

func Distance(locationList1 []int, locationList2 []int) int {

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
