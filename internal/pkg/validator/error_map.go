package validator

import (
	"fmt"
	"strings"
)

type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

type ErrorMap struct {
	Errors map[string]error
}

func NewErrorMap(errs map[string]error) error {
	return &ErrorMap{Errors: errs}
}

func (e ErrorMap) Error() string {
	errorMessage := make([]string, 0)

	for key, er := range e.Errors {
		errorMessage = append(errorMessage, fmt.Sprintf("%s:%s", key, er.Error()))
	}
	return strings.Join(errorMessage, ";")
}
