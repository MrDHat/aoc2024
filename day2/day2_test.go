package day2_test

import (
	"aoc2024/day2"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsReportSafe(t *testing.T) {
	tests := []struct {
		name     string
		report   []int
		expected bool
	}{
		{
			"Safe because the levels are all decreasing by 1 or 2",
			[]int{7, 6, 4, 2, 1},
			true,
		},
		{
			"Unsafe because 2 7 is an increase of 5",
			[]int{1, 2, 7, 8, 9},
			false,
		},
		{
			"Unsafe because 6 2 is a decrease of 4",
			[]int{9, 7, 6, 2, 1},
			false,
		},
		{
			"Unsafe because 1 3 is increasing but 3 2 is decreasing",
			[]int{1, 3, 2, 4, 5},
			false,
		},
		{
			"Unsafe because 4 4 is neither an increase or a decrease",
			[]int{8, 6, 4, 4, 1},
			false,
		},
		{
			"Safe because the levels are all increasing by 1, 2, or 3",
			[]int{1, 3, 6, 7, 9},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, day2.IsReportSafe(tt.report), tt.expected)
		})
	}
}
