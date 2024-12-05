package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Rules []Rule

var totalTime = time.Now()

func (r Rules) IsApplicable(pages []int) Rules {
	foundRules := map[Rule]int{}
	for _, rule := range r {
		if slices.Contains(pages, rule.First) && slices.Contains(pages, rule.Second) {
			foundRules[rule]++
		}
	}
	newRules := Rules{}
	for rule, _ := range foundRules {
		newRules = append(newRules, rule)
	}
	return newRules
}

func (r Rules) DoesConform(pages []int, rules Rules) bool {
	good := true
	for _, rule := range rules {
		if !rule.IsFollowed(pages) {
			good = false
			break
		}
	}
	return good
}

func (r Rules) Correct(pages []int, c chan<- []int) []int {
	good := false
	for !good {
		for _, rule := range r {
			firstIndex := slices.Index(pages, rule.First)
			secondIndex := slices.Index(pages, rule.Second)
			if firstIndex > secondIndex {
				pages[firstIndex], pages[secondIndex] = pages[secondIndex], pages[firstIndex]
			}
			if r.DoesConform(pages, r) {
				good = true
			}
		}
	}
	c <- pages
	return pages
}

type Rule struct {
	First  int
	Second int
}

func (r Rule) IsFollowed(update []int) bool {
	firstIndex := slices.Index(update, r.First)
	secondIndex := slices.Index(update, r.Second)
	return firstIndex < secondIndex
}

func NewRule(input string) Rule {
	parts := strings.Split(input, "|")
	if len(parts) != 2 {
		log.Fatalf("Got rule with invalid number of inputs: %d", len(parts))
	}
	first, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatalf("Got rule with invalid first number: %s", err.Error())
	}
	second, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatalf("Got rule with invalid second number: %s", err.Error())
	}
	return Rule{
		First:  first,
		Second: second,
	}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error opening input file: %s", err.Error())
	}
	defer f.Close()
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatalf("Error reading input: %s", err.Error())
	}

	var rules Rules
	var updates [][]int
	toRules := true
	for _, line := range strings.Split(string(b), "\n") {
		if line == "" {
			toRules = false
			continue
		}
		if toRules {
			rules = append(rules, NewRule(line))
		} else {
			pages := strings.Split(line, ",")
			var pageNums []int
			for _, page := range pages {
				pageNum, err := strconv.Atoi(page)
				if err != nil {
					log.Fatalf("Error converting page number to int: %s", err.Error())
				}
				pageNums = append(pageNums, pageNum)
			}
			updates = append(updates, pageNums)
		}
	}

	var validUpdates [][]int
	var invalidUpdates [][]int
	for _, update := range updates {
		if !rules.DoesConform(update, rules.IsApplicable(update)) {
			invalidUpdates = append(invalidUpdates, update)
			continue
		}
		validUpdates = append(validUpdates, update)
	}

	validTotal := 0
	for _, v := range validUpdates {
		middle := (len(v) - 1) / 2
		validTotal += v[middle]
	}
	fmt.Printf("Valid total: %d\n", validTotal)

	invalidTotal := 0

	returnVals := make(chan []int)
	runningRoutines := 0
	for _, v := range invalidUpdates {
		applicableRules := rules.IsApplicable(v)
		runningRoutines++
		go applicableRules.Correct(v, returnVals)
	}

	totalTries := 0
	c := time.Tick(time.Minute)
	go func() {
		for {
			select {
			case <-c:
				fmt.Printf("%d left!\n", runningRoutines)
				fmt.Printf("Total tries: %d\n", totalTries)
				fmt.Printf("Tries / second: %f\n", float64(totalTries)/time.Since(totalTime).Seconds())
			}
		}
	}()
	for runningRoutines > 0 {
		select {
		case fixed := <-returnVals:
			invalidTotal += fixed[(len(fixed)-1)/2]
			runningRoutines--
		default:
			totalTries++
		}
	}
	fmt.Printf("InvalidTotal: %d\n", invalidTotal)
	fmt.Printf("Total Time: %s\n", time.Since(totalTime).String())
}
