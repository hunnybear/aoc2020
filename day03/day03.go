//usr/bin/env go run $0 $@ ; exit

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
	x int8
	y int8
}

func score_run(the_map [][]byte, slope_x int, slope_y int) int {
	trees := 0
	x := 0
	
	for y := 0; y < len(the_map); y += slope_y {
		if the_map[y][x % len(the_map[y])] == tree {
			trees ++
		}

		x += slope_x
	}

	return trees
}

func main() {

	input, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	the_map := bytes.Split(bytes.TrimFunc(input, unicode.IsSpace), []byte("\n"))

	trees := score_run(the_map, 3, 1)
	
	fmt.Println(trees)
}