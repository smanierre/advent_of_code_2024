package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

type File struct {
	ID     int
	Length int
}

func (f File) Len() int {
	return f.Length * len(strconv.Itoa(f.ID))
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
	t := time.Now()
	var disk [][]string
	space := false
	id := -1
	for _, v := range string(b) {
		n, err := strconv.Atoi(string(v))
		if err != nil {
			log.Fatalf("Error converting input to integer: %s", err.Error())
		}
		if !space {
			id++
		}
		for range n {
			if space {
				disk = append(disk, []string{"."})
			} else {
				disk = append(disk, []string{strconv.Itoa(id)})
			}
		}
		space = !space
	}

	var disk2 []File
	id = -1
	space = false
	for _, v := range string(b) {
		n, err := strconv.Atoi(string(v))
		if err != nil {
			log.Fatalf("Error converting input to integer: %s", err.Error())
		}
		if !space {
			id++
		}
		if space {
			for range n {
				disk2 = append(disk2, File{ID: -1, Length: 1})
			}
		} else {
			disk2 = append(disk2, File{ID: id, Length: n})
		}
		space = !space
	}
	fmt.Printf("Preparation time: %s\n", time.Since(t))

	partOne(disk)
	partTwo(disk2)
}

func partOne(disk [][]string) {
	t := time.Now()
	for i := len(disk) - 1; i > 0; i-- {
		if disk[i][0] == "." {
			continue
		}
		slot := getFirstSpace(disk, len(disk[i]))
		if slot == -1 {
			continue
		}
		if slot > i {
			continue
		}
		tmp := append(disk[:slot], disk[i])
		tmp = append(tmp, disk[slot+len(disk[i]):i]...)
		for range len(disk[i]) {
			tmp = append(tmp, []string{"."})
		}
		if i+1 == len(disk) {
			disk = tmp
			continue
		}
		tmp = append(tmp, disk[i+1:]...)
		disk = tmp
	}
	total := 0
	for i, v := range disk {
		n, _ := strconv.Atoi(string(v[0]))
		total += i * n
	}
	fmt.Println(time.Since(t))
	fmt.Println(total)
}

func partTwo(disk []File) {
	t := time.Now()
	lastId := 9999
	for i := len(disk) - 1; i > 0; i-- {
		if disk[i].ID == -1 {
			continue
		}
		slot := getFirstFileSpace(disk, disk[i])
		if slot == -1 {
			continue
		}
		if slot > i {
			continue
		}
		if disk[i].ID > lastId {
			continue
		}
		tmp := append(disk[:slot], disk[i])
		tmp = append(tmp, disk[slot+disk[i].Len():i]...)
		for range disk[i].Length * len(strconv.Itoa(disk[i].ID)) {
			tmp = append(tmp, File{ID: -1, Length: 1})
		}
		if i+1 == len(disk) {
			disk = tmp
			continue
		}
		tmp = append(tmp, disk[i+1:]...)
		disk = tmp
	}
	total := 0
	index := 0
	for _, v := range disk {
		if v.ID == -1 {
			index++
			continue
		}
		total += v.ID * index
		index += v.Len()
	}
	fmt.Println(time.Since(t))
	fmt.Println(total)
}

func getFirstSpace(disk [][]string, size int) int {
	count := 0
	index := 0
	for i, v := range disk {
		if v[0] == "." {
			if count == 0 {
				count++
				index = i

			} else {
				count++
			}
			if count == size {
				return index
			}
		}
	}
	return -1
}

func getFirstFileSpace(disk []File, f File) int {
	count := 0
	index := 0
	for i, v := range disk {
		if v.ID == -1 {
			if count == 0 {
				count++
				index = i
			} else {
				count++
			}
			if count == f.Len() {
				return index
			}
		} else {
			count = 0
		}
	}
	return -1
}
