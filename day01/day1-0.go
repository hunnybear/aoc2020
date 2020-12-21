package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

//var pattern regexp.Regexp = regexp.Compile("^(\d+)-(\d+)\s+(\w)\s+(\w+)\s*$")

// Return number valid data looks like:
//  ```1-3 a: abcde
//     1-3 b: cdefg
//     2-9 c: ccccccccc
//  ```

func is_valid(line string) bool {
	//match := pattern.FindAllStringSubmatch(line, 1)
	//fmt.Println(match)
	//return match != nil
	return true
}


func main() {
	file, err := os.Open("data")
	valid := 0

	re := regexp.MustCompile(`a(x*)b`)
	// Indices:
	//    01234567   012345678
	//    -ab-axb-   -axxb-ab-
	fmt.Println(re.FindAllStringSubmatchIndex("-ab-", -1))
	fmt.Println(re.FindAllStringSubmatchIndex("-axxb-", -1))
	fmt.Println(re.FindAllStringSubmatchIndex("-ab-axb-", -1))
	fmt.Println(re.FindAllStringSubmatchIndex("-axxb-ab-", -1))
	fmt.Println(re.FindAllStringSubmatchIndex("-foo-", -1))

	fmt.Println("be-gin")

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