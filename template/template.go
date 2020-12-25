//usr/bin/env go run $0 $@ ; exit

package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	input_file := ""
	
	if len(os.Args) > 1  {
		input_file = "test_input"
	} else {
		input_file = "input"
	}
	input, err := ioutil.ReadFile(input_file)
	if err != nil {log.Fatal(err)}

	fmt.Println(answer)
}
