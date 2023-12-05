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
	"time"
)

func main() {
	filename := "input.txt"
	f, err := os.Open(filename)
	fatalOnError("Read file", err)
	defer f.Close()

	start := time.Now()
	lowest := firstPart(f)
	fmt.Printf("First part in %s, %d\n", time.Since(start), lowest)

	start = time.Now()
	_, err = f.Seek(0, io.SeekStart)
	fatalOnError("Seek", err)
	lowest = secondPart(f)
	fmt.Printf("Second part in %s, %d\n", time.Since(start), lowest)

}

func firstPart(r io.Reader) int64 {
	scanner := bufio.NewScanner(r)

	scanner.Scan()
	se, _ := strings.CutPrefix(scanner.Text(), "seeds: ")
	ses := strings.Split(se, " ")
	seeds := make([]int64, len(ses))
	for i, v := range ses {
		seeds[i] = Atoi(v)
	}
	scanner.Scan()
	scanner.Scan()

	kMaps := kindMaps{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			update(seeds, kMaps)
			kMaps = kindMaps{}
			scanner.Scan()
			continue
		}
		var source, dest, interval int64
		fmt.Sscanf(line, "%d %d %d", &dest, &source, &interval)
		km := kindMap{
			source:   source,
			dest:     dest,
			interval: interval,
		}
		kMaps = append(kMaps, km)
	}
	update(seeds, kMaps)

	return slices.Min(seeds)
}

func update(seeds []int64, kMaps kindMaps) {
	slices.SortFunc(kMaps, kmSort)
	for i, v := range seeds {
		seeds[i] = kMaps.destValue(v)
	}
}

type kindMap struct {
	source   int64
	dest     int64
	interval int64
}

type kindMaps []kindMap

func (km kindMaps) destValue(source int64) int64 {
	for _, v := range km {
		if source < v.source {
			break
		}
		if source < v.source+v.interval {
			return v.dest + source - v.source
		}
	}
	return source
}

func secondPart(r io.Reader) int64 {
	scanner := bufio.NewScanner(r)

	scanner.Scan()
	se, _ := strings.CutPrefix(scanner.Text(), "seeds: ")
	ses := strings.Split(se, " ")
	scanner.Scan()
	scanner.Scan()

	all := []kindMaps{}
	kMaps := kindMaps{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			slices.SortFunc(kMaps, kmSort)
			all = append(all, slices.Clone(kMaps))
			kMaps = kindMaps{}
			scanner.Scan()
			continue
		}
		var source, dest, interval int64
		fmt.Sscanf(line, "%d %d %d", &dest, &source, &interval)
		km := kindMap{
			source:   source,
			dest:     dest,
			interval: interval,
		}
		kMaps = append(kMaps, km)
	}
	slices.SortFunc(kMaps, kmSort)
	all = append(all, slices.Clone(kMaps))

	var min int64 = math.MaxInt64
	for i, v := range ses {
		if i == len(ses)-1 {
			break
		}
		if i%2 == 1 {
			continue
		}
		maxLoop := Atoi(ses[i+1])
		var loop int64
		for loop < maxLoop {
			step := Atoi(v) + loop
			for _, km := range all {
				step = km.destValue(step)
			}
			if step < min {
				min = step
			}
			loop++
		}
	}

	return min
}

func kmSort(a, b kindMap) int {
	if a.source < b.source {
		return -1
	}
	if a.source > b.source {
		return 1
	}
	return 0
}

func fatalOnError(message string, err error) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}

func Atoi(v string) int64 {
	val, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		log.Fatalf("convert %q: %v", v, err)
	}
	return val
}
