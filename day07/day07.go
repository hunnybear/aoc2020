//usr/bin/env go run $0 $@ ; exit

package main

import (
	"fmt"
	"log"
)

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {log.Fatal(err)}

	fmt.Println(answer)
}
