package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
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
	total := firstPart(f)
	fmt.Printf("First part in %s, %d\n", time.Since(start), total)

	start = time.Now()
	_, err = f.Seek(0, io.SeekStart)
	fatalOnError("Seek", err)
	total64 := secondPart(f)
	fmt.Printf("Second part in %s, %d\n", time.Since(start), total64)

}

func firstPart(r io.Reader) int {
	scanner := bufio.NewScanner(r)

	scanner.Scan()
	ts, _ := strings.CutPrefix(scanner.Text(), "Time: ")
	times := slices.DeleteFunc(strings.Split(ts, " "), func(s string) bool {
		return s == ""
	})

	scanner.Scan()
	ls, _ := strings.CutPrefix(scanner.Text(), "Distance: ")
	lengths := slices.DeleteFunc(strings.Split(ls, " "), func(s string) bool {
		return s == ""
	})

	waysToWin := make([]int, len(times))
	for i, t := range times {
		time := Atoi(t)
		record := Atoi(lengths[i])
		for j := 1; j <= time; j++ {
			sc := j * (time - j)
			if sc > record {
				waysToWin[i]++
			}
		}
	}

	score := 1
	for _, v := range waysToWin {
		score *= v
	}
	return score
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

func secondPart(r io.Reader) uint64 {
	scanner := bufio.NewScanner(r)

	scanner.Scan()
	ts, _ := strings.CutPrefix(scanner.Text(), "Time: ")
	time := Atoi64(strings.ReplaceAll(ts, " ", ""))

	scanner.Scan()
	ls, _ := strings.CutPrefix(scanner.Text(), "Distance: ")
	record := Atoi64(strings.ReplaceAll(ls, " ", ""))

	fmt.Println(time, record)
	var waysToWin uint64
	var j uint64
	for ; j <= time; j++ {
		sc := j * (time - j)
		if sc > record {
			waysToWin++
		}
	}

	return waysToWin
}

func Atoi64(v string) uint64 {
	val, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		log.Fatalf("convert %q: %v", v, err)
	}
	return uint64(val)
}
