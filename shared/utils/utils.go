package utils

import "fmt"

func GetValueOrDefault[T any](value *T) T {
	if value == nil {
		var zeroValue T
		return zeroValue
	}
	return *value
}

func FormatStringPtr(p *string) string {
	if p == nil {
		return "n/a"
	}
	return *p
}

func FormatFloatPtr(p *float64) string {
	if p == nil {
		return "n/a"
	}
	return fmt.Sprintf("%.8f", *p)
}

func FormatInt64Ptr(p *int64) string {
	if p == nil {
		return "n/a"
	}
	return fmt.Sprintf("%d", *p)
}
