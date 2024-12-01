package p1distance_test

import (
	"aoc2024/p1distance"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistance(t *testing.T) {

	tests := []struct {
		name     string
		list1    []int
		list2    []int
		expected int
	}{
		{
			"Total distance should be 11",
			[]int{3,
				4,
				2,
				1,
				3,
				3},
			[]int{4,
				3,
				5,
				3,
				9,
				3},
			11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := p1distance.Distance(tt.list1, tt.list2)
			assert.Equal(t, res, tt.expected)
		})
	}
}
