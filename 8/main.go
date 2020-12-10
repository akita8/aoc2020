package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	operation string
	parameter int
}

func (i *instruction) flip() bool {
	switch i.operation {
	case "nop":
		i.operation = "jmp"
		return true
	case "acc":
		return false
	case "jmp":
		i.operation = "nop"
		return true
	default:
		log.Fatalf("got unexpected operation %s", i.operation)
		return false
	}
}

func first(instructions []*instruction) (int, bool) {
	var accumulator, i int
	valid := true
	visited := make(map[int]bool)

	for {
		if i == len(instructions) {
			break
		}
		if visited[i] {
			valid = false
			break
		}
		inst := instructions[i]
		visited[i] = true
		switch inst.operation {
		case "nop":
			i++
		case "acc":
			accumulator += inst.parameter
			i++
		case "jmp":
			i += inst.parameter
		default:
			log.Fatalf("got unexpected operation %s", inst.operation)
		}
	}

	return accumulator, valid
}

func second(instructions []*instruction) (int, bool) {
	for i := 0; i < len(instructions); i++ {
		if instructions[i].flip() {
			acc, valid := first(instructions)
			if valid {
				return acc, valid
			}
			instructions[i].flip()
		}
	}
	return 0, false
}

func main() {
	f, err := os.Open("./8.txt")
	if err != nil {
		log.Fatalf("unable to open file %+v", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var instructions []*instruction
	for scanner.Scan() {
		line := scanner.Text()
		components := strings.Split(line, " ")
		parameter, err := strconv.Atoi(components[1])
		if err != nil {
			log.Fatalf("unable to convert %s to int: %+v", components[1], err)
		}
		instructions = append(instructions, &instruction{
			operation: components[0],
			parameter: parameter,
		})
	}
	acc, _ := first(instructions)
	log.Println(acc)
	acc, valid := second(instructions)
	if valid {
		log.Println(acc)
	} else {
		log.Println("no valid solutions :(")
	}

}
