package parser

import (
	"fmt"
	"strconv"
	"strings"
)

func intSliceToString(slice []int) string {
	strSlice := make([]string, len(slice))
	for i, val := range slice {
		strSlice[i] = strconv.Itoa(val)
	}
	return strings.Join(strSlice, " ")
}

func cronToString(s schedule) string {

	output := "minute\t\t" + intSliceToString(s.minutes) + "\n"
	output += "hour\t\t" + intSliceToString(s.hours) + "\n"
	output += "day of month\t" + intSliceToString(s.daysOfMonth) + "\n"
	output += "month\t\t" + intSliceToString(s.months) + "\n"
	output += "day of week\t" + intSliceToString(s.daysOfWeek) + "\n"
	output += "command\t\t" + s.command
	return output

}

func validateString(expression []string) error {
	if len(expression) != 6 {
		return fmt.Errorf("invalid cron string")
	}
	return nil
}
