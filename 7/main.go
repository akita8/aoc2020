package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func findFirst(rules map[string]map[string]int, bagType string, found map[string]bool) int {
	var count int
	for k, v := range rules {
		if v[bagType] != 0 && !found[k] {
			found[k] = true
			count += 1 + findFirst(rules, k, found)
		}
	}
	return count
}

func findSecond(rules map[string]map[string]int, bagType string) int {
	var count int
	for k, v := range rules[bagType] {
		count += v
		res := findSecond(rules, k)
		if res != 0 {
			count += v * res
		}
	}
	return count
}
func first(rules map[string]map[string]int) {
	log.Println(findFirst(rules, "shiny gold", make(map[string]bool)))
}

func second(rules map[string]map[string]int) {
	log.Println(findSecond(rules, "shiny gold"))
}

func main() {
	f, err := os.Open("./7.txt")
	if err != nil {
		log.Fatalf("unable to open file %+v", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	rules := make(map[string]map[string]int)
	for scanner.Scan() {
		line := scanner.Text()
		r := strings.NewReplacer(" bags", "", " bag", "", ".", "", ",", "")
		line = r.Replace(line)
		components := strings.Split(line, " contain ")
		key := components[0]
		contained := strings.Split(components[1], " ")
		if contained[0] != "no" {
			rules[key] = make(map[string]int)
			for i := 0; i < len(contained); i += 3 {
				count, err := strconv.Atoi(contained[i])
				if err != nil {
					log.Fatalf("unable to convert %s to int: %v", contained[i], err)
				}
				rules[key][fmt.Sprintf("%s %s", contained[i+1], contained[i+2])] = count
			}
		}
	}

	first(rules)
	second(rules)
}
