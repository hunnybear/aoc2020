//usr/bin/env go run $0 $@ ; exit

package main

import (
	"bytes"
	"io/ioutil"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type validator func(string) bool

var required_fields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
var optional_fields = []string{"cid"}
var eye_colors = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

var color_re = regexp.MustCompile(`^#[0-9a-f]{6}$`)
var pid_re = regexp.MustCompile(`^[0-9]{9}$`)

var part_2_validators = make(map[string]validator)

func get_fields(record []byte) map[string]string {
	fields := make(map[string]string)
	split_record := strings.Fields(string(record))
	for i := 0; i < len(split_record); i++ {
		key_val := strings.Split(split_record[i], ":")
		fields[key_val[0]] = key_val[1]
	}
	return fields
}

func check_record_part_1(record []byte, required []string) bool {
	// miight make sense to roll both of these into one,
	// but it's late, this is a toy, and I'm lazy
	fields := get_fields(record)
	for i := 0; i < len(required); i++ {
		// check for missing required records
		_, ok := fields[required[i]]

		if ok != true {
			return false
		}
	}
	return true
}

func check_record_part_2(record []byte, required []string) bool {
	fields := get_fields(record)

	for i := 0; i < len(required); i++ {
		val, ok := fields[required[i]]
		if ok != true {
			return false
		}
		this_validator, ok := part_2_validators[required[i]]
		if ok != true {
			fmt.Printf("%s is missing", required[i])
			log.Panic("Missing validator!")
		}
		if this_validator(val) == false {
			fmt.Printf("%s failed with %s\n", required[i], val)
			return false
		}
	}

	return true
}

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {log.Fatal(err)}

	// Setup the part 2 validators
	part_2_validators["byr"] = byr
	part_2_validators["iyr"] = iyr
	part_2_validators["eyr"] = eyr
	part_2_validators["hgt"] = hgt
	part_2_validators["hcl"] = color_re.MatchString
	part_2_validators["ecl"] = ecl
	part_2_validators["pid"] = pid_re.MatchString


	records := bytes.Split(bytes.TrimFunc(input, unicode.IsSpace), []byte("\n\n"))
	valid_part_1 := 0
	valid_part_2 := 0
	for i := 0; i < len(records); i++ {
		if check_record_part_1(records[i], required_fields) == true {
			valid_part_1 += 1
		}
		if check_record_part_2(records[i], required_fields) {
			valid_part_2 += 1
		}
	}

	fmt.Println(valid_part_1)
	fmt.Println(valid_part_2)
}

func int_between(value string, min int, max int) bool {
	int_val, err := strconv.Atoi(value)
	if err != nil {
		return false
	}
	if int_val >= min && int_val <= max {
		return true
	}
	return false
}
func byr(year string) bool {
	return int_between(year, 1920, 2002)
}

func iyr(year string) bool {
	return int_between(year, 2010, 2020)
}

func eyr(year string) bool {
	return int_between(year, 2020, 2030)
}

func hgt(height string) bool {
	if height[len(height) - 2: len(height)] == "in" {
		return int_between(height[0:len(height)-2], 59, 76)
	} else if height[len(height) - 2: len(height)] == "cm" {
		return int_between(height[0:len(height) -2], 150, 193)
	} else {
		return false
	}
}

func ecl(color string) bool {
	for i := 0; i < len(eye_colors); i++ {
		if eye_colors[i] == color {
			return true
		}
	}
	return false
}

