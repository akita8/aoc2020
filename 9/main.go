package main

import (
	"bufio"
	"errors"
	"log"
	"math"
	"os"
	"strconv"
)

func first(numbers []int, preambleLength int) (int, error) {
	for i := preambleLength; i < len(numbers); i++ {
		num := numbers[i]
		invalid := true
		preamble := numbers[i-preambleLength : i]
		for j := 0; j < preambleLength-1; j++ {
			for k := j + 1; k < len(preamble); k++ {
				if preamble[j]+preamble[k] == num {
					invalid = false
					break
				}
			}
		}
		if invalid {
			return num, nil
		}
	}

	return 0, errors.New("no valid solutions :(")
}

func second(numbers []int, invalidNumber int) (int, error) {
	var set []int
	for i := 2; i < len(numbers)-1; i++ {
		for j := 0; i+j < len(numbers); j++ {
			sum := 0
			for _, num := range numbers[j : i+j] {
				sum += num
			}
			if sum == invalidNumber {
				set = numbers[j : i+j]
				break
			}
		}
	}
	if set == nil {
		return 0, errors.New("no valid solutions :(")
	}
	min := math.MaxInt64
	max := 0
	for _, num := range set {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	return min + max, nil
}

func main() {
	f, err := os.Open("./9.txt")
	if err != nil {
		log.Fatalf("unable to open file %+v", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var numbers []int
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("unable to convert %s to int: %+v", line, err)
		}
		numbers = append(numbers, num)
	}
	invalidNumber, err := first(numbers, 25)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(invalidNumber)
	weakness, err := second(numbers, invalidNumber)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(weakness)
}
