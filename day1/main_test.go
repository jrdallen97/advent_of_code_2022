package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopElf(t *testing.T) {
	tests := []struct {
		input         []int
		expectedIndex int
	}{
		{
			[]int{100, 0, 99},
			0,
		},
		{
			[]int{99, 0, 100},
			2,
		},
		{
			[]int{98, 100, 99},
			1,
		},
	}

	for _, test := range tests {
		idx, val := TopElf(test.input)
		assert.Equal(t, test.expectedIndex, idx)
		assert.Equal(t, test.input[test.expectedIndex], val)
	}
}
