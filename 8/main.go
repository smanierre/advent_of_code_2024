package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type Antenna struct {
	Count       int
	Coordinates []Coordinate
}

type Coordinate struct {
	X int
	Y int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error opening input file: %s", err.Error())
	}
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatalf("Error reading input: %s", err.Error())
	}
	var grid [][]string
	for _, line := range strings.Split(string(b), "\n") {
		grid = append(grid, strings.Split(line, ""))
	}
	antennas := map[string]Antenna{}
	for y, row := range grid {
		for x, cell := range row {
			if cell != "." {
				if a, ok := antennas[cell]; ok {
					a.Count++
					a.Coordinates = append(a.Coordinates, Coordinate{X: x, Y: y})
					antennas[cell] = a
				} else {
					antennas[cell] = Antenna{Count: 1, Coordinates: []Coordinate{{X: x, Y: y}}}
				}
			}
		}
	}
	validAntennas := map[string]Antenna{}
	for k, v := range antennas {
		if v.Count > 1 {
			validAntennas[k] = v
		}
	}
	partOne(len(grid[0]), len(grid), validAntennas)
	partTwo(len(grid[0]), len(grid), validAntennas)
}

func partOne(xBound, yBound int, validAntennas map[string]Antenna) {
	t := time.Now()
	uniqueAntinodes := map[Coordinate]int{}
	for _, v := range validAntennas {
		for _, c1 := range v.Coordinates {
			for _, c2 := range v.Coordinates {
				if c1.X == c2.X && c1.Y == c2.Y {
					continue // Same antenna, skip
				}
				distanceX, distanceY := getDistance(c1, c2)
				antinodes := getAntinodeCoords(c1, c2, distanceX, distanceY)
				for _, a := range antinodes {
					if a.X >= 0 && a.X < xBound && a.Y >= 0 && a.Y < yBound {
						uniqueAntinodes[a] = 1
					}
				}
			}
		}
	}
	fmt.Println(len(uniqueAntinodes))
	fmt.Println(time.Since(t))
}

func partTwo(xBound, yBound int, validAntennas map[string]Antenna) {
	t := time.Now()
	uniqueAntinodes := map[Coordinate]int{}
	for _, v := range validAntennas {
		for _, c1 := range v.Coordinates {
			for _, c2 := range v.Coordinates {
				if c1.X == c2.X && c1.Y == c2.Y {
					continue // Same antenna, skip
				}
				distanceX, distanceY := getDistance(c1, c2)
				antinodes := getResonantAntinodeCoords(c1, c2, distanceX, distanceY, xBound, yBound)
				for _, a := range antinodes {
					if a.X >= 0 && a.X < xBound && a.Y >= 0 && a.Y < yBound {
						uniqueAntinodes[a] = 1
					}
				}
			}
		}
	}
	fmt.Println(len(uniqueAntinodes))
	fmt.Println(time.Since(t))
}

func getDistance(a1, a2 Coordinate) (int, int) {
	return a1.X - a2.X, a1.Y - a2.Y
}

func getAntinodeCoords(c1, c2 Coordinate, distanceX, distanceY int) []Coordinate {
	return []Coordinate{
		{X: c1.X + distanceX, Y: c1.Y + distanceY},
		{X: c2.X - distanceX, Y: c2.Y - distanceY},
	}
}

func getResonantAntinodeCoords(c1, c2 Coordinate, distanceX, distanceY, xBound, yBound int) []Coordinate {
	coords := []Coordinate{c1, c2} // Antenna's are now coordinates as well
	x := c1.X
	y := c1.Y
	for {
		x = x + distanceX
		y = y + distanceY
		if x < 0 || x >= xBound || y < 0 || y >= yBound {
			break
		}
		coords = append(coords, Coordinate{X: x, Y: y})
	}
	for {
		x = x - distanceX
		y = y - distanceY
		if x < 0 || x >= xBound || y < 0 || y >= yBound {
			break
		}
		coords = append(coords, Coordinate{X: x, Y: y})
	}
	return coords
}
