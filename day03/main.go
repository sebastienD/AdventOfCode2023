package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"unicode"

	"golang.org/x/exp/maps"
)

func main() {
	filename := "input.txt"
	f, err := os.Open(filename)
	fatalOnError("Read file", err)
	defer f.Close()

	m := newMatrice(f)
	sum := firstPart(m)
	fmt.Printf("First part, %d\n", sum)

	_, err = f.Seek(0, io.SeekStart)
	fatalOnError("Seek", err)
	sum64 := secondPart(m)
	fmt.Printf("Second part, %d\n", sum64)

}

func firstPart(mat matrice) (sum int) {
	for x, line := range mat {
		var num string
		var charDetected bool
		for y, _ := range line {
			if mat.isDigit(x, y) {
				if len(num) == 0 {
					charDetected = false
					num = ""
				}
				num += string(mat[x][y])
				if mat.haveCharAround(x, y) {
					charDetected = true
				}
				if !mat.isOut(x, y+1) {
					continue
				}
			}
			if len(num) > 0 && charDetected {
				sum += Atoi(num)
			}
			charDetected = false
			num = ""
		}
	}
	return
}

func secondPart(mat matrice) (sum int64) {
	points := mat.computeNumbers()
	for x, line := range mat {
		for y, _ := range line {
			isGear, allPoints := mat.isGear(x, y, points)
			if isGear {
				ratio := int64(Atoi(allPoints[0].val)) * int64(Atoi(allPoints[1].val))
				sum += ratio
			}
		}
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

type point struct {
	val        string
	startIndex int
	endIndex   int
}

func (p point) hasNoValue() bool {
	return p.val == ""
}

//	--- y y y
//	--- 0 1 2
//
// x 0
// x 1
// x 2
type matrice [][]byte

func newMatrice(r io.Reader) matrice {
	fileScanner := bufio.NewScanner(r)
	in := make([][]byte, 0, 10)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		in = append(in, []byte(line))
	}
	return matrice(in)
}

func (m matrice) isDigit(x, y int) bool {
	pos := m[x][y]
	return unicode.IsDigit(rune(pos))
}

func (m matrice) isDot(x, y int) bool {
	return m[x][y] == '.'
}

func (m matrice) computeNumbers() [][]point {
	points := make([][]point, 0, 10)
	for x, line := range m {
		var p point
		ps := []point{}
		for y, c := range line {
			if m.isDigit(x, y) {
				if p.hasNoValue() {
					p.startIndex = y
				}
				p.val += string(c)
				if m.isLastDigit(x, y) {
					p.endIndex = y
					ps = append(ps, p)
					p = point{}
				}
			}
		}
		points = append(points, ps)
	}
	return points
}

func (m matrice) haveCharAround(x, y int) bool {
	return m.isChar(x-1, y-1) || m.isChar(x-1, y) || m.isChar(x-1, y+1) ||
		m.isChar(x, y-1) || m.isChar(x, y+1) ||
		m.isChar(x+1, y-1) || m.isChar(x+1, y) || m.isChar(x+1, y+1)
}

func (m matrice) isChar(x, y int) bool {
	if m.isOut(x, y) {
		return false
	}
	if m.isDigit(x, y) || m.isDot(x, y) {
		return false
	}
	return true
}

func (m matrice) isLastDigit(x, y int) bool {
	if m.isOut(x, y+1) {
		return true
	}
	return !m.isDigit(x, y+1)
}

func (m matrice) isOut(x, y int) bool {
	return x < 0 || y < 0 || x >= len(m) || y >= len(m[x])
}

func (m matrice) isGear(x, y int, points [][]point) (bool, []point) {
	if m[x][y] != '*' {
		return false, nil
	}

	res := make(map[point]struct{}, 0)
	res[m.pointHere(x-1, y-1, points)] = struct{}{}
	res[m.pointHere(x-1, y, points)] = struct{}{}
	res[m.pointHere(x-1, y+1, points)] = struct{}{}
	res[m.pointHere(x, y-1, points)] = struct{}{}
	res[m.pointHere(x, y+1, points)] = struct{}{}
	res[m.pointHere(x+1, y-1, points)] = struct{}{}
	res[m.pointHere(x+1, y, points)] = struct{}{}
	res[m.pointHere(x+1, y+1, points)] = struct{}{}

	delete(res, point{})

	allNumbers := maps.Keys(res)

	if len(allNumbers) != 2 {
		return false, nil
	}
	return true, allNumbers
}

func (m matrice) PrintAround(x, y int) {
	stringAt := func(x, y int) string {
		if m.isOut(x, y) {
			return ""
		}
		return string(m[x][y])
	}
	fmt.Printf("%s%s%s\n", stringAt(x-1, y-1), stringAt(x-1, y), stringAt(x-1, y+1))
	fmt.Printf("%s%s%s\n", stringAt(x, y-1), stringAt(x, y), stringAt(x, y+1))
	fmt.Printf("%s%s%s\n\n", stringAt(x+1, y-1), stringAt(x+1, y), stringAt(x+1, y+1))
}

func (m matrice) pointHere(x, y int, points [][]point) point {
	if m.isOut(x, y) || !m.isDigit(x, y) {
		return point{}
	}
	line := points[x]
	for _, p := range line {
		if y >= p.startIndex && y <= p.endIndex {
			return p
		}
	}
	return point{}
}
