package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type StructTest struct {
	Name string `validate:"required"`
}

func TestValidateStruct(t *testing.T) {
	test := &StructTest{Name: "Test"}
	err := ValidateStruct(test)

	assert.NoError(t, err)
}

func TestValidateStructError(t *testing.T) {
	test := &StructTest{}
	err := ValidateStruct(test)

	assert.Error(t, err)
}
