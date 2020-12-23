//usr/bin/env go run $0 $@ ; exit

package main

import (
	"bytes"
	"io/ioutil"
	"fmt"
	"log"
	"strings"
	"unicode"
)

var part_1_required = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
var part_1_optional = []string{"cid"}

func get_fields(record []byte) map[string]string {
	fields := make(map[string]string)
	split_record := strings.Fields(string(record))
	for i := 0; i < len(split_record); i++ {
		key_val := strings.Split(split_record[i], ":")
		fields[key_val[0]] = key_val[1]
	}
	return fields
}

func check_record(record []byte, required []string) bool {
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

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {log.Fatal(err)}

	records := bytes.Split(bytes.TrimFunc(input, unicode.IsSpace), []byte("\n\n"))
	valid := 0
	for i := 0; i < len(records); i++ {
		if check_record(records[i], part_1_required) == true {
			valid += 1
		}
	}
	fmt.Println(valid)
}