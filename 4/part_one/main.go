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

func IsXMAS(x, y int, data [][]string) int {
	if data[y][x] != "X" {
		return 0
	}
	dirs := GetSafeDirs(x, y, len(data[0]), len(data))
	return FindXMAS(x, y, dirs, data)
}

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
	UpRight
	UpLeft
	DownRight
	DownLeft
)

func GetSafeDirs(x, y, width, height int) []Direction {
	dirs := []Direction{}
	// Check up, need at least 3 rows above
	if y > 2 {
		dirs = append(dirs, Up)
	}
	// Check Down, need at least 3 rows below
	if y < height-3 {
		dirs = append(dirs, Down)
	}
	// Check Left, need at least 3 columns before
	if x > 2 {
		dirs = append(dirs, Left)
	}
	// Check Right, need at least 3 columns after
	if x < width-3 {
		dirs = append(dirs, Right)
	}
	// Check up right, need 3 rows above and 3 columns after
	if x < width-3 && y > 2 {
		dirs = append(dirs, UpRight)
	}
	// Check up Left, need 3 rows above and 3 columns before
	if x > 2 && y > 2 {
		dirs = append(dirs, UpLeft)
	}
	// Check Down right, need 3 rows below and 3 columns after
	if x < width-3 && y < height-3 {
		dirs = append(dirs, DownRight)
	}
	// Check Down left, need 3 rows below and 3 columns before
	if x > 2 && y < height-3 {
		dirs = append(dirs, DownLeft)
	}
	return dirs
}

func FindXMAS(x, y int, dirs []Direction, data [][]string) int {
	found := 0
	for _, dir := range dirs {
		switch dir {
		case Up:
		loopUp:
			for i := 1; i < 4; i++ {
				switch i {
				case 1:
					if data[y-i][x] != "M" {
						break loopUp
					}
				case 2:
					if data[y-i][x] != "A" {
						break loopUp
					}
				case 3:
					if data[y-i][x] == "S" {
						found++
					}
				}
			}
		case Down:
		loopDown:
			for i := 1; i < 4; i++ {
				switch i {
				case 1:
					if data[y+i][x] != "M" {
						break loopDown
					}
				case 2:
					if data[y+i][x] != "A" {
						break loopDown
					}
				case 3:
					if data[y+i][x] == "S" {
						found++
					}
				}
			}
		case Left:
		loopLeft:
			for i := 1; i < 4; i++ {
				switch i {
				case 1:
					if data[y][x-i] != "M" {
						break loopLeft
					}
				case 2:
					if data[y][x-i] != "A" {
						break loopLeft
					}
				case 3:
					if data[y][x-i] == "S" {
						found++
					}
				}
			}
		case Right:
		loopRight:
			for i := 1; i < 4; i++ {
				switch i {
				case 1:
					if data[y][x+i] != "M" {
						break loopRight
					}
				case 2:
					if data[y][x+i] != "A" {
						break loopRight
					}
				case 3:
					if data[y][x+i] == "S" {
						found++
					}
				}
			}
		case UpRight:
		loopUpRight:
			for i := 1; i < 4; i++ {
				switch i {
				case 1:
					if data[y-i][x+i] != "M" {
						break loopUpRight
					}
				case 2:
					if data[y-i][x+i] != "A" {
						break loopUpRight
					}
				case 3:
					if data[y-i][x+i] == "S" {
						found++
					}
				}
			}
		case UpLeft:
		loopUpLeft:
			for i := 1; i < 4; i++ {
				switch i {
				case 1:
					if data[y-i][x-i] != "M" {
						break loopUpLeft
					}
				case 2:
					if data[y-i][x-i] != "A" {
						break loopUpLeft
					}
				case 3:
					if data[y-i][x-i] == "S" {
						found++
					}
				}
			}
		case DownRight:
		loopDownRight:
			for i := 1; i < 4; i++ {
				switch i {
				case 1:
					if data[y+i][x+i] != "M" {
						break loopDownRight
					}
				case 2:
					if data[y+i][x+i] != "A" {
						break loopDownRight
					}
				case 3:
					if data[y+i][x+i] == "S" {
						found++
					}
				}
			}
		case DownLeft:
		loopDownLeft:
			for i := 1; i < 4; i++ {
				switch i {
				case 1:
					if data[y+i][x-i] != "M" {
						break loopDownLeft
					}
				case 2:
					if data[y+i][x-i] != "A" {
						break loopDownLeft
					}
				case 3:
					if data[y+i][x-i] == "S" {
						found++
					}
				}
			}
		}
	}
	return found
}
