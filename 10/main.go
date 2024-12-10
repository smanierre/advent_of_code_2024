package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Coordinate struct {
	X    int
	Y    int
	Val  int
	Done bool
}

func (c Coordinate) GetNextSteps(m [][]int) []Direction {
	var dirs []Direction
	if c.X-1 > -1 && m[c.Y][c.X-1]-1 == c.Val {
		dirs = append(dirs, Left)
	}
	if c.X+1 < len(m[0]) && m[c.Y][c.X+1]-1 == c.Val {
		dirs = append(dirs, Right)
	}
	if c.Y-1 > -1 && m[c.Y-1][c.X]-1 == c.Val {
		dirs = append(dirs, Up)
	}
	if c.Y+1 < len(m) && m[c.Y+1][c.X]-1 == c.Val {
		dirs = append(dirs, Down)
	}
	return dirs
}

func (c Coordinate) Move(d Direction) Coordinate {
	c2 := Coordinate{}
	switch d {
	case Up:
		c2 = Coordinate{X: c.X, Y: c.Y - 1, Val: c.Val + 1}
	case Down:
		c2 = Coordinate{X: c.X, Y: c.Y + 1, Val: c.Val + 1}
	case Left:
		c2 = Coordinate{X: c.X - 1, Y: c.Y, Val: c.Val + 1}
	case Right:
		c2 = Coordinate{X: c.X + 1, Y: c.Y, Val: c.Val + 1}
	}
	if c2.Val == 9 {
		c2.Done = true
	}
	return c2
}

func (c Coordinate) FindPaths(m [][]int, peaks map[Coordinate]int) {
	for _, dir := range c.GetNextSteps(m) {
		coord := c.Move(dir)
		if coord.Done {
			peaks[coord]++
		}
		coord.FindPaths(m, peaks)
	}
}

func main() {
	t := time.Now()
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error opening input file: %s", err.Error())
	}
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatalf("Error reading input: %s", err.Error())
	}
	var m [][]int
	var trailheads []Coordinate
	for Y, line := range strings.Split(string(b), "\n") {
		var l []int
		for X, spot := range strings.Split(line, "") {
			n, _ := strconv.Atoi(spot)
			l = append(l, n)
			if n == 0 {
				trailheads = append(trailheads, Coordinate{X: X, Y: Y, Val: n})
			}
		}
		m = append(m, l)
	}
	fmt.Printf("Preparation time: %s\n", time.Since(t))
	partOne(m, trailheads)
	partTwo(m, trailheads)
}

func partOne(m [][]int, trailheads []Coordinate) {
	t := time.Now()
	peaks := map[Coordinate]int{}
	total := 0
	for _, c := range trailheads {
		c.FindPaths(m, peaks)
		total += len(peaks)
		peaks = map[Coordinate]int{}
	}

	fmt.Println(total)
	fmt.Println("Part one:", time.Since(t))
}

func partTwo(m [][]int, trailheads []Coordinate) {
	t := time.Now()
	peaks := map[Coordinate]int{}
	total := 0
	for _, c := range trailheads {
		c.FindPaths(m, peaks)
		for _, v := range peaks {
			total += v
		}
		peaks = map[Coordinate]int{}
	}

	fmt.Println(total)
	fmt.Println("Part one:", time.Since(t))
}
