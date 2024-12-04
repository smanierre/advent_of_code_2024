package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("../part_one/input.txt")
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
	occurrences := map[int]int{}
	for _, i := range listTwo {
		if _, ok := occurrences[i]; !ok {
			occurrences[i] = 1
		} else {
			occurrences[i]++
		}
	}
	similarity := 0
	for _, i := range listOne {
		if _, ok := occurrences[i]; ok {
			similarity += i * occurrences[i]
		}
	}
	fmt.Println(similarity)
}
