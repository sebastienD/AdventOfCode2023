package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	filename := "input.txt"
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Read file %s: %v", filename, err)
	}
	defer f.Close()

	sum := firstPart(f)
	fmt.Printf("First part, %d\n", sum)

	f.Seek(0, io.SeekStart)
	sum = secondPart(f)
	fmt.Printf("Second part, %d\n", sum)

}

func firstPart(r io.Reader) (sum int) {
	fileScanner := bufio.NewScanner(r)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		first, second := "", ""
		for _, char := range line {
			if unicode.IsDigit(char) {
				if first == "" {
					first = string(char)
				} else {
					second = string(char)
				}
			}
		}
		if first == "" {
			continue
		}
		if second == "" {
			second = first
		}
		val, err := strconv.Atoi(first + second)
		if err != nil {
			log.Fatalf("Convertion: %v\n", err)
		}
		sum += val
	}

	return
}

func secondPart(r io.Reader) (sum int) {
	var digits = []struct {
		show string
		val  string
	}{
		{"one", "1"},
		{"two", "2"},
		{"three", "3"},
		{"four", "4"},
		{"five", "5"},
		{"six", "6"},
		{"seven", "7"},
		{"eight", "8"},
		{"nine", "9"},
		{"1", "1"},
		{"2", "2"},
		{"3", "3"},
		{"4", "4"},
		{"5", "5"},
		{"6", "6"},
		{"7", "7"},
		{"8", "8"},
		{"9", "9"},
		{"0", "0"},
	}

	fileScanner := bufio.NewScanner(r)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		first, second := "", ""
		firstIndex, lastIndex := -2, -2
		for _, digit := range digits {
			firstI := strings.Index(line, digit.show)
			if firstI > -1 {
				if firstIndex == -2 || firstI < firstIndex {
					firstIndex = firstI
					first = digit.val
				}
			}
			lastI := strings.LastIndex(line, digit.show)
			if lastI > -1 {
				if lastIndex == -2 || lastI > lastIndex {
					lastIndex = lastI
					second = digit.val
				}
			}
		}

		if first == "" {
			continue
		}
		if second == "" {
			second = first
		}
		val, err := strconv.Atoi(first + second)
		if err != nil {
			log.Fatalf("Convertion: %v\n", err)
		}
		sum += val
	}
	return
}
