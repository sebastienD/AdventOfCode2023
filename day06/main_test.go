package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstPart(t *testing.T) {
	input := `Time:      7  15   30
Distance:  9  40  200`
	total := firstPart(strings.NewReader((input)))
	assert.Equal(t, 288, total)
}

func TestSecondPart(t *testing.T) {
	input := `Time:      7  15   30
Distance:  9  40  200`
	total := secondPart(strings.NewReader((input)))
	assert.Equal(t, uint64(71503), total)
}
