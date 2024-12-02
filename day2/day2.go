package day2

func isValidAscendingSequence(a, b int) bool {
	if a > b || a == b || b-a > 3 {
		return false
	}

	return true
}

func removeIdx(data []int, idx int) []int {
	ret := make([]int, 0)
	ret = append(ret, data[:idx]...)
	return append(ret, data[idx+1:]...)
}

func UnsafeIndex(report []int) int {
	shouldBeAscending := true
	if report[0] > report[1] {
		shouldBeAscending = false
	}

	for i := 0; i < len(report)-1; i++ {
		if shouldBeAscending {
			if !isValidAscendingSequence(report[i], report[i+1]) {
				return i
			}
		} else {
			if !isValidAscendingSequence(report[i+1], report[i]) {
				return i
			}
		}
	}
	return -1
}

func NoOfSafeReports(reports [][]int, canTolerate bool) int {
	noOfSafeReports := 0

	if len(reports) == 0 {
		return noOfSafeReports
	}

	for _, report := range reports {
		unsafeIdx := UnsafeIndex(report)
		if unsafeIdx == -1 {
			noOfSafeReports++
		} else if canTolerate {
			// run the same with index & index + 1
			// if any of them returns true, its safe!
			secondRunIdx := UnsafeIndex(removeIdx(report, unsafeIdx))
			if secondRunIdx == -1 {
				noOfSafeReports++
			} else {
				thirdRunIdx := UnsafeIndex(removeIdx(report, unsafeIdx+1))
				if thirdRunIdx == -1 {
					noOfSafeReports++
				}
			}
		}
	}

	return noOfSafeReports
}
