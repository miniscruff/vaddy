package vaddy

import (
	"errors"
	"strconv"
	"strings"
)

// ValidationError is what we return on invalid validations.
// Output of the error is a single string of:
// `Key [Index] Message [(Help)]`
type ValidationError struct {
	Key     string
	Message string
	Help    string
	Index   *int
}

func (v *ValidationError) Error() string {
	if v == nil {
		return ""
	}

	sb := &strings.Builder{}
	sb.WriteString(v.Key)

	if v.Index != nil {
		sb.WriteString("[")
		sb.WriteString(strconv.Itoa(*v.Index))
		sb.WriteString("]")
	}

	sb.WriteString(" ")
	sb.WriteString(v.Message)

	if v.Help != "" {
		sb.WriteString(" ( ")
		sb.WriteString(v.Help)
		sb.WriteString(" )")
	}

	return sb.String()
}

func expandErrorKey(err error, key string) error {
	// TODO: use errors.Is???
	ve, isValidationError := err.(*ValidationError)
	if !isValidationError {
		return &ValidationError{
			Message: err.Error(),
			Key:     key,
		}
	}

	ve.Key = key

	return ve
}

func expandErrorIndex(err error, index int) error {
	// TODO: use errors.Is???
	ve, isValidationError := err.(*ValidationError)
	if !isValidationError {
		return &ValidationError{
			Message: err.Error(),
			Index:   &index,
		}
	}

	ve.Index = &index

	return ve
}

func expandErrorKeyIndex(err error, key string, index int) error {
	// TODO: use errors.Is???
	ve, isValidationError := err.(*ValidationError)
	if !isValidationError {
		return &ValidationError{
			Message: err.Error(),
			Index:   &index,
			Key:     key,
		}
	}

	ve.Index = &index
	ve.Key = key

	return ve
}

// Join combines all errors into one.
func Join(errs ...error) error {
	return errors.Join(errs...)
}

// JoinAnd will return the first non-nil error,
// or nil if all errors are nil.
func JoinAnd(errs ...error) error {
	for _, e := range errs {
		if e != nil {
			return e
		}
	}

	return nil
}

// TODO:
// ValdiationErrors type
// Join should build ^
// Error() for ^ should support TextMarshal, JSONMarshal
