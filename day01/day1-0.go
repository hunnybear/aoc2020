package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"strconv"
)

// Return number valid data looks like:
//  ```1-3 a: abcde
//     1-3 b: cdefg
//     2-9 c: ccccccccc
//  ```

var re = regexp.MustCompile(`(\d+)-(\d+)\s+([a-zA-Z]):\s+([a-zA-Z]+)$`)

type valid struct {
	part_0 bool
	part_1 bool
}

func is_valid(line string) valid {
	res := re.FindAllStringSubmatch(line, -1)
	min, err := strconv.Atoi(res[0][1])
	if err != nil {
		fmt.Println(line)
		fmt.Println(res)
		log.Fatal(err)
	}	
	max, err := strconv.Atoi(res[0][2])
	if err != nil {
		fmt.Println(line)
		fmt.Println(res)
		log.Fatal(err)
	}

	target := res[0][3]
	full_string := res[0][4]

	count := strings.Count(full_string, target)

	// part 2 is min/max are 1-indexed indices. Valid == one and only one of the
	// indicated indices indicates the input string



	return valid{
		part_0: min <= count && count <= max,
		part_1: (full_string[min + 1] == target) != (full_string[max+1] == target)
	}
}

func is_valid_part_1(line string) bool {

}


func main() {
	file, err := os.Open("data")
	valid_part_0 := 0
	valid_part_1 := 0

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		validity := is_valid(scanner.Text())
		fmt.Println(validity)
		if validity.part_0 {
			valid_part_0 += 1
		}
		if validity.part_1 {
			valid_part_1 += 1
		}
	}

	fmt.Println(valid_part_0)
	fmt.Println(valid_part_1)
}