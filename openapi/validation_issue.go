package openapi

import "reflect"

type ValidationIssueCode int

const (
	ValidationIssueCodeRequired = iota + 1
	ValidationIssueCodeNullable
	ValidationIssueCodeMin
	ValidationIssueCodeMax
	ValidationIssueCodeMinLength
	ValidationIssueCodeMaxLength
	ValidationIssueCodeEnum
	ValidationIssueCodePattern
)

type ValidationIssue struct {
	Code     ValidationIssueCode
	Name     string
	Value    reflect.Value
	Expected interface{}
	Actual   interface{}
}

func NewValidationIssue(
	code ValidationIssueCode,
	name string,
	value reflect.Value,
	expected interface{},
	actual interface{},
) *ValidationIssue {
	return &ValidationIssue{
		Code:     code,
		Name:     name,
		Value:    value,
		Expected: expected,
		Actual:   actual,
	}
}
