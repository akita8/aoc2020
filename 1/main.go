package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// bruteforce!
func first(numbers []int) {
	for _, n := range numbers {
		for _, o := range numbers {
			if n+o == 2020 {
				log.Println(n * o)
				return
			}
		}
	}
}

// bruteforce**3!
func second(numbers []int) {
	for _, n := range numbers {
		for _, o := range numbers {
			for _, oo := range numbers {
				if n+o+oo == 2020 {
					log.Println(n * o * oo)
					return
				}
			}
		}
	}
}

func main() {
	f, err := os.Open("./1.txt")
	if err != nil {
		log.Fatalf("unable to open file %+v", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var numbers []int
	for scanner.Scan() {
		num := scanner.Text()
		i, err := strconv.Atoi(num)
		if err != nil {
			log.Fatalf("unable to convert to int %s :%+v", num, err)
		}
		numbers = append(numbers, i)
	}
	first(numbers)
	second(numbers)
}
