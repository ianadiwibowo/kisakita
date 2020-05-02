package handler

import "fmt"

// InvalidValidationError is used mostly in request payload validation
type InvalidValidationError struct {
	Detail string
}

func (e *InvalidValidationError) Error() string {
	return e.Detail
}

// InvalidAttributeError is used mostly in request payload validation
type InvalidAttributeError struct {
	Field        string
	ValidatorTag string
}

func (e *InvalidAttributeError) Error() string {
	return fmt.Sprintf("%s is %s", e.Field, e.ValidatorTag)
}
