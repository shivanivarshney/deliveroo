package parser

import (
	"fmt"
	"strconv"
	"strings"
)

// defining the error variable for invalid cron string
var (
	errInvalidCron = fmt.Errorf("invalid cron string")
	// errInvalidMinute     = fmt.Errorf("invalid minute range")
	// errInvalidHour       = fmt.Errorf("invalid hour range")
	// errInvalidDayOfMonth = fmt.Errorf("invalid day of month range")
	// errInvalidMonth      = fmt.Errorf("invalid month range")
	// errInvalidDayOFWeek  = fmt.Errorf("invalid day of week range")
)

// function to convert int slice to final string
func intSliceToString(slice []int) string {
	strSlice := make([]string, len(slice))
	for i, val := range slice {
		strSlice[i] = strconv.Itoa(val)
	}
	return strings.Join(strSlice, " ")
}

// function to format the schedule as per the output
func cronToString(s schedule) string {

	output := "minute\t\t" + intSliceToString(s.minutes) + "\n"
	output += "hour\t\t" + intSliceToString(s.hours) + "\n"
	output += "day of month\t" + intSliceToString(s.daysOfMonth) + "\n"
	output += "month\t\t" + intSliceToString(s.months) + "\n"
	output += "day of week\t" + intSliceToString(s.daysOfWeek) + "\n"
	output += "command\t\t" + s.command
	return output

}

// function to validate our schedule values have all exact fields or not
func validateString(schedule []string) error {
	fmt.Println("schedule", schedule)
	if len(schedule) != 6 {
		return errInvalidCron
	}

	// minutes, err := strconv.Atoi(expression[0])
	// if err != nil {
	// 	// Handle the error if the string cannot be converted to an integer
	// 	return errInvalidMinute
	// }

	// // Handle the error if the string, if string is out of range
	// if minutes < 0 || minutes > 59 {
	// 	return errInvalidMinute
	// }

	// hour, err := strconv.Atoi(expression[1])
	// if err != nil {
	// 	// Handle the error if the string cannot be converted to an integer
	// 	return errInvalidHour
	// }

	// // Handle the error if the string, if string is out of range
	// if hour < 0 || hour > 23 {
	// 	fmt.Println(expression[1])
	// 	return errInvalidHour
	// }

	// daysOfMonth, err := strconv.Atoi(expression[2])
	// if err != nil {
	// 	// Handle the error if the string cannot be converted to an integer
	// 	return errInvalidDayOfMonth
	// }

	// // Handle the error if the string, if string is out of range
	// if daysOfMonth < 1 || daysOfMonth > 31 {
	// 	return errInvalidDayOfMonth
	// }

	// month, err := strconv.Atoi(expression[3])
	// if err != nil {
	// 	// Handle the error if the string cannot be converted to an integer
	// 	return errInvalidMonth
	// }

	// // Handle the error if the string, if string is out of range
	// if month < 1 || month > 12 {
	// 	fmt.Println(expression[3])
	// 	return errInvalidMonth
	// }

	// daysOfWeek, err := strconv.Atoi(expression[4])
	// if err != nil {
	// 	// Handle the error if the string cannot be converted to an integer
	// 	return errInvalidDayOFWeek
	// }

	// // Handle the error if the string, if string is out of range
	// if daysOfWeek < 0 || daysOfWeek > 6 {
	// 	return errInvalidDayOFWeek
	// }

	return nil
}
