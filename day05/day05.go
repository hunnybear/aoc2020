//usr/bin/env go run $0 $@ ; exit

// part 1: first 7 chars are front-back, last 3 are side-side
// find 2d pos (row, col), sead ID is row * 8 + col
// looking for highest seat id
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
	"unicode"
)

const ROWS int = 128
const COLS int = 8

var offsets = map[byte]int{
	70:0, 66: 1, 76: 0, 82: 1,
}

func partition(size int, instructions string) int {
	offset := 0
	for i := 0; i < len(instructions); i++ {
		size /= 2
		offset += (size * offsets[instructions[i]]) 
	}
	if size != 1 {
		log.Fatal("something is wrong, did not pinpoint")
	}
	return offset
}

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {log.Fatal(err)}

	highest := 0
	mine := -1

	lines := strings.Split(strings.TrimFunc(string(input), unicode.IsSpace), "\n")
	ids := make([]int, len(lines))

	for line_i := 0; line_i < len(lines); line_i++ {
		row := partition(ROWS, lines[line_i][0:7])
		col := partition(COLS, lines[line_i][7:len(lines[line_i])])
	
		id := row * 8 + col
		ids[line_i] = id
		if id > highest {
			highest = id
		}
	}
	sort.Ints(ids)
	for id_i := 1; id_i < len(ids); id_i++ {
		if ids[id_i] - ids[id_i-1] != 1 {
			if mine != -1 {
				log.Fatal("already found mine!")
			}
			mine = ids[id_i] - 1
		}
	}
	fmt.Println(mine)
	fmt.Println(highest)
}
