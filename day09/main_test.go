package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstPart(t *testing.T) {
	input :=
		`0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`
	sum := firstPart(strings.NewReader((input)))
	assert.Equal(t, int64(114), sum)
}

func TestSecondPart(t *testing.T) {
	input :=
		`0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`
	sum := secondPart(strings.NewReader((input)))
	assert.Equal(t, int64(2), sum)
}
