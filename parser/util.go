package parser

import (
	"fmt"
	"strconv"
	"strings"
)

// defining the variable for validating string
var (
	fieldNames     = []string{"minute", "hour", "dayOfMonth", "month", "dayOfWeek"}
	errInvalidCron = fmt.Errorf("invalid number of fields in cron expression")
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
func validateString(fields []string) error {

	if len(fields) != 6 {
		return errInvalidCron
	}

	fields = fields[:len(fields)-1]

	// Define valid ranges for each field
	validRanges := map[string][2]int{
		"minute":     {0, 59},
		"hour":       {0, 23},
		"dayOfMonth": {1, 31},
		"month":      {1, 12},
		"dayOfWeek":  {0, 6},
	}

	for i, field := range fields {

		validRange := validRanges[fieldNames[i]]

		errMsg := fmt.Errorf("invalid value %s for %s field", field, fieldNames[i])

		if field == "*" {
			continue
		}

		if strings.Contains(field, "-") {
			bounds := strings.Split(field, "-")
			if len(bounds) != 2 {
				return errMsg
			}
			start, err := strconv.Atoi(bounds[0])
			if err != nil {
				return errMsg
			}
			end, err := strconv.Atoi(bounds[1])
			if err != nil {
				return errMsg
			}
			if start < validRange[0] || end > validRange[1] || start > end {
				return errMsg
			}
			continue
		}

		if strings.Contains(field, "/") {
			parts := strings.Split(field, "/")
			if len(parts) != 2 {
				return errMsg
			}

			if parts[0] != "*" {
				value, err := strconv.Atoi(parts[0])
				if err != nil {
					return errMsg
				}
				if value < validRange[0] || value > validRange[1] {
					return errMsg
				}
			}

			step, err := strconv.Atoi(parts[1])
			if err != nil || step <= 0 || step > validRange[1] {
				return errMsg
			}
			continue
		}

		value, err := strconv.Atoi(field)
		if err != nil {
			return errMsg
		}

		if value < validRange[0] || value > validRange[1] {
			return errMsg
		}

	}
	return nil
}
