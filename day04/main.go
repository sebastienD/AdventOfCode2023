package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

func main() {
	filename := "input.txt"
	f, err := os.Open(filename)
	fatalOnError("Read file", err)
	defer f.Close()

	score := firstPart(f)
	fmt.Printf("First part, %d\n", score)

	_, err = f.Seek(0, io.SeekStart)
	fatalOnError("Seek", err)
	total := secondPart(f)
	fmt.Printf("Second part, %d\n", total)

}

func firstPart(r io.Reader) int64 {
	scanner := bufio.NewScanner(r)

	var sum int64
	for scanner.Scan() {
		line := scanner.Text()
		content := strings.Split(strings.Split(line, ":")[1], "|")

		firstP := content[0]
		second := slices.DeleteFunc(strings.Split(content[1], " "), func(e string) bool {
			return e == ""
		})
		intersect := make([]string, 0)
		for _, v := range second {
			if strings.Contains(firstP, " "+v+" ") {
				intersect = append(intersect, v)
			}
		}
		pow := len(intersect)
		if pow == 0 {
			continue
		}
		sum += int64(math.Pow(2, float64(len(intersect)-1)))
	}

	return sum
}

func secondPart(r io.Reader) (total int) {
	scanner := bufio.NewScanner(r)

	cards := make(map[int]int, 0)
	current := 1
	for scanner.Scan() {
		line := scanner.Text()
		start := strings.Split(line, ":")
		content := strings.Split(start[1], "|")

		firstP := content[0]
		second := slices.DeleteFunc(strings.Split(content[1], " "), func(e string) bool {
			return e == ""
		})
		nbIntersect := 0
		for _, v := range second {
			if strings.Contains(firstP, " "+v+" ") {
				nbIntersect++
			}
		}

		cards[current]++
		for nbIntersect > 0 {
			cards[current+nbIntersect] += cards[current]
			nbIntersect--
		}
		fmt.Println(cards)
		current++
	}

	for _, v := range maps.Values(cards) {
		total += v
	}

	return
}

func fatalOnError(message string, err error) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}

func Atoi(v string) int {
	val, err := strconv.Atoi(v)
	if err != nil {
		log.Fatalf("convert %q: %v", v, err)
	}
	return val
}
