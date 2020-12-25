//usr/bin/env go run $0 $@ ; exit
// test result must be 4

package main

import (
	"bytes"
	"fmt"
	"log"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"strconv"
	"unicode"
)

const NO_OTHER = "no other bags"

type ruleset_entry map[string] int
type ruleset map[string] ruleset_entry

type contain_res struct {
	contains bool
	total int
}

// These amount to the same structure but I want to keep them separate

type cache_entry map[string] int
type cache map[string] cache_entry

var contains_cache = make(cache)

var rule_re = regexp.MustCompile(`^(.*) bags contain (.*)\.`)
var contents_re = regexp.MustCompile(`(\d+) ([a-zA-Z ]+) bags?|^no other bags$`)

func contain_count(kind string, container string, rules ruleset) int {

	res, ok := contains_cache[container][kind]
	if ok {return res}

	count := 0

	for next_container, next_count := range rules[container] {
		if next_container == kind {
			return next_count
		} else {
			count += next_count * contain_count(kind, next_container, rules)
		}
	}
	contains_cache[container][kind] = count
	return count
}

func max_container_count(kind string, rules ruleset) int {
	count := 0
	for contained_kind, contained_kind_count := range rules[kind] {
		count += contained_kind_count * (1 + max_container_count(contained_kind, rules))
	}

	return count
}

func can_contain_kind(kind string, container string, count int, rules ruleset) contain_res {
	count_res := contain_count(kind, container, rules)
	return contain_res{contains:count_res > 0, total:count_res} 
}

func generate_ruleset_and_cache(input string) ruleset {
	my_ruleset := make(ruleset)
	for _, line := range strings.Split(input, "\n") {

		match := rule_re.FindAllStringSubmatch(line, -1)
			
		kind := match[0][1]

		contains_cache[kind] = make(cache_entry)

		my_ruleset[kind] = make(ruleset_entry)

		contents_match := contents_re.FindAllStringSubmatch(match[0][2], -1)
		if contents_match[0][0] == NO_OTHER{continue}
		//fmt.Println(kind)
		for _, match := range contents_match {

				entry_count, err := strconv.Atoi(string(match[1]))
				if err != nil {
					log.Fatal(err)
				}
			
				//fmt.Printf("entry (%d) is %s\n", entry_count, contents_entry[2])

				my_ruleset[kind][match[2]] = entry_count
			}
		
		//fmt.Println(contents_match)
	}
	return my_ruleset
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

	my_ruleset := generate_ruleset_and_cache(string(input))

	able_to_contain := make(map[string]int)
	total_able_gold := 0

	for kind, _ := range my_ruleset {
		does_contain := can_contain_kind("shiny gold", kind, 1, my_ruleset)
		if does_contain.contains {total_able_gold++}
		able_to_contain[kind] = does_contain.total

	}

	fmt.Println(total_able_gold)
	fmt.Println(max_container_count("shiny gold", my_ruleset))

}
