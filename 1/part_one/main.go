package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	bytes, err := io.ReadAll(input)
	if err != nil {
		log.Fatal(err.Error())
	}
	items := strings.Split(string(bytes), "\n")
	var listOne, listTwo []int
	for _, i := range items {
		pair := strings.Split(i, "   ")
		if len(pair) != 2 {
			continue
		}
		itemOne, err := strconv.Atoi(pair[0])
		if err != nil {
			log.Fatal(err.Error())
		}
		itemTwo, err := strconv.Atoi(pair[1])
		if err != nil {
			log.Fatal(err.Error())
		}
		listOne = append(listOne, itemOne)
		listTwo = append(listTwo, itemTwo)
	}

	slices.Sort(listOne)
	slices.Sort(listTwo)

	var distances []int
	for i := range len(listOne) {
		distances = append(distances, int(math.Abs(float64(listOne[i]-listTwo[i]))))
	}
	total := 0
	for _, d := range distances {
		total += d
	}
	fmt.Println(total)
}
