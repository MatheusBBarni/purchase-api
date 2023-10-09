package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertStringToUint(t *testing.T) {
	value, err := ConvertStringToUint("1")

	assert.NoError(t, err)
	assert.Equal(t, value, uint(1))
}

func TestConvertStringToUintError(t *testing.T) {
	_, err := ConvertStringToUint("1.99")

	assert.Error(t, err)
}

func TestConvertFloatToTwoDecimals(t *testing.T) {
	value := ConvertFloatToTwoDecimals(20.998278)

	assert.Equal(t, value, 20.99)
}

func TestConvertStringToDate(t *testing.T) {
	value, err := ConvertStringToDate("2023-10-08")

	assert.NoError(t, err)
	assert.Contains(t, value.String(), "2023-10-08")
}

func TestConvertStringToDateError(t *testing.T) {
	_, err := ConvertStringToDate("2023-02-31")

	assert.Error(t, err)
}
