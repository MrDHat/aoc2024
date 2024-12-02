package day2_test

import (
	"aoc2024/day2"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoOfSafeReports(t *testing.T) {
	tests := []struct {
		name        string
		report      []int
		canTolerate bool
		expected    int
	}{
		{
			"1. Safe at tolerance level 0 because the levels are all decreasing by 1 or 2",
			[]int{7, 6, 4, 2, 1},
			false,
			1,
		},
		{
			"2. Unsafe at tolerance level 0 because 2 7 is an increase of 5",
			[]int{1, 2, 7, 8, 9},
			false,
			0,
		},
		{
			"3. Unsafe at tolerance level 0 because 6 2 is a decrease of 4",
			[]int{9, 7, 6, 2, 1},
			false,
			0,
		},
		{
			"4. Unsafe at tolerance level 0 because 1 3 is increasing but 3 2 is decreasing",
			[]int{1, 3, 2, 4, 5},
			false,
			0,
		},
		{
			"5. Unsafe at tolerance level 0 because 4 4 is neither an increase or a decrease",
			[]int{8, 6, 4, 4, 1},
			false,
			0,
		},
		{
			"6. Safe at tolerance level 0 because the levels are all increasing by 1, 2, or 3",
			[]int{1, 3, 6, 7, 9},
			false,
			1,
		},
		{"7. Safe at tolerance level 1 without removing any level",
			[]int{7, 6, 4, 2, 1},
			true,
			1,
		},
		{"8. Unsafe at tolerance level 1 regardless of which level is removed",
			[]int{1, 2, 7, 8, 9},
			true,
			0,
		},
		{"9. Unsafe at tolerance level 1 regardless of which level is removed",
			[]int{9, 7, 6, 2, 1},
			true,
			0,
		},
		{"10. Safe at tolerance level 1 by removing the second level, 3",
			[]int{1, 3, 2, 4, 5},
			true,
			1,
		},
		{"11. Safe at tolerance level 1 by removing the third level, 4",
			[]int{8, 6, 4, 4, 1},
			true,
			1,
		},
		{"12. Safe at tolerance level 1 without removing any level",
			[]int{1, 3, 6, 7, 9},
			true,
			1,
		},
		{"13. Safe at tolerance level 1 by removing 1 level",
			[]int{75, 77, 72, 70, 69},
			true,
			1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, day2.NoOfSafeReports([][]int{tt.report}, tt.canTolerate))
		})
	}
}
