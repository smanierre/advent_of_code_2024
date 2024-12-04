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
