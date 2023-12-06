package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstPart(t *testing.T) {
	input := ``
	sum := firstPart(strings.NewReader((input)))
	assert.Equal(t, 0, sum)
}
