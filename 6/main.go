package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func first(f *os.File) {
	scanner := bufio.NewScanner(f)
	var count int

	for scanner.Scan() {
		answers := make(map[rune]bool)
		for {
			s := scanner.Text()
			if len(s) == 0 {
				break
			}
			for _, c := range []rune(s) {
				answers[c] = true
			}
			if !scanner.Scan() {
				break
			}
		}
		count += len(answers)
	}

	log.Println(count)
}

func second(f *os.File) {
	scanner := bufio.NewScanner(f)
	var count int

	for scanner.Scan() {
		lineCount := 0
		answers := make(map[rune]int)
		for {
			s := scanner.Text()
			if len(s) == 0 {
				break
			}
			for _, c := range []rune(s) {
				answers[c] += 1
			}
			lineCount++
			if !scanner.Scan() {
				break
			}
		}
		for _, v := range answers {
			if v == lineCount {
				count++
			}
		}

	}

	log.Println(count)
}

func main() {
	f, err := os.Open("./6.txt")
	if err != nil {
		log.Fatalf("unable to open file %+v", err)
	}
	defer f.Close()

	first(f)
	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		log.Fatalf("unable to rewind file pointer: %+v", err)
	}
	second(f)
}
