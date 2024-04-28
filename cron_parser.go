package main

import (
	parser "deliveroo/parser"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("herre..")

	// checking if no input provided through command line
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "<cron_string>")
		return
	}

	// extract the input from argument & split that based on " " (whitespaces)
	cronString := os.Args[1]
	fields := strings.Fields(cronString)

	// call parser function to parse fields to get the cron range
	output, err := parser.ParseFields(fields)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(output)

}
