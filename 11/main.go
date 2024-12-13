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

func main() {
	t := time.Now()
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error opening input file: %v\n", err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatalf("Error reading input file: %v\n", err)
	}
	strNums := strings.Split(string(b), " ")
	var nums []uint64
	for _, s := range strNums {
		n, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			log.Fatalf("Error converting num: %v\n", err)
		}
		nums = append(nums, n)
	}
	partOneArray := make([]uint64, 0)
	partOneArray = append(partOneArray, nums...)
	fmt.Printf("Setup time: %s\n", time.Since(t))
	partOne(partOneArray)
	partTwo(nums)
}

func partOne(nums []uint64) {
	t := time.Now()
	var tmpArray []uint64
	for _, n := range nums {
		tmpArray = append(tmpArray, Blink(n, 25)...)
	}
	fmt.Printf("Part One: %s\n", time.Since(t))
	fmt.Println(len(tmpArray))
}

func partTwo(nums []uint64) {
	t := time.Now()
	var total uint64
	for _, n := range nums {
		CachedBlink(n, 5000, &total)
	}
	fmt.Printf("Part Two: %s\n", time.Since(t))
	fmt.Println(total)
}

func Blink(i, count uint64) []uint64 {
	if count == 0 {
		return []uint64{i}
	}
	newNums := []uint64{}
	if i == 0 {
		newNums = append(newNums, 1)
	}
	if s := strconv.FormatUint(i, 10); len(s)%2 == 0 {
		n1, err := strconv.ParseUint(s[:(len(s)/2)], 10, 64)
		if err != nil {
			log.Fatalf("Error converting num: %v\n", err)
		}
		s2 := s[len(s)/2:]
		s2 = strings.TrimLeft(s2, "0")
		if s2 == "" {
			s2 = "0"
		}
		n2, err := strconv.ParseUint(s2, 10, 64)
		if err != nil {
			log.Fatalf("Error converting num: %v\n", err)
		}
		newNums = append(newNums, Blink(n1, count-1)...)
		newNums = append(newNums, Blink(n2, count-1)...)
		return newNums
	}
	newNums = append(newNums, i*2024)
	return Blink(newNums[0], count-1)
}

type CacheEntry struct {
	Value           uint64
	RemainingBlinks uint64
}

var cache = map[CacheEntry]uint64{}

func CachedBlink(i, count uint64, total *uint64) []uint64 {
	if count == 0 {
		*total += 1
		return []uint64{}
	}
	newNums := []uint64{}
	if i == 0 {
		return CachedBlink(1, count-1, total)
	}
	c := CacheEntry{
		Value:           i,
		RemainingBlinks: count,
	}
	if v, ok := cache[c]; ok {
		*total += v
		return []uint64{}
	}
	if s := strconv.FormatUint(i, 10); len(s)%2 == 0 {
		n1, err := strconv.ParseUint(s[:(len(s)/2)], 10, 64)
		if err != nil {
			log.Fatalf("Error converting num: %v\n", err)
		}
		s2 := s[len(s)/2:]
		s2 = strings.TrimLeft(s2, "0")
		if s2 == "" {
			s2 = "0"
		}
		n2, err := strconv.ParseUint(s2, 10, 64)
		if err != nil {
			log.Fatalf("Error converting num: %v\n", err)
		}
		var tmpTotal uint64 = 0
		newNums = append(newNums, CachedBlink(n1, count-1, &tmpTotal)...)
		newNums = append(newNums, CachedBlink(n2, count-1, &tmpTotal)...)
		if c.RemainingBlinks > 2 {
			cache[c] = tmpTotal
		}
		*total += tmpTotal
		return newNums
	}
	newNums = append(newNums, i*2024)
	return CachedBlink(newNums[0], count-1, total)
}
