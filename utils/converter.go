package utils

import (
	"math"
	"strconv"
	"time"
)

func ConvertStringToUint(value string) (uint, error) {
	parsedValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}

	return uint(parsedValue), nil
}

func ConvertFloatToTwoDecimals(value float64) float64 {
	return math.Floor(value*100) / 100
}

func ConvertStringToDate(value string) (time.Time, error) {
	layout := "2006-01-02"
	return time.Parse(layout, value)
}
