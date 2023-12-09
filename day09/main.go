package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	filename := "input.txt"
	f, err := os.Open(filename)
	fatalOnError("Read file", err)
	defer f.Close()

	start := time.Now()
	sum := firstPart(f)
	fmt.Printf("First part in %s, %d\n", time.Since(start), sum)

	start = time.Now()
	_, err = f.Seek(0, io.SeekStart)
	fatalOnError("Seek", err)
	sum = secondPart(f)
	fmt.Printf("Second part in %s, %d\n", time.Since(start), sum)

}

func firstPart(r io.Reader) int64 {
	scanner := bufio.NewScanner(r)

	dataset := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		h := strings.Fields(line)
		histo := make([]int, len(h))
		for i, v := range h {
			histo[i] = Atoi(v)
		}
		dataset = append(dataset, histo)
	}

	var sum int64

	for _, current := range dataset {

		histo := newHistory(current)
		for !contentsOnlyZero(histo) {
			histo = computeNextDiff(histo)
		}
		sum += nextValue(histo)
	}

	return sum
}

func secondPart(r io.Reader) int64 {
	scanner := bufio.NewScanner(r)

	dataset := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		h := strings.Fields(line)
		histo := make([]int, len(h))
		for i, v := range h {
			histo[i] = Atoi(v)
		}
		dataset = append(dataset, histo)
	}

	var sum int64
	for _, current := range dataset {
		histo := newHistory(current)
		for !contentsOnlyZero(histo) {
			histo = computeNextDiff(histo)
		}
		sum += previousValue(histo)
	}

	return sum
}

type history [][]int

func newHistory(histo []int) history {
	history := make([][]int, 0)
	history = append(history, histo)
	return history
}

func computeNextDiff(h history) history {
	last := h[len(h)-1]

	next := make([]int, len(last)-1)
	for i := 0; i < len(last)-1; i++ {
		next[i] = last[i+1] - last[i]
	}

	return append(h, next)
}

func contentsOnlyZero(h history) bool {
	last := h[len(h)-1]
	for _, v := range last {
		if v != 0 {
			return false
		}
	}
	return true
}

func nextValue(h history) int64 {
	var sum int64
	for _, histo := range h {
		sum += int64(histo[len(histo)-1])
	}
	return sum
}

func previousValue(h history) int64 {
	var diff int64
	for i := len(h) - 2; i > -1; i-- {
		diff = int64(h[i][0]) - diff
	}
	return diff
}

func fatalOnError(message string, err error) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}

func Atoi(v string) int {
	val, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		log.Fatalf("convert %q: %v", v, err)
	}
	return int(val)
}
