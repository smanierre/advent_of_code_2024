package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err.Error())
	}
	reports := strings.Split(string(b), "\n")
	safeReports := 0
	for _, report := range reports {
		nums := convertToNums(report)
		direction, err := getDirection(nums)
		if err != nil {
			continue
		}
		if isSafe(nums, direction, true) {
			safeReports++
		}
	}
	fmt.Println(safeReports)
}

type Direction interface {
	IsPairOk(current, next int) bool
}

type Up struct{}

func (u Up) IsPairOk(current, next int) bool {
	return next-current < 4 && next-current > 0
}

type Down struct{}

func (d Down) IsPairOk(current, next int) bool {
	return current-next < 4 && current-next > 0
}

func getDirection(report []int) (Direction, error) {
	numUp := 0
	numDown := 0
	prev := 0
	duplicate := 0
	for i, v := range report {
		if i == 0 {
			prev = v
			continue
		}
		if v == prev {
			duplicate++
			continue
		}
		if v > prev {
			numUp++
		} else {
			numDown++
		}
		prev = v
	}
	if numUp > 1 && numDown > 1 {
		return nil, errors.New("too many direction changes, not safe")
	}
	if numUp > 0 && numDown > 0 && duplicate > 0 {
		return nil, errors.New("too many direction changes with duplicate, not safe")
	}
	if numUp > 1 {
		return Up{}, nil
	}
	return Down{}, nil
}

func convertToNums(report string) []int {
	parts := strings.Split(report, " ")
	var nums []int
	for _, part := range parts {
		n, err := strconv.Atoi(part)
		if err != nil {
			log.Fatal(err.Error())
		}
		nums = append(nums, n)
	}
	return nums
}

func isSafe(nums []int, d Direction, recurse bool) bool {
	previous := nums[0]
	for i := 1; i < len(nums); i++ {
		if !d.IsPairOk(previous, nums[i]) {
			if !recurse {
				return false
			}
			var withoutPrevious []int
			var withoutCurrent []int
			// Need to do it like this to not modify existing slice
			withoutPrevious = append(withoutPrevious, nums[:i-1]...)
			withoutPrevious = append(withoutPrevious, nums[i:]...)
			withoutCurrent = append(withoutCurrent, nums[:i]...)
			withoutCurrent = append(withoutCurrent, nums[i+1:]...)
			if isSafe(withoutCurrent, d, false) || isSafe(withoutPrevious, d, false) {
				return true
			} else {
				return false
			}
		}
		previous = nums[i]
	}
	return true
}
