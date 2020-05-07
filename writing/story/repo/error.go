package repo

import (
	"fmt"
	"reflect"
)

// RecordNotFoundError is used when repository search failed to find an entity
type RecordNotFoundError struct {
	Entity interface{}
	ID     int
}

func (e *RecordNotFoundError) Error() string {
	return fmt.Sprintf("%s with key %v is not found", e.getClassName(e.Entity), e.ID)
}

func (e *RecordNotFoundError) getClassName(object interface{}) string {
	return reflect.TypeOf(object).Name()
}
