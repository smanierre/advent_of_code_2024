package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatalf("Error opening file: %s", err.Error())
	}
	defer f.Close()
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatalf("Error reading file: %s", err.Error())
	}

	dontRegex, err := regexp.Compile(`don't\(\)`)
	if err != nil {
		log.Fatalf("Error compiling regex: %s", err.Error())
	}
	doRegex, err := regexp.Compile(`do\(\)`)
	if err != nil {
		log.Fatalf("Error compiling regex: %s", err.Error())
	}
	done := false
	for !done {
		// Get index of next do/don't
		dontIndex := dontRegex.FindIndex(b)
		doIndex := doRegex.FindIndex(b)
		// If there's no more don'ts, we're done
		if dontIndex == nil {
			done = true
			continue
		}
		// If there is a don't but no do, cut off everything after the don't
		if doIndex == nil {
			b = b[:dontIndex[0]]
			done = true
			continue
		}
		// If there is a do before a don't, remove it as the space between all previous don't-do has been removed
		if doIndex[1] < dontIndex[0] {
			b = append(b[:doIndex[0]], b[doIndex[1]:]...)
		} else { // Remove the area between the don't and do
			b = append(b[:dontIndex[0]], b[doIndex[1]:]...)
		}
	}

	// Copied from Part 1 since we are now just doing the same thing on the remaining
	regex, err := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)
	if err != nil {
		log.Fatalf("Error compiling regex: %s", err.Error())
	}
	results := regex.FindAllSubmatch(b, -1)
	total := 0
	for _, result := range results {
		var numOne, numTwo int
		for i := 1; i < 3; i++ {

			if i == 1 {
				numOne, err = strconv.Atoi(string(result[i]))
			} else {
				numTwo, err = strconv.Atoi(string(result[i]))
			}
		}
		total += numOne * numTwo
	}
	fmt.Println(total)
}
