package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
)

func first(codes []string) {

	var maxID int
	for _, code := range codes {
		var row, column int
		for i, c := range code[:7] {
			if c == 'B' {
				row += int(math.Pow(2, 6.-float64(i)))
			}
		}
		for i, c := range code[7:] {
			if c == 'R' {
				column += int(math.Pow(2, 2.-float64(i)))
			}
		}
		ID := row*8 + column
		if ID > maxID {
			maxID = ID
		}
	}
	log.Println(maxID)
}

func second(codes []string) {
	var IDs []int
	for _, code := range codes {
		var row, column int
		for i, c := range code[:7] {
			if c == 'B' {
				row += int(math.Pow(2, 6.-float64(i)))
			}
		}
		for i, c := range code[7:] {
			if c == 'R' {
				column += int(math.Pow(2, 2.-float64(i)))
			}
		}
		IDs = append(IDs, row*8+column)

	}
	sort.Slice(IDs, func(i, j int) bool { return IDs[i] < IDs[j] })
	for i, id := range IDs {
		if i+IDs[0] != id {
			log.Println(id - 1)
			break
		}
	}
}

func main() {
	f, err := os.Open("./5.txt")
	if err != nil {
		log.Fatalf("unable to open file %+v", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var codes []string
	for scanner.Scan() {
		codes = append(codes, scanner.Text())
	}
	first(codes)
	second(codes)
}
