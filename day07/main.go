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

	"golang.org/x/exp/maps"
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

	hands := []myHand{}
	for scanner.Scan() {
		line := scanner.Text()
		var hand string
		var bid int
		_, err := fmt.Sscanf(line, "%s %d", &hand, &bid)
		fatalOnError("Scanf", err)
		hands = append(hands, newHandType(hand, bid))
	}

	slices.SortFunc(hands, func(a, b myHand) int {
		if a.handType > b.handType {
			return 1
		}
		if a.handType < b.handType {
			return -1
		}
		for i := range a.withValue {
			if a.withValue[i] > b.withValue[i] {
				return 1
			}
			if a.withValue[i] < b.withValue[i] {
				return -1
			}
			continue
		}
		return 0
	})

	var total int64
	for i, h := range hands {
		total += (int64(i) + 1) * int64(h.bid)
	}

	return total
}

func secondPart(r io.Reader) int64 {
	return 0
}

type myHand struct {
	hand      string
	handType  int
	bid       int
	withValue string
}

func newHandType(hand string, bid int) myHand {
	withValue := strings.ReplaceAll(hand, "A", "Z")
	withValue = strings.ReplaceAll(withValue, "K", "Y")
	withValue = strings.ReplaceAll(withValue, "Q", "X")
	withValue = strings.ReplaceAll(withValue, "J", "W")
	withValue = strings.ReplaceAll(withValue, "T", "V")
	return myHand{
		hand:      hand,
		handType:  calcType(hand),
		bid:       bid,
		withValue: withValue,
	}
}

func calcType(hand string) int {
	h := []byte(hand)
	dedoub := make(map[byte]int)
	for _, c := range h {
		dedoub[c]++
	}
	if isFiveOfAKind(dedoub) {
		return 10
	}
	if isFourOfAKind(dedoub) {
		return 9
	}
	if isFullHouse(dedoub) {
		return 8
	}
	if isThreeOfAKind(dedoub) {
		return 7
	}
	if isTwoPair(dedoub) {
		return 6
	}
	if isOnePair(dedoub) {
		return 5
	}
	return 4
}

func isFiveOfAKind(dedoub map[byte]int) bool {
	return len(dedoub) == 1
}

func isFourOfAKind(dedoub map[byte]int) bool {
	values := maps.Values(dedoub)
	return len(dedoub) == 2 && (values[0] == 4 || values[1] == 4)
}

func isFullHouse(dedoub map[byte]int) bool {
	values := maps.Values(dedoub)
	return len(dedoub) == 2 && (values[0] == 3 || values[1] == 3)
}

func isThreeOfAKind(dedoub map[byte]int) bool {
	values := maps.Values(dedoub)
	return len(dedoub) == 3 && (values[0] == 3 || values[1] == 3 || values[2] == 3)
}

func isTwoPair(dedoub map[byte]int) bool {
	return len(dedoub) == 3
}

func isOnePair(dedoub map[byte]int) bool {
	return len(dedoub) == 4
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
