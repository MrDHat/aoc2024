package day3_test

import (
	"aoc2024/day3"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindUncorruptedData(t *testing.T) {
	resp := day3.FindUncorruptedData("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))mul(2000 , 1000)")
	assert.Equal(t, []string{
		"mul(2,4)",
		"mul(5,5)",
		"mul(11,8)",
		"mul(8,5)",
	}, resp)

}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			"Should return 8",
			"mul(2,4)",
			8,
		},
		{
			"Should return 28000",
			"mul(500,56)",
			28000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, day3.Multiply(tt.input))
		})
	}
}

func TestFindUncorruptedDataWithPast(t *testing.T) {

	resp, _ := day3.FindUncorruptedDataWithPast("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64]do()(mul(11,8)undo()?mul(8,5))", true)
	assert.Equal(t, []string{
		"mul(2,4)",
		"mul(11,8)",
		"mul(8,5)",
	}, resp)

	resp2, _ := day3.FindUncorruptedDataWithPast("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64]do()(mul(11,8)undo()?mul(8,5))", true)
	assert.Equal(t, []string{
		"mul(2,4)",
		"mul(11,8)",
		"mul(8,5)",
	}, resp2)

	resp3, _ := day3.FindUncorruptedDataWithPast("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64]do()(mul(11,8)undo()?mul(8,5))", false)
	assert.Equal(t, []string{
		"mul(11,8)",
		"mul(8,5)",
	}, resp3)
}
