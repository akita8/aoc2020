package main

import (
	"bufio"
	"log"
	"os"
)

func first(right, down int) int {
	// uglyyyyyyyyyy
	f, err := os.Open("./3.txt")
	if err != nil {
		log.Fatalf("unable to open file %+v", err)
	}
	defer f.Close()
	var treeCount, rowLength, x, y int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := scanner.Text()

		if y%down != 0 || rowLength == 0 {
			rowLength = len(row)
			y++
			continue
		}

		x += right
		if x >= rowLength {
			x -= rowLength
		}
		if rune(row[x]) == '#' {
			treeCount++
		}
		y++
	}
	return treeCount
}

func second() int {
	count := 1
	params := []struct {
		right, down int
	}{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	for _, p := range params {
		count *= first(p.right, p.down)
	}
	return count
}

func main() {

	log.Println(first(3, 1))
	log.Println(second())
}
