package errors

import "errors"

var (
	ErrEmptyMetricName = errors.New("metric name cannot be empty")
)
