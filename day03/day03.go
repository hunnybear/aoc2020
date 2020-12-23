//usr/bin/env go run $0 $@ ; exit
// 1934917632 is too low

// right 3, down 1
package main

import (
	"bytes"
	"fmt"
	"log"
//	"io"
	"io/ioutil"
	"unicode"
)

var open byte = '.'
var tree byte = '#'

type slope struct {
	x int
	y int
}

func score_run(the_map [][]byte, my_slope slope) int {
	trees := 0
	x := 0
	
	for y := 0; y < len(the_map); y += my_slope.y {
		if the_map[y][x % len(the_map[y])] == tree {
			trees ++
		}

		x += my_slope.x
	}

	return trees
}

func main() {

	input, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	part_2_slopes := [5]slope { 
		slope{x:1,y:1},
		slope{x:3,y:1},
		slope{x:5,y:1},
		slope{x:7,y:1},
		slope{x:1,y:2},
	}

	the_map := bytes.Split(bytes.TrimFunc(input, unicode.IsSpace), []byte("\n"))

	part_1_trees := score_run(the_map, slope{x:1, y:3})
	
	// totals are mulitiplied cumulatively, so start with 1 rather than
	// 0.
	part_2_answer := 1

	for i := 0; i < len(part_2_slopes); i++ {
		run_score := score_run(the_map, part_2_slopes[i])
		part_2_answer *= run_score
	}

	fmt.Println("part 01")
	fmt.Println(part_1_trees)
	fmt.Println("part 02")
	fmt.Println(part_2_answer)
}