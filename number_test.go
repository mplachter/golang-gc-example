package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHighestOccuringNumbers(t *testing.T) {

	tcs := []struct {
		numbers []int
		result  []int
	}{
		{[]int{1, 1, 2, 2, 3, 3, 4}, []int{1, 2, 3}},
		{[]int{1, 2, 3, 2, 4, 5, 4, 6, 4}, []int{4}},
	}

	for _, tc := range tcs {
		nm := NumberStorage{
			numbers: tc.numbers,
		}
		numbers := nm.HighestOccuringNumbers()
		assert.Equal(t, tc.result, numbers)
	}

}
