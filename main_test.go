package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PrimesBelow(t *testing.T) {
	assert.Equal(t, []int{2, 3, 5, 7}, primesBelow(10))
	assert.Equal(t, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}, primesBelow(100))
	assert.Equal(t, []int{}, primesBelow(1))
	assert.Equal(t, []int{}, primesBelow(2))
	assert.Equal(t, []int{2}, primesBelow(3))
}
