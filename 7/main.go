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

type Operation int

const (
	Add Operation = iota
	Mult
	Concat
)

type Equation struct {
	Answer   int
	Operands []int
}

func (e Equation) GetSolution(concat bool) int {
	if validAdd := e.Solve(e.Operands, Add, 0, concat); validAdd {
		return e.Answer
	}
	if validMul := e.Solve(e.Operands, Mult, 0, concat); validMul {
		return e.Answer
	}
	return 0
}

func (e Equation) Solve(remainingNumbers []int, operation Operation, prevTotal int, concat bool) bool {
	var total int
	if operation == Mult {
		if prevTotal == 0 {
			prevTotal = 1
		}
		total = prevTotal * remainingNumbers[0]
	} else if operation == Add {
		total = remainingNumbers[0] + prevTotal
	} else if operation == Concat {
		prevStr := strconv.Itoa(prevTotal)
		curStr := strconv.Itoa(remainingNumbers[0])
		newStr := prevStr + curStr
		total, _ = strconv.Atoi(newStr)
	}
	if len(remainingNumbers) == 1 {
		return total == e.Answer
	}
	validAdd := e.Solve(remainingNumbers[1:], Add, total, concat)
	validMult := e.Solve(remainingNumbers[1:], Mult, total, concat)
	var validConcat bool
	if concat {
		validConcat = e.Solve(remainingNumbers[1:], Concat, total, concat)
	}
	return validAdd || validMult || validConcat
}

func (e Equation) AllAdd() int {
	total := 0
	for _, o := range e.Operands {
		total += o
	}
	return total
}

func (e Equation) AllTimes() int {
	total := 0
	for i, operand := range e.Operands {
		if i == 0 {
			total += operand
		} else {
			total *= operand
		}
	}
	return total
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
	var equations []Equation

	for _, line := range strings.Split(string(b), "\n") {
		var o []int
		c := strings.Index(line, ":")
		e := Equation{}
		answer, err := strconv.Atoi(line[:c])
		if err != nil {
			log.Fatalf("Error parsing answer: %s", err.Error())
		}
		e.Answer = answer
		for _, v := range strings.Split(strings.TrimSpace(line[c+1:]), " ") {
			n, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalf("Error parsing operand: %s", err.Error())
			}
			o = append(o, n)
		}
		e.Operands = o
		equations = append(equations, e)
	}
	total := 0
	for _, e := range equations {
		total += e.GetSolution(false)
	}
	fmt.Println(time.Since(t))
	fmt.Println(total)
	t = time.Now()
	total = 0
	for _, e := range equations {
		total += e.GetSolution(true)
	}
	fmt.Println(time.Since(t))
	fmt.Println(total)
}
