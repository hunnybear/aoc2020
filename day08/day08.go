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

type path []int


type command struct {
	// I had this as [3]byte, but that might just be me totally missing the point
	instruction []byte
	value int
}

type swap struct {
	index int
	instruction []byte
}


func get_commands(instruction []byte, commands []command) map[int]int {
	res_map := make(map[int]int)
	for cmd_i, command := range commands {
		if bytes.Equal(command.instruction, instruction) {
			res_map[cmd_i] = command.value
		} 
	}
	return res_map
}

func filter_paths_by_origin(paths []path, origin int) []path {
	fmt.Printf("filter start:\n%d\n\n", paths)
	for i := len(paths) - 1; i >= 0; i-- {
		if paths[i][0] != origin {
			paths = append(paths[:i], paths[i+1:]...)
		}
	}
	fmt.Printf("fitler res:\n%d\n\n", paths)
	return paths
}

func find_swap(paths []path, commands []command, start int, jumps map[int]int) swap {
	// Find the swap that will get from target starting position to one of the 
	// provided paths

	checked := make(map[int]bool)
	//TOREMOVE
	if len(checked) > 30{return swap{}}

	for _, this_path := range paths {
		for _, position := range this_path {
			this_checked, ok := checked[position]
			if this_checked == true {
				continue
			} else if ok == false {
				fmt.Println(this_checked)
			}

			if bytes.Equal(commands[position-1].instruction, JMP) {
				paths_to_prev := trace_paths(path{position-1}, commands, jumps)
				
				fmt.Printf("pre filter jmp -> acc (%d):\n%d\n", position, paths_to_prev)
				paths_to_prev = filter_paths_by_origin(paths_to_prev, 0)

				if len(paths_to_prev) > 1 {
					log.Panic("I really don't know how")
				} else if len(paths_to_prev) == 1 {
					return swap {
						index: position - 1,
						instruction: ACC}
				}
			}

			accs := get_commands(ACC, commands)
			paths_to_prev := []path{}
			for acc_idx, acc_val := range accs {
				if acc_idx + acc_val == position {
					paths_to_prev = append(paths_to_prev, trace_paths(path{acc_idx}, commands, jumps)...)

				}
			}
			fmt.Printf("pre filter acc-> jmp:\n%d\n", paths_to_prev)
			paths_to_prev = filter_paths_by_origin(paths_to_prev, 0)
			if len(paths_to_prev) > 1 {
				log.Panic("this be bad")
			} else if len(paths_to_prev) == 1 {
				return swap{index: paths_to_prev[0][len(paths_to_prev[0])-1],
							instruction:JMP}
			}

			checked[position] = true
		} 
	}

	return swap{}
}


func trace_paths(end path, commands []command, jumps map[int]int) []path {
	paths := []path{}
	target := end[0]

	prev := []int{}

	if ! bytes.Equal(commands[target-1].instruction, JMP) {
		prev = append(prev, target-1)
	}

	for idx, jump_val := range jumps {
		if idx + jump_val == end[0] {
			fmt.Printf("%d hits %d (%d)\n", idx, end[0], jump_val)
			prev = append(prev, idx) 
		}
	}

	if len(prev) < 1 {
		paths = append(paths, []int{})
	}

	for _, prev_idx := range prev {

		for _, path_to_here := range trace_paths(path{prev_idx}, commands, jumps) {
			// each path_to_here is its own path to here. sensibly.
			fmt.Printf("path to here: %d\n\n", path_to_here)

			paths = append(paths, append(path_to_here, end...))
		}
	}

	return paths

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

	// part 2

	jumps := get_commands(JMP, commands)

	tails := trace_paths(path{len(commands)}, commands, jumps)
	fmt.Println(tails)

	res := find_swap(tails, commands, 0, jumps)
	fmt.Println(res)

}
