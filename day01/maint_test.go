package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstPart(t *testing.T) {
	input := `1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet
	`
	sum := firstPart(strings.NewReader(input))
	assert.Equal(t, 142, sum)
}

func TestSecondPart(t *testing.T) {
	input := `two1nine
	eightwothree
	abcone2threexyz
	xtwone3four
	4nineeightseven2
	zoneight234
	7pqrstsixteen
	`
	sum := secondPart(strings.NewReader(input))
	assert.Equal(t, 281, sum)
}
