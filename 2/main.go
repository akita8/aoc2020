package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type row struct {
	lower  int
	upper  int
	target rune
	passwd string
}

func first(rows []row) {
	var validCount int
	for _, r := range rows {
		count := 0
		for _, s := range []rune(r.passwd) {
			if r.target == s {
				count++
			}
		}
		if count >= r.lower && count <= r.upper {
			validCount++
		}
	}
	log.Println(validCount)
}

func second(rows []row) {
	var validCount int
	for _, r := range rows {
		letters := []rune(r.passwd)
		if len(letters) >= r.lower && len(letters) >= r.upper {
			// I know double if but I don't want a loooong condition
			if (r.target == letters[r.lower-1]) != (r.target == letters[r.upper-1]) {
				validCount++
			}
		}
	}
	log.Println(validCount)
}

func main() {
	f, err := os.Open("./2.txt")
	if err != nil {
		log.Fatalf("unable to open file %+v", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var rows []row
	for scanner.Scan() {
		s := scanner.Text()
		rowParts := strings.Split(s, " ")
		bounds := strings.Split(rowParts[0], "-")
		lower, err := strconv.Atoi(bounds[0])
		if err != nil {
			log.Fatalf("unable to convert lower bound %s %+v", bounds[0], err)
		}
		upper, err := strconv.Atoi(bounds[1])
		if err != nil {
			log.Fatalf("unable to convert lower bound %s %+v", bounds[1], err)
		}
		r := row{
			target: rune(rowParts[1][0]),
			passwd: rowParts[2],
			lower:  lower,
			upper:  upper,
		}
		rows = append(rows, r)
	}
	first(rows)
	second(rows)
}
