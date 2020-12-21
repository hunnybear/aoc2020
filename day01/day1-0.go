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

func is_valid(line string) bool {
	fmt.Println(line)
	res := re.FindAllStringSubmatch(line, -1)
	fmt.Println(res)
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
	count := strings.Count(res[0][4], res[0][3])

	return min <= count && count <= max

	return true
}


func main() {
	file, err := os.Open("data")
	valid := 0

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if is_valid(scanner.Text()){
			valid += 1
		}
	}

	fmt.Println(valid)
}