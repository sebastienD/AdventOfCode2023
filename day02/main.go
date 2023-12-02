package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := "input.txt"
	f, err := os.Open(filename)
	fatalOnError("Read file", err)
	defer f.Close()

	sum := firstPart(f)
	fmt.Printf("First part, %d\n", sum)

	_, err = f.Seek(0, io.SeekStart)
	fatalOnError("Seek", err)
	sum = secondPart(f)
	fmt.Printf("Second part, %d\n", sum)

}

func firstPart(r io.Reader) (sum int) {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		content := strings.Split(line, ":")
		if len(content) < 2 {
			continue
		}
		draws := strings.Split(content[1], ";")
		drawPossible := true
		for _, draw := range draws {
			colors := strings.Split(draw, ",")
			for _, colorCounter := range colors {
				cc := strings.Split(colorCounter, " ")
				if !IsPossible(cc[1], cc[2]) {
					drawPossible = false
					goto FINISH
				}
			}
		}
	FINISH:
		if drawPossible {
			game := strings.Split(content[0], " ")[1]
			i := Atoi(game)
			sum += i
		}
	}
	return
}

func IsPossible(v string, color string) bool {
	val := Atoi(v)
	maXColors := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	return val <= maXColors[color]
}

func secondPart(r io.Reader) (sum int) {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		content := strings.Split(line, ":")
		if len(content) < 2 {
			continue
		}
		draws := strings.Split(content[1], ";")
		highestColors := map[string]int{
			"red":   0,
			"blue":  0,
			"green": 0,
		}
		for _, draw := range draws {
			colors := strings.Split(draw, ",")
			for _, colorCounter := range colors {
				cc := strings.Split(colorCounter, " ")

				color := cc[2]
				counter := Atoi(cc[1])
				if counter > highestColors[color] {
					highestColors[color] = counter
				}
			}
		}
		sum += computePower(highestColors)
	}
	return
}

func computePower(m map[string]int) int {
	power := 1
	for _, v := range m {
		power *= v
	}
	return power
}

func Atoi(v string) int {
	val, err := strconv.Atoi(v)
	if err != nil {
		log.Fatalf("convert %s: %v", v, err)
	}
	return val
}

func fatalOnError(message string, err error) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}
