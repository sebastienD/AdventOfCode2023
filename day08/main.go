package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

func main() {
	filename := "input.txt"
	f, err := os.Open(filename)
	fatalOnError("Read file", err)
	defer f.Close()

	start := time.Now()
	steps := firstPart(f)
	fmt.Printf("First part in %s, %d\n", time.Since(start), steps)

	start = time.Now()
	_, err = f.Seek(0, io.SeekStart)
	fatalOnError("Seek", err)
	allSteps := secondPart(f)
	fmt.Printf("Second part in %s, %d\n", time.Since(start), allSteps)

}

func firstPart(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	instructions := scanner.Text()

	network := map[string][2]string{}
	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		var k, l, r string
		fmt.Sscanf(line, `%3s = (%3s, %3s)`, &k, &l, &r)
		network[k] = [2]string{l, r}
	}

	var steps int
	step := "AAA"
	for {
		for _, c := range instructions {
			instruction := network[step]
			if c == rune('L') {
				step = instruction[0]
			} else {
				step = instruction[1]
			}
			steps++
			if step == "ZZZ" {
				goto FINISH1
			}
		}
	}
FINISH1:

	return steps
}

func secondPart(r io.Reader) int64 {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	instructions := scanner.Text()

	network := map[string][2]string{}
	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		var k, l, r string
		fmt.Sscanf(line, `%3s = (%3s, %3s)`, &k, &l, &r)
		network[k] = [2]string{l, r}
	}

	steps := []string{}
	for k := range network {
		if k[2] == 'A' {
			steps = append(steps, k)
		}
	}

	var stepCount int
	firstZ := make([]int, len(steps))
	for {
		for _, c := range instructions {
			for i := range steps {
				tuple := network[steps[i]]
				if c == rune('L') {
					steps[i] = tuple[0]
				} else {
					steps[i] = tuple[1]
				}
			}
			stepCount++
			for i, v := range steps {
				if v[2] == 'Z' {
					firstZ[i] = stepCount
				}
			}
			done := true
			for _, v := range firstZ {
				done = done && v > 0
			}
			if done {
				goto FINISH2
			}
		}
	}
FINISH2:

	var result int64 = 1
	for _, v := range firstZ {
		result = lcm(result, int64(v))
	}

	return result
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

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int64) int64 {
	return int64(math.Abs(float64(a*b)) / float64(gcd(a, b)))
}
