package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
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
	sum := firstPart(f)
	fmt.Printf("First part in %s, %d\n", time.Since(start), sum)

	//start = time.Now()
	//_, err = f.Seek(0, io.SeekStart)
	//fatalOnError("Seek", err)
	//lowest = secondPart(f)
	//fmt.Printf("Second part in %s, %d\n", time.Since(start), lowest)

}

func firstPart(r io.Reader) int {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	return 0
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
