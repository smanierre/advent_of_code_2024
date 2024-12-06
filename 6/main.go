package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type Direction int

func (d Direction) NextDirection() Direction {
	switch d {
	case Up:
		return Right
	case Down:
		return Left
	case Left:
		return Up
	case Right:
		return Down
	}
	return d
}

func (d Direction) NextCoordinate(current *Coordinate, xLen, yLen int) error {
	switch d {
	case Up:
		current.Y = current.Y - 1
	case Down:
		current.Y = current.Y + 1
	case Left:
		current.X = current.X - 1
	case Right:
		current.X = current.X + 1
	}
	// Check for next Coordinate out of bounds on the x-axis
	if current.X < 0 || current.X > xLen-1 {
		return errors.New("off grid")
	}
	// Check for next Coordinate out of bounds on the y-axis
	if current.Y < 0 || current.Y > yLen-1 {
		return errors.New("off grid")
	}
	return nil
}

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Coordinate struct {
	X      int
	Y      int
	Facing Direction
}

func (c *Coordinate) NextObstacle(room [][]string) (int, bool) {
	totalMoves := 0
	for {
		room[c.Y][c.X] = "X"
		err := c.Facing.NextCoordinate(c, len(room[0]), len(room))
		if err != nil {
			return totalMoves, true
		}
		if room[c.Y][c.X] == "#" {
			c.Backtrack()
			c.Facing = c.Facing.NextDirection()
			return totalMoves, false
		}
		if room[c.Y][c.X] != "X" {
			totalMoves++
		}
	}
}

func (c *Coordinate) Backtrack() {
	switch c.Facing {
	case Up:
		c.Y = c.Y + 1
	case Down:
		c.Y = c.Y - 1
	case Left:
		c.X = c.X + 1
	case Right:
		c.X = c.X - 1
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
	var room [][]string
	for _, line := range strings.Split(string(b), "\n") {
		room = append(room, strings.Split(line, ""))
	}
	newRoomFunc := NewRoomPrep(room)

	position := FindStart(room)
	count := 1 // Count the starting position
	done := false
	for !done {
		distanceTraveled, finished := position.NextObstacle(room)
		count += distanceTraveled
		done = finished
	}
	fmt.Printf("Spots visited: %d\n", count)
	fmt.Printf("Part 1 took: %s\n", time.Since(t))

	t = time.Now()
	loops := 0
	for i, row := range room {
		for j, _ := range row {
			newRoom := newRoomFunc()
			// Check to see if there's an obstacle there, if so skip
			if newRoom[i][j] == "#" {
				continue
			}
			// Never went there before, so not possible
			if room[i][j] == "." {
				continue
			}
			position := FindStart(newRoom)
			if position.Y == i && position.X == j {
				continue
			}
			newRoom[i][j] = "#"

			for {
				loop := LoopChecker(position, newRoom)
				if loop {
					loops++
					break
				}
				break
			}
		}
	}
	fmt.Printf("Loops Possible: %d\n", loops)
	fmt.Println(time.Since(t))
}

func FindStart(room [][]string) *Coordinate {
	position := &Coordinate{}
	for i, _ := range room {
		for j, _ := range room[i] {
			switch room[i][j] {
			case "^":
				position.Y = i
				position.X = j
				position.Facing = Up
				return position
			case "v":
				position.Y = i
				position.X = j
				position.Facing = Down
				return position
			case ">":
				position.Y = i
				position.X = j
				position.Facing = Right
				return position
			case "<":
				position.Y = i
				position.X = j
				position.Facing = Left
				return position
			}
		}
	}
	return nil
}

func NewRoomPrep(room [][]string) func() [][]string {
	r := make([][]string, len(room))
	for i := range room {
		r[i] = append([]string{}, room[i]...)
	}
	return func() [][]string {
		j := make([][]string, len(r))
		for i := range r {
			j[i] = append([]string{}, r[i]...)
		}
		return j
	}
}

func LoopChecker(coord *Coordinate, room [][]string) bool {
	var visitedCoords []Coordinate
	for {
		err := coord.Facing.NextCoordinate(coord, len(room[0]), len(room))
		if err != nil {
			return false
		}
		if room[coord.Y][coord.X] == "#" {
			coord.Backtrack()
			coord.Facing = coord.Facing.NextDirection()
			continue
		}
		if room[coord.Y][coord.X] == "." {
			room[coord.Y][coord.X] = "X"
			visitedCoords = []Coordinate{{coord.X, coord.Y, coord.Facing}}
			continue
		}
		if visitedCoords[0].X == coord.X && visitedCoords[0].Y == coord.Y {
			// Back where we started with no new coords, its a loop
			return true
		}
		visitedCoords = append(visitedCoords, Coordinate{coord.X, coord.Y, coord.Facing})
	}
}
