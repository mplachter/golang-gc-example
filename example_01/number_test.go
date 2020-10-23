package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHighestOccuringNumbers(t *testing.T) {
	nm := NumberStorage{
		numbers: []int{
			1, 1, 2, 2, 3,
		},
	}

	numbers := nm.HighestOccuringNumbers()

	assert.Equal(t, []int{1, 2}, numbers)
}
