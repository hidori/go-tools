package openapi

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/hidori/go-tools/runtime"
	"github.com/pkg/errors"
)

type Validator interface {
	Validate(schema *openapi3.Schema, name string, value interface{}) ([]*ValidationIssue, error)
}

func NewValidator() Validator {
	return &validator{}
}

type validator struct{}

func (v *validator) Validate(schema *openapi3.Schema, name string, value interface{}) ([]*ValidationIssue, error) {
	rv := reflect.ValueOf(value)
	rt := reflect.TypeOf(value)

	issues, err := v.validate(schema, name, rv, rt.Kind())
	if err != nil {
		return nil, errors.Wrap(err, "fail to validate()")
	}

	return issues, nil
}

func (v *validator) validate(schema *openapi3.Schema, name string, value reflect.Value, kind reflect.Kind) ([]*ValidationIssue, error) {
	switch kind {
	case reflect.Bool:
		return nil, nil

	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Float32,
		reflect.Float64:
		return v.validateNumber(schema, name, value)

	case reflect.String:
		return v.validateString(schema, name, value)

	case reflect.Pointer:
		return v.validatePointer(schema, name, value)

	case reflect.Slice:
		return v.validateSlice(schema, name, value)

	case reflect.Struct:
		return v.validateStruct(schema, name, value)
	}

	return nil, fmt.Errorf("unsupported type '%v' of '%s'", kind, name)
}

func (v *validator) validateNumber(schema *openapi3.Schema, name string, value reflect.Value) ([]*ValidationIssue, error) {
	var issues []*ValidationIssue

	{
		actual := value.Convert(reflect.TypeOf(float64(0))).Float()

		if schema.Min != nil {
			expected := *schema.Min
			if actual < expected {
				issues = append(issues, NewValidationIssue(ValidationIssueCodeMin, name, value, expected, actual))
			}
		}

		if schema.Max != nil {
			expected := *schema.Max
			if expected < actual {
				issues = append(issues, NewValidationIssue(ValidationIssueCodeMax, name, value, expected, actual))
			}
		}
	}

	{
		_issues, err := v.validateEnum(schema, name, value)
		if err != nil {
			return nil, errors.Wrap(err, "fail to validateEnum()")
		}

		issues = append(issues, _issues...)
	}

	return issues, nil
}

func (v *validator) validateString(schema *openapi3.Schema, name string, value reflect.Value) ([]*ValidationIssue, error) {
	var issues []*ValidationIssue

	{
		actual := uint64(len(value.String()))

		{
			expected := schema.MinLength
			if actual < expected {
				issues = append(issues, NewValidationIssue(ValidationIssueCodeMinLength, name, value, expected, actual))
			}
		}

		if schema.MaxLength != nil {
			expected := *schema.MaxLength
			if expected < actual {
				issues = append(issues, NewValidationIssue(ValidationIssueCodeMaxLength, name, value, expected, actual))
			}
		}
	}

	if schema.Pattern != "" {
		expected := schema.Pattern
		actual := value.String()

		pattern, err := regexp.Compile(expected)
		if err != nil {
			return nil, errors.Wrap(err, "fail to regexp.Compile()")
		}

		if !pattern.MatchString(actual) {
			issues = append(issues, NewValidationIssue(ValidationIssueCodePattern, name, value, expected, actual))
		}
	}

	{
		_issues, err := v.validateEnum(schema, name, value)
		if err != nil {
			return nil, errors.Wrap(err, "fail to validateEnum()")
		}

		issues = append(issues, _issues...)
	}

	return issues, nil
}

func (v *validator) validateEnum(schema *openapi3.Schema, name string, value reflect.Value) ([]*ValidationIssue, error) {
	if len(schema.Enum) < 1 {
		return nil, nil
	}

	var (
		issues   []*ValidationIssue
		contains bool
	)

	expected := schema.Enum
	actual := value

	for i := range expected {
		rv := reflect.ValueOf(expected[i])

		_expected, err := runtime.Recover1(func() (reflect.Value, error) {
			return rv.Convert(actual.Type()), nil
		})
		if err != nil {
			return nil, errors.Wrap(err, "fail to rv.Convert()")
		}

		if _expected.Equal(actual) {
			contains = true
		}
	}

	if !contains {
		issues = append(issues, NewValidationIssue(ValidationIssueCodeEnum, name, value, expected, actual))
	}

	return issues, nil
}

func (v *validator) validatePointer(schema *openapi3.Schema, name string, value reflect.Value) ([]*ValidationIssue, error) {
	if value.IsNil() {
		if schema.Nullable {
			return nil, nil
		}

		return []*ValidationIssue{
			NewValidationIssue(ValidationIssueCodeNullable, name, value, schema.Nullable, !schema.Nullable),
		}, nil
	}

	rv := value.Elem()
	rt := rv.Type()

	issues, err := v.validate(schema, name, rv, rt.Kind())
	if err != nil {
		return nil, errors.Wrap(err, "fail to validate()")
	}

	return issues, nil
}

func (v *validator) validateSlice(schema *openapi3.Schema, name string, value reflect.Value) ([]*ValidationIssue, error) {
	var issues []*ValidationIssue

	for i := 0; i < value.Len(); i++ {
		name = fmt.Sprintf("%s[%d]", name, i)

		rv := value.Index(i)
		rt := rv.Type()

		_issues, err := v.validate(schema.Items.Value, name, rv, rt.Kind())
		if err != nil {
			return nil, errors.Wrap(err, "fail to validate()")
		}

		issues = append(issues, _issues...)
	}

	return issues, nil
}

func (v *validator) validateStruct(schema *openapi3.Schema, name string, value reflect.Value) ([]*ValidationIssue, error) {
	var issues []*ValidationIssue

	for _name, schemaRef := range schema.Properties {
		field := v.getStructField(value, _name)

		if name != "" {
			_name = fmt.Sprintf("%s.%s", name, _name)
		}

		if field == nil {
			return nil, fmt.Errorf("no field for '%s'", _name)
		}

		rv := value.FieldByName(field.Name)
		rt := rv.Type()

		_issues, err := v.validate(schemaRef.Value, _name, rv, rt.Kind())
		if err != nil {
			return nil, errors.Wrap(err, "fail to v.Validate()")
		}

		issues = append(issues, _issues...)
	}

	return issues, nil
}

func (v *validator) getStructField(rv reflect.Value, propertyName string) *reflect.StructField {
	count := rv.NumField()

	for i := 0; i < count; i++ {
		field := rv.Type().Field(i)

		tag := strings.Split(field.Tag.Get("json"), ",")
		if len(tag) < 1 {
			continue
		}

		if tag[0] == propertyName {
			return &field
		}
	}

	return nil
}
