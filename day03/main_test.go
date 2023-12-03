package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstPart(t *testing.T) {
	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
	mat := newMatrice(strings.NewReader(input))
	sum := firstPart(mat)
	assert.Equal(t, 4361, sum)
}

func TestSecondPart(t *testing.T) {
	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
	mat := newMatrice(strings.NewReader(input))
	sum := secondPart(mat)
	assert.Equal(t, int64(467835), sum)
}

func TestComputeNumbers(t *testing.T) {
	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
	mat := newMatrice(strings.NewReader(input))
	numbers := mat.computeNumbers()
	assert.Equal(t, numbers[0], []point{{
		val:        "467",
		startIndex: 0,
		endIndex:   2,
	}, {
		val:        "114",
		startIndex: 5,
		endIndex:   7,
	}})
	assert.Equal(t, 0, len(numbers[1]))
	assert.Equal(t, 2, len(numbers[2]))
	assert.Equal(t, 0, len(numbers[3]))
	assert.Equal(t, 1, len(numbers[4]))
	assert.Equal(t, 1, len(numbers[5]))
	assert.Equal(t, 1, len(numbers[6]))
	assert.Equal(t, 1, len(numbers[7]))
	assert.Equal(t, 0, len(numbers[8]))
	assert.Equal(t, 2, len(numbers[9]))
}

func TestSpecialComputeNumbers(t *testing.T) {
	input := `...811.656.
...........
....871*...
..*.....222
48.448.....`
	mat := newMatrice(strings.NewReader(input))
	numbers := mat.computeNumbers()
	assert.Equal(t, numbers[0], []point{{
		val:        "811",
		startIndex: 3,
		endIndex:   5,
	}, {
		val:        "656",
		startIndex: 7,
		endIndex:   9,
	}})
	assert.Equal(t, 0, len(numbers[1]))
	assert.Equal(t, 1, len(numbers[2]))
	assert.Equal(t, 1, len(numbers[3]))
	assert.Equal(t, 2, len(numbers[4]))
}
