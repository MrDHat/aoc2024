package day2

func isValidAscendingSequence(a, b int) bool {
	if a > b || a == b || b-a > 3 {
		return false
	}

	return true
}

func IsReportSafe(report []int) bool {
	isSafe := true
	if len(report) == 0 {
		return true
	}
	shouldBeAscending := true
	if report[0] > report[1] {
		shouldBeAscending = false
	}

	for i := 0; i < len(report)-1; i++ {
		if shouldBeAscending {
			if !isValidAscendingSequence(report[i], report[i+1]) {
				isSafe = false
				break
			}
		} else {
			if !isValidAscendingSequence(report[i+1], report[i]) {
				isSafe = false
				break
			}
		}
	}

	return isSafe
}

func NoOfSafeReports(reports [][]int) int {
	noOfSafeReports := 0

	if len(reports) == 0 {
		return noOfSafeReports
	}

	for _, report := range reports {
		if IsReportSafe(report) {
			noOfSafeReports++
		}
	}

	return noOfSafeReports
}
