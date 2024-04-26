package parser

import (
	"strconv"
	"strings"
)

type schedule struct {
	minutes     []int
	hours       []int
	daysOfMonth []int
	months      []int
	daysOfWeek  []int
	command     string
}

func ParseFields(fields []string) (string, error) {

	err := validateString(fields)

	if err != nil {
		return "", err
	}

	schedule := schedule{
		minutes:     expandField(fields[0], 0, 60),
		hours:       expandField(fields[1], 0, 24),
		daysOfMonth: expandField(fields[2], 1, 31),
		months:      expandField(fields[3], 1, 12),
		daysOfWeek:  expandField(fields[4], 1, 7),
		command:     fields[5],
	}

	return cronToString(schedule), nil

}

func expandField(field string, minVal, maxVal int) []int {
	if field == "*" {
		result := make([]int, maxVal)
		for i := 0; i < maxVal; i++ {
			result[i] = i
		}
		return result
	}

	var result []int
	if strings.Contains(field, "/") {
		parts := strings.Split(field, "/")
		step, _ := strconv.Atoi(parts[1])
		for i := 0; i < maxVal; i += step {
			result = append(result, i)
		}
		return result
	}

	ranges := strings.Split(field, ",")
	for _, r := range ranges {
		if strings.Contains(r, "-") {
			bounds := strings.Split(r, "-")
			start, _ := strconv.Atoi(bounds[0])
			end, _ := strconv.Atoi(bounds[1])
			for i := start; i <= end; i++ {
				result = append(result, i)
			}
		} else {
			val, _ := strconv.Atoi(r)
			result = append(result, val)
		}
	}
	return result
}
