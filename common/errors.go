package common

import "errors"

// define variables for errors
var (
	ErrInvalidCronExp        = errors.New("invalid fields in cron expression")
	ErrInvalidIntervalLimits = errors.New("invalid interval limit")
	ErrInvalidIntervalRanges = errors.New("invalid ranges specified for interval of field")
	ErrInvalidExpSteps       = errors.New("invalid steps specified for field")
)
