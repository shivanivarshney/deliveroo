package common

// specify all fields allowed
var Fields = []string{"minute", "hour", "dayOfMonth", "month", "dayOfWeek", "command"}

// Define valid ranges for each field
var ValidRanges = map[string][2]int{
	"minute":     {0, 59},
	"hour":       {0, 23},
	"dayOfMonth": {1, 31},
	"month":      {1, 12},
	"dayOfWeek":  {0, 6},
}
