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

func main() {
	trees := 0
	
	input, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	the_map := bytes.Split(bytes.TrimFunc(input, unicode.IsSpace), []byte("\n"))

	x := 0
	
	for y := 0; y < len(the_map); y ++ {
		if the_map[y][x % len(the_map[y])] == tree {
			trees ++
		}

		x += 3
	}
	fmt.Println(trees)
}