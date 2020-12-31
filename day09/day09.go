//usr/bin/env go run $0 $@ ; exit

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const TEST_RESULT := 0

func main() {
	input_file := ""
	
	if len(os.Args) > 1  {
		testing := true
		input_file = "test_input"
	} else {
		testing := false
		input_file = "input"
	}
	input, err := ioutil.ReadFile(input_file)
	if err != nil {log.Fatal(err)}

	fmt.Println(answer)

	if testing && answer != TEST_RESULT {
		log.Error(fmt.Sprintf"Test answer was wrong, got %d, expected %d", TEST_RESULT, answer)
	}
}
