package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatalf("Error opening input file: %s\n", err.Error())
	}
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatalf("Error reading input file: %s\n", err.Error())
	}
	tempRows := strings.Split(string(b), "\n")
	var rows [][]string
	for _, row := range tempRows {
		rows = append(rows, strings.Split(row, ""))
	}
	found := 0
	for y, row := range rows {
		for x, _ := range row {
			found += IsXMAS(x, y, rows)
		}
	}
	fmt.Println(found)
}

// Above copied from part 1

func IsXMAS(x, y int, data [][]string) int {
	// Key off A and search from there since it's in the middle
	if data[y][x] != "A" {
		return 0
	}
	dirs := GetSafeDirs(x, y, len(data[0]), len(data))
	// If we can't search each direction, disregard
	if len(dirs) != 4 {
		return 0
	}
	return FindXMAS(x, y, dirs, data)
}

type Direction int

const (
	UpRight Direction = iota
	UpLeft
	DownRight
	DownLeft
)

func GetSafeDirs(x, y, width, height int) []Direction {
	dirs := []Direction{}
	// Only care about the diagonals now
	// Check up right, need 1 rows above and 1 columns after
	if x < width-1 && y > 0 {
		dirs = append(dirs, UpRight)
	}
	// Check up Left, need 1 rows above and 1 columns before
	if x > 0 && y > 0 {
		dirs = append(dirs, UpLeft)
	}
	// Check Down right, need 1 rows below and 1 columns after
	if x < width-1 && y < height-1 {
		dirs = append(dirs, DownRight)
	}
	// Check Down left, need 1 rows below and 1 columns before
	if x > 0 && y < height-1 {
		dirs = append(dirs, DownLeft)
	}
	return dirs
}

func FindXMAS(x, y int, dirs []Direction, data [][]string) int {
	found := 0
	for _, dir := range dirs {
		// Only need to check up right and left, and compare to the opposite side based on letter
		switch dir {
		case UpRight:
			if (data[y-1][x+1] == "M" && data[y+1][x-1] == "S") || (data[y-1][x+1] == "S" && data[y+1][x-1] == "M") {
				found++
			}
		case UpLeft:
			if (data[y-1][x-1] == "M" && data[y+1][x+1] == "S") || (data[y-1][x-1] == "S" && data[y+1][x+1] == "M") {
				found++
			}
		default:
			continue
		}
	}
	if found == 2 {
		return 1
	}
	return 0
}
