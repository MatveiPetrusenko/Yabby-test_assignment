package main

import (
	"testing"
)

const height = 4
const width = 5
const colour = 3

func BenchmarkSearchMatches(b *testing.B) {
	var matrix matrices

	matrix.initialInputMatrix(height, width)
	matrix.initialVisitedMatrix(height, width)

	matrix.insertRandValue(colour)
	matrix.fillUpEmptyMatrix(colour)

	for i := 0; i < b.N; i++ {
		searchMatches(height, width, colour, matrix)
	}
}
