package main

import (
	parser "deliveroo/parser"
	"fmt"
	"os"
	"strings"
)

func main() {

	// checking if no input provided through command line
	if len(os.Args) != 2 {
		panic(fmt.Sprintf("Command line arguments supplied are not valid: Example format: %s", "*/15 0 1,15 * 1-5 /usr/bin/find"))
	}

	// extract the input from argument & split that based on " " (whitespaces)
	cronString := os.Args[1]
	fields := strings.Fields(cronString)

	// call parser function to parse fields to get the cron range
	output, err := parser.ParseFields(fields)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(output)

}
