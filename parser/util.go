package parser

import (
	"deliveroo/common"
	"fmt"
	"strconv"
	"strings"
)

// intSliceToString convert int slice to final string
func intSliceToString(slice []int) string {
	strSlice := make([]string, len(slice))
	for i, val := range slice {
		strSlice[i] = strconv.Itoa(val)
	}
	return strings.Join(strSlice, " ")
}

// getFieldValues get fields values from schedule
func (s *schedule) getFieldValues(field string) []int {
	switch field {
	case "minute":
		return s.minutes
	case "hour":
		return s.hours
	case "dayOfMonth":
		return s.daysOfMonth
	case "month":
		return s.months
	case "dayOfWeek":
		return s.daysOfWeek
	default:
		return nil
	}
}

// cronToString format the schedule as per the output
func cronToString(s schedule) string {

	var lines []string
	maxFieldLength := 0

	// Calculate the maximum field length
	for _, field := range common.Fields {
		if len(field) > maxFieldLength {
			maxFieldLength = len(field)
		}
	}

	// Generate lines for each field
	for _, field := range common.Fields {

		var line string

		if field == "command" {
			line = fmt.Sprintf("%s%s", "command", strings.Repeat(" ", maxFieldLength+2-len("command"))+s.command)
		} else {
			line = fmt.Sprintf("%s%s%s", field, strings.Repeat(" ", maxFieldLength+2-len(field)), intSliceToString(s.getFieldValues(field)))

		}
		lines = append(lines, line)
	}

	return strings.Join(lines, "\n")

}

// validateString validate our schedule values have all exact fields or not (#modifiable when expansion for new strings)
func validateString(fields []string) error {

	if len(fields) != 6 {
		return common.ErrInvalidCronExp
	}

	fields = fields[:len(fields)-1]

	for i, field := range fields {

		validRange := common.ValidRanges[common.Fields[i]]

		if field == "*" {
			continue
		}

		if strings.Contains(field, "-") {
			bounds := strings.Split(field, "-")
			if len(bounds) != 2 {
				return common.ErrInvalidIntervalRanges
			}
			start, err := strconv.Atoi(bounds[0])
			if err != nil {
				return common.ErrInvalidCronExp
			}
			end, err := strconv.Atoi(bounds[1])
			if err != nil {
				return common.ErrInvalidCronExp
			}
			if start < validRange[0] || end > validRange[1] || start > end {
				return common.ErrInvalidIntervalLimits
			}
			continue
		}

		if strings.Contains(field, "/") {
			parts := strings.Split(field, "/")
			if len(parts) != 2 {
				return common.ErrInvalidIntervalRanges
			}

			if parts[0] != "*" {
				value, err := strconv.Atoi(parts[0])
				if err != nil {
					return common.ErrInvalidCronExp
				}
				if value < validRange[0] || value > validRange[1] {
					return common.ErrInvalidIntervalLimits
				}
			}

			step, err := strconv.Atoi(parts[1])
			if err != nil || step <= 0 || step > validRange[1] {
				return common.ErrInvalidExpSteps
			}
			continue
		}

		value, err := strconv.Atoi(field)
		if err != nil {
			return common.ErrInvalidCronExp
		}

		if value < validRange[0] || value > validRange[1] {
			return common.ErrInvalidIntervalLimits
		}

	}
	return nil
}
