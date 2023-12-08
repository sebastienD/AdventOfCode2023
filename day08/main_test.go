package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstPart(t *testing.T) {
	tt := map[string]struct {
		input    string
		expected int
	}{
		"simple": {`RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`, 2},
		"multiple": {`LLR

	AAA = (BBB, BBB)
	BBB = (AAA, ZZZ)
	ZZZ = (ZZZ, ZZZ)`, 6},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			steps := firstPart(strings.NewReader((tc.input)))
			assert.Equal(t, tc.expected, steps)
		})
	}
}

func TestSecondPart(t *testing.T) {
	input := `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`
	steps := secondPart(strings.NewReader((input)))
	assert.Equal(t, int64(6), steps)
}
