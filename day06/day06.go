//usr/bin/env go run $0 $@ ; exit
// for part 2, 3050 is too low, can't figure out my bug gonna move on for the moment

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type score struct {
	any int
	all int
}

func union(a []byte, b []byte) []byte {

	res := []byte{}
	checked_chars := make(map[byte]bool)
	OuterLoop: for _, qa := range a {
		checked, _ := checked_chars[qa]
		if checked {continue}
		checked_chars[qa] = true	
		inner_checked := make(map[byte]bool)
		for _, qb := range b {
			checked, _ := inner_checked[qb]
			if checked {continue}
			if qa == qb {
				res = append(res, qb)
				continue OuterLoop
			}
			inner_checked[qb] = true
		}
	}
	fmt.Printf("Union: %s | %s = %s\n", a, b, res)
	return res
}

func score_group(group []byte) score {
	fmt.Printf("%s\n", group)
	any_answers := make(map[byte] bool)
	passengers := bytes.Split(group, []byte{'\n'})
	init_all := true
	all_answers := []byte{}

	for _, passenger := range passengers {
		if init_all == true{
			all_answers = passenger
			init_all = false
		} else {
			all_answers = union(all_answers, passenger)
		}
		for _, question :=  range passenger {	
			any_answers[question] = true
		}
		
	}

	return score{any:len(any_answers), all:len(all_answers)}
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

	groups := bytes.Split(input, []byte{'\n', '\n'})

	any_score := 0
	all_score := 0
	group_score := score{}

	for group_i := 0; group_i < len(groups); group_i++ {
		group_score = score_group(groups[group_i])
		any_score += group_score.any
		all_score += group_score.all
		fmt.Printf("all score (running): %d\n", all_score)
	}
	fmt.Println("")
	fmt.Println(any_score)
	fmt.Println(all_score)
}
