package main

import (
	parser "deliveroo/parser"
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "<cron_string>")
		return
	}

	cronString := os.Args[1]
	fields := strings.Fields(cronString)

	output, err := parser.ParseFields(fields)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(output)

}
