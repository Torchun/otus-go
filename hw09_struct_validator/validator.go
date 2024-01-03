package hw09structvalidator

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// define error types to be printed.
var (
	errTypeNotStruct = fmt.Errorf("interface type is not string")
	errCaseDefault   = fmt.Errorf("tag validation case fallback to default")
	errCaseMismatch  = fmt.Errorf("tag validation case mismatch")
	errCaseSubsetLen = fmt.Errorf("tag validation case subset len mismatch")
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

// concat all matched errors into one string.
func (v ValidationErrors) Error() string {
	var resultErr string
	for _, item := range v {
		resultErr += fmt.Sprintf("field: %s | err: %s\n", item.Field, item.Err)
	}

	return resultErr
}

func Validate(v interface{}) error {
	var errSlice ValidationErrors

	// get reflection on passed interface type & value
	rType := reflect.TypeOf(v)
	rValue := reflect.ValueOf(v)

	// ensure passed interface kind is structure
	if rType.Kind().String() != "struct" {
		return errTypeNotStruct
	}

	// I want to keep debug lines for future understanding
	// fmt.Printf("fields count: %s\n", strconv.Itoa(rType.NumField()))

	// extract each field's type, value, tag=="validate"
	for i := 0; i < rType.NumField(); i++ {
		xType := rType.Field(i)
		xValue := rValue.Field(i)
		tagValue := xType.Tag.Get("validate")

		// skip any field to next i except tagged as "validate"
		if tagValue == "" {
			continue
		}
		// I want to keep debug lines for future understanding
		// fmt.Printf("xType: %s\n", xType.Type.String())
		// fmt.Printf("xValue: %s\n", xValue.String())
		// fmt.Printf("tagValue: %s\n", tagValue)

		switch xType.Type.String() {
		case "string":
			err := tagStringValidate(xValue.String(), tagValue)
			if err != nil {
				errSlice = append(errSlice, ValidationError{
					Field: xType.Name,
					Err:   err,
				})
			}

		case "[]string":
			for _, item := range xValue.Interface().([]string) {
				err := tagStringValidate(item, tagValue)
				if err != nil {
					errSlice = append(errSlice, ValidationError{
						Field: xType.Name,
						Err:   err,
					})
				}
			}

		case "int":
			err := tagIntValidate(xValue.Interface().(int), tagValue)
			if err != nil {
				errSlice = append(errSlice, ValidationError{
					Field: xType.Name,
					Err:   err,
				})
			}

		case "[]int":
			for _, item := range xValue.Interface().([]int) {
				err := tagIntValidate(item, tagValue)
				if err != nil {
					errSlice = append(errSlice, ValidationError{
						Field: xType.Name,
						Err:   err,
					})
				}
			}

		default:
			continue
		}
	}

	return errSlice
}

func tagStringValidate(data string, tag string) error {
	anyRule := strings.Split(tag, "|")

	for _, value := range anyRule {
		rulesSlice := strings.Split(value, ":")
		switch rulesSlice[0] {
		case "len":
			intValue, err := strconv.Atoi(rulesSlice[1])
			if err != nil {
				return fmt.Errorf("atoi error - %w", err)
			}

			if len(data) != intValue {
				return fmt.Errorf("tagStringValidate len error - %w", errCaseMismatch)
			}
		case "regexp":
			matchString, err := regexp.MatchString(rulesSlice[1], data)
			if err != nil {
				return fmt.Errorf("tagStringValidate regexp error - %w", err)
			}

			if !matchString {
				return fmt.Errorf("tagStringValidate regexp mismatch error - %w", errCaseMismatch)
			}
		case "in":
			for _, item := range strings.Split(rulesSlice[1], ",") {
				if !strings.Contains(data, item) {
					return fmt.Errorf("tagStringValidate in contains error - %w", errCaseMismatch)
				}
			}
		default:
			return fmt.Errorf("tagStringValidate default error - %w", errCaseDefault)
		}
	}

	return nil
}

func tagIntValidate(data int, tag string) error {
	anyRule := strings.Split(tag, "|")

	for _, value := range anyRule {
		rulesSlice := strings.Split(value, ":")
		switch rulesSlice[0] {
		case "min":
			intValue, err := strconv.Atoi(rulesSlice[1])
			if err != nil {
				return fmt.Errorf("atoi error - %w", err)
			}

			if data < intValue {
				return fmt.Errorf("tagIntValidate min error - %w", errCaseMismatch)
			}
		case "max":
			intValue, err := strconv.Atoi(rulesSlice[1])
			if err != nil {
				return fmt.Errorf("atoi error - %w", err)
			}

			if data > intValue {
				return fmt.Errorf("tagIntValidate max error - %w", errCaseMismatch)
			}
		case "in":
			sliceTagValue := strings.Split(rulesSlice[1], ",")
			if len(sliceTagValue) != 2 {
				fmt.Printf("len > 2: %s", sliceTagValue)
				return fmt.Errorf("tagIntValidate in error - %w", errCaseSubsetLen)
			}

			intValueMin, err := strconv.Atoi(sliceTagValue[0])
			if err != nil {
				return fmt.Errorf("tagIntValidate in atoi min error - %w", err)
			}

			intValueMax, err := strconv.Atoi(sliceTagValue[0])
			if err != nil {
				return fmt.Errorf("tagIntValidate in atoi max error - %w", err)
			}

			if data < intValueMin || data > intValueMax {
				return fmt.Errorf("tagIntValidate in error - %w", errCaseMismatch)
			}
		default:
			return fmt.Errorf("tagIntValidate default error - %w", errCaseDefault)
		}
	}

	return nil
}
