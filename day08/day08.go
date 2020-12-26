//usr/bin/env go run $0 $@ ; exit
// test expected answer: 5

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"unicode"
)

var JMP = []byte{'j', 'm', 'p'}
var ACC = []byte{'a', 'c', 'c'}
var NOP = []byte{'n', 'o', 'p'}

type command struct {
	// I had this as [3]byte, but that might just be me totally missing the point
	instruction []byte
	value int
}

func main() {
	input_file := ""
	
	if len(os.Args) > 1  {
		input_file = "test_input"
	} else {
		input_file = "input"
	}
	input, err := ioutil.ReadFile(input_file)
	if err != nil {log.Fatal(err)}
	input = bytes.TrimFunc(input, unicode.IsSpace)
	lines  := bytes.Split(input, []byte{'\n'})
	answer := 0

	commands := make([]command, len(lines), len(lines))

	
	for i, line := range lines {
		value, err := strconv.Atoi(string(line[4:]))
		if err != nil {
			log.Fatal(err)
		}

		commands[i] = command{
			instruction: line[:3],
			value: value}
	}
	seen := make(map[int]bool)
	iterations := 0
	for i:= 0; i < len(commands); i++ {
		this_seen, ok := seen[i]
		if this_seen && ok {break} 
		command := commands[i]
		if bytes.Equal(command.instruction, JMP) {
			i += commands[i].value -1 // (-1 to account for the loop's ++)
		} else if bytes.Equal(command.instruction, ACC) {
			answer += command.value
		}
		iterations ++
		seen[i] = true
	}


	fmt.Println(answer)
}
